package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"regexp"
	"runtime"
	"strconv"
	"syscall"
	"time"

	"github.com/bannerchi/dorylus/models"
	"github.com/bannerchi/dorylus/syslib"
	"github.com/bannerchi/dorylus/tcp"
	"github.com/bannerchi/dorylus/util"
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

	//get load average
	regexpGetLoadAverage, _ := regexp.Compile("get_load_average")

	//get proc status by pid
	regexpGetProcStatus, _ := regexp.Compile(`^get_proc_status_-?\d+$`)

	regexpGetNumber, _ := regexp.Compile(`-?\d+$`)

	// run job from task
	regexpRunJob, _ := regexp.Compile(`^run_task_-?\d+$`)

	// remove job by taskid
	regexpRmjob, _ := regexp.Compile(`^rm_task_-?\d+$`)

	if regexpGetLoadAverage.MatchString(req) {
		resvMsg = tcp.NewEchoPacket([]byte(syslib.GetLoadAverage()), false)
	}

	if regexpGetProcStatus.MatchString(req) {
		var pid int
		strPid := regexpGetNumber.FindString(req)
		pid, _ = strconv.Atoi(strPid)
		resvMsg = tcp.NewEchoPacket([]byte(syslib.GetProcStatusByPid(pid)), false)
	}

	if regexpRunJob.MatchString(req) {
		var taskId int
		strTaskId := regexpGetNumber.FindString(req)
		taskId, _ = strconv.Atoi(strTaskId)
		resvMsg = tcp.NewEchoPacket([]byte(syslib.RunTask(taskId)), false)
	}

	if regexpRmjob.MatchString(req) {
		var taskId int
		strTaskId := regexpGetNumber.FindString(req)
		taskId, _ = strconv.Atoi(strTaskId)
		resvMsg = tcp.NewEchoPacket([]byte(syslib.RmTaskById(taskId)), false)
	}

	c.AsyncWritePacket(tcp.NewEchoPacket(resvMsg.Serialize(), true), time.Second)
	return true
}

func (this *Callback) OnClose(c *tcp.Conn) {
	fmt.Println("OnClose:", c.GetExtraData())
}

func main() {
	conf := util.GetConfig()

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
