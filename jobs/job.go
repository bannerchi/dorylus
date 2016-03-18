package jobs

import (
	"bytes"
	"fmt"
	//"html/template"
	"os/exec"
	"runtime/debug"
	"github.com/bannerchi/dorylus/models"
	//"strings"
	"sync"
	"log"
	"time"
)

//var mailTpl *template.Template

//func init() {
//	mailTpl, _ = template.New("mail_tpl").Parse(`
//	你好 {{.username}}，<br/>
//
//<p>以下是任务执行结果：</p>
//
//<p>
//任务ID：{{.task_id}}<br/>
//任务名称：{{.task_name}}<br/>
//执行时间：{{.start_time}}<br />
//执行耗时：{{.process_time}}秒<br />
//执行状态：{{.status}}
//</p>
//<p>-------------以下是任务执行输出-------------</p>
//<p>{{.output}}</p>
//<p>
//--------------------------------------------<br />
//本邮件由系统自动发出，请勿回复<br />
//如果要取消邮件通知，请登录到系统进行设置<br />
//</p>
//`)
//
//}

type Job struct {
	id         int
	logId      int64
	name       string
	task       *models.Task
	runFunc    func(time.Duration) (string, string, error, bool)
	running    sync.Mutex
	status     int
	Concurrent bool
}

func NewJobFromTask(task *models.Task) (*Job, error) {
	if task.Id < 1 {
		return nil, fmt.Errorf("ToJob: 缺少id")
	}
	job := NewCommandJob(task.Id, task.TaskName, task.Command)
	job.task = task
	job.Concurrent = task.Concurrent == 0
	return job, nil
}

func NewCommandJob(id int, name string, command string) *Job {
	job := &Job{
		id:   id,
		name: name,
	}
	job.runFunc = func(timeout time.Duration) (string, string, error, bool) {
		bufOut := new(bytes.Buffer)
		bufErr := new(bytes.Buffer)
		cmd := exec.Command("/bin/bash", "-c", command)
		cmd.Stdout = bufOut
		cmd.Stderr = bufErr
		cmd.Start()
		err, isTimeout := runCmdWithTimeout(cmd, timeout)

		return bufOut.String(), bufErr.String(), err, isTimeout
	}
	return job
}

func (j *Job) Status() int {
	return j.status
}

func (j *Job) GetName() string {
	return j.name
}

func (j *Job) GetId() int {
	return j.id
}

func (j *Job) GetLogId() int64 {
	return j.logId
}

func (j *Job) Run() {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err, "\n", string(debug.Stack()))
		}
	}()

	t := time.Now()

	if j.Concurrent {
		j.running.Lock()
		defer j.running.Unlock()
	}

	if workPool != nil {
		workPool <- true
		defer func() {
			<-workPool
		}()
	}

	j.status = 1
	defer func() {
		j.status = 0
	}()

	timeout := time.Duration(time.Hour * 24)
	if j.task.Timeout > 0 {
		timeout = time.Second * time.Duration(j.task.Timeout)
	}

	cmdOut, cmdErr, err, isTimeout := j.runFunc(timeout)

	ut := time.Now().Sub(t) / time.Millisecond

	// 插入日志
	tasklogs := new(models.TaskLog)
	tasklogs.TaskId = j.id
	tasklogs.Output = cmdOut
	tasklogs.Error = cmdErr
	tasklogs.ProcessTime = int(ut)
	tasklogs.CreateTime = t.Unix()

	if isTimeout {
		tasklogs.Status = models.TASK_TIMEOUT
		tasklogs.Error = fmt.Sprintf("任务执行超过 %d 秒\n----------------------\n%s\n", int(timeout/time.Second), cmdErr)
	} else if err != nil {
		tasklogs.Status = models.TASK_ERROR
		tasklogs.Error = err.Error() + ":" + cmdErr
	}
	j.logId, _ = models.TaskLogAdd(tasklogs)

	// 更新上次执行时间
	j.task.PrevTime = t.Unix()
	j.task.ExecuteTimes++
	models.UpdateTask(j.task.Id, j.task)

	// 发送邮件通知
//	if (j.task.Notify == 1 && err != nil) || j.task.Notify == 2 {
//		user, uerr := models.UserGetById(j.task.UserId)
//		if uerr != nil {
//			return
//		}
//
//		var title string
//
//		data := make(map[string]interface{})
//		data["task_id"] = j.task.Id
//		data["username"] = user.UserName
//		data["task_name"] = j.task.TaskName
//		data["start_time"] = beego.Date(t, "Y-m-d H:i:s")
//		data["process_time"] = float64(ut) / 1000
//		data["output"] = cmdOut
//
//		if isTimeout {
//			title = fmt.Sprintf("任务执行结果通知 #%d: %s", j.task.Id, "超时")
//			data["status"] = fmt.Sprintf("超时（%d秒）", int(timeout/time.Second))
//		} else if err != nil {
//			title = fmt.Sprintf("任务执行结果通知 #%d: %s", j.task.Id, "失败")
//			data["status"] = "失败（" + err.Error() + "）"
//		} else {
//			title = fmt.Sprintf("任务执行结果通知 #%d: %s", j.task.Id, "成功")
//			data["status"] = "成功"
//		}
//
//		content := new(bytes.Buffer)
//		mailTpl.Execute(content, data)
//		ccList := make([]string, 0)
//		if j.task.NotifyEmail != "" {
//			ccList = strings.Split(j.task.NotifyEmail, "\n")
//		}
//		if !mail.SendMail(user.Email, user.UserName, title, content.String(), ccList) {
//			beego.Error("发送邮件超时：", user.Email)
//		}
//	}
}
