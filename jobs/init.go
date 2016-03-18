package jobs

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
	"github.com/astaxie/beego/config"
)
var conf config.Configer
//func InitJobs() {
//	list, _ := models.TaskGetList(1, 1000000, "status", 1)
//	for _, task := range list {
//		job, err := NewJobFromTask(task)
//		if err != nil {
//			beego.Error("InitJobs:", err.Error())
//			continue
//		}
//		AddJob(task.CronSpec, job)
//	}
//}

func init() {
	var err error
	env := os.Getenv("DORYLUS_ENV")
	if env == "dev" || env == "" {
		env = "dev"
	}
	conf, err = config.NewConfig("ini", "conf/" + env + ".conf")
	if err != nil {
		log.Fatal("config error")
	}
}

func GetConfig() config.Configer{
	return conf
}

func runCmdWithTimeout(cmd *exec.Cmd, timeout time.Duration) (error, bool) {
	done := make(chan error)
	go func() {
		done <- cmd.Wait()
	}()

	var err error
	select {
	case <-time.After(timeout):
	//beego.Warn(fmt.Sprintf("任务执行时间超过%d秒，进程将被强制杀掉: %d", int(timeout/time.Second), cmd.Process.Pid))
		log.Printf("任务执行时间超过%d秒，进程将被强制杀掉: %d", int(timeout / time.Second), cmd.Process.Pid)
		go func() {
			<-done // 读出上面的goroutine数据，避免阻塞导致无法退出
		}()
		if err = cmd.Process.Kill(); err != nil {
			log.Printf("进程无法杀掉: %d, 错误信息: %s", cmd.Process.Pid, err)
			//beego.Error(fmt.Sprintf("进程无法杀掉: %d, 错误信息: %s", cmd.Process.Pid, err))
		}
		return err, true
	case err = <-done:
		return err, false
	}
}
