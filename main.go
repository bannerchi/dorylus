package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/bannerchi/dorylus/models"
	"github.com/bannerchi/dorylus/tcp"
	Config "github.com/bannerchi/dorylus/util/config"
	Filter "github.com/bannerchi/dorylus/util/filter"
)

type Callback struct{}

func (this *Callback) OnConnect(c *tcp.Conn) bool {
	addr := c.GetRawConn().RemoteAddr()
	c.PutExtraData(addr)
	fmt.Println("OnConnect:", addr)
	return true
}

func (this *Callback) OnMessage(c *tcp.Conn, p tcp.Packet) bool {
	//check type
	echoPacket := p.(*tcp.EchoPacket)
	var resvMsg = new(tcp.EchoPacket)
	req := string(echoPacket.GetBody())
	resv := Filter.ResponsFilter(req)

	resvMsg = tcp.NewEchoPacket(resv, false)

	c.AsyncWritePacket(tcp.NewEchoPacket(resvMsg.Serialize(), true), time.Second)
	return true
}

func (this *Callback) OnClose(c *tcp.Conn) {
	fmt.Println("OnClose:", c.GetExtraData())
}

func main() {
	conf := Config.GetConfig()

	tcpPort := conf.String("tcp.port")

	// init models
	models.Init()

	//set cpus for max
	runtime.GOMAXPROCS(runtime.NumCPU())

	// creates a tcp listener
	tcpAddr, err := net.ResolveTCPAddr("tcp4", tcpPort)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	// creates a server
	config := &tcp.Config{
		PacketSendChanLimit:    20,
		PacketReceiveChanLimit: 20,
	}
	srv := tcp.NewServer(config, &Callback{}, &tcp.EchoProtocol{})

	// starts service
	go srv.Start(listener, time.Second)
	fmt.Println("listening:", listener.Addr())

	// catchs system signal
	chSig := make(chan os.Signal)
	signal.Notify(chSig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("Signal: ", <-chSig)

	// stops service
	srv.Stop()
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
