package jobs

import (
	"log"
	"net"
	"os/exec"
	"time"

	"github.com/bannerchi/dorylus/models"
)

func InitOneJobByTaskId(taskId int) string {
	task, err := models.GetTaskById(taskId)
	if err != nil {
		log.Printf("InitJob error : %v", err.Error())
	}

	job, err := NewJobFromTask(task)

	if err != nil {
		log.Printf("InitJob error : %v", err.Error())
	}

	return AddJob(task.CronSpec, job)
}

func GetLocalIp() string {
	var ips string
	sliceIp := []string{}

	addrs, err := net.InterfaceAddrs()

	if err != nil {
		log.Fatal(err)
	}

	for _, address := range addrs {
		// check loop ip address
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				sliceIp = append(sliceIp, ipnet.IP.String())
			}
		}
	}

	for index, ip := range sliceIp {
		if len(sliceIp) > index+1 {
			ips = ips + ip + "|"
		} else {
			ips = ips + ip
		}
	}

	return ips
}

func runCmdWithTimeout(cmd *exec.Cmd, timeout time.Duration) (error, bool) {
	done := make(chan error)
	go func() {
		done <- cmd.Wait()
	}()

	var err error
	select {
	case <-time.After(timeout):
		log.Printf("Task run over%d sec，process will be killed: %d", int(timeout/time.Second), cmd.Process.Pid)
		go func() {
			<-done // 读出上面的goroutine数据，避免阻塞导致无法退出
		}()
		if err = cmd.Process.Kill(); err != nil {
			log.Printf("Process can't be killed: %d, errMsg: %s", cmd.Process.Pid, err)
		}
		return err, true
	case err = <-done:
		return err, false
	}
}
