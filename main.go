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

	"github.com/bannerchi/dorylus/syslib"
	"github.com/bannerchi/dorylus/tcp"
	//"github.com/bannerchi/dorylus/models"
	"github.com/astaxie/beego/config"
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
	req := string(echoPacket.GetBody())

	var resvMsg = new(tcp.EchoPacket)
	switch req {
	case "get_load_average":
		resvMsg = tcp.NewEchoPacket([]byte(syslib.GetLoadAverage()), false)
	}

	//fmt.Printf("OnMessage:[%v] [%v]\n", echoPacket.GetLength(), req)
	c.AsyncWritePacket(tcp.NewEchoPacket(resvMsg.Serialize(), true), time.Second)
	return true
}

func (this *Callback) OnClose(c *tcp.Conn) {
	fmt.Println("OnClose:", c.GetExtraData())
}

func main() {
	var err error
	env := os.Getenv("DORYLUS_ENV")
	if env == "dev" || env == "" {
		env = "dev"
	}
	conf, err := config.NewConfig("ini", "conf/"+env+".conf")

	if err != nil {
		log.Fatal(err)
	}

	tcpPort := conf.String("tcp.port")
	fmt.Println(tcpPort)
	//models.Init()

	// task_log := new(models.TaskLog)
	// task_log.TaskId = 1
	// task_log.ProcessTime = 20
	// task_log.Output = "sas"
	// task_log.CreateTime = time.Now().Unix()
	// models.TaskLogAdd(task_log)

	// task := new(models.Task)
	// task.UserId = 212
	// error := models.UpdateTask(2, task)
	// if error != nil {
	// 	fmt.Println(error)
	// } else {
	// 	fmt.Println("jobs done")
	// }

	runtime.GOMAXPROCS(runtime.NumCPU())

	// creates a tcp listener
	tcpAddr, err := net.ResolveTCPAddr("tcp4", ":8989")
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	// creates a server
	config := &tcp.Config{
		PacketSendChanLimit:    20,
		PacketReceiveChanLimit: 20,
	}
	srv := tcp.NewServer(config, &Callback{}, &tcp.EchoProtocol{})

	// status := syslib.GetLoadAverage()

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
