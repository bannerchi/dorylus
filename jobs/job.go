package jobs

import (
	"bytes"
	"fmt"
	"github.com/bannerchi/dorylus/models"
	"os/exec"
	"runtime/debug"
//"strings"
	"log"
	"time"
)

type Job struct {
	id         int
	logId      int64
	name       string
	task       *models.Task
	runFunc    func(time.Duration) (string, string, error, bool)
	status     int // >0 is running
	Concurrent bool
	pid        int
}

func NewJobFromTask(task *models.Task) (*Job, error) {
	if task.Id < 1 {
		return nil, fmt.Errorf("ToJob: missing task_id")
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
		job.task.Pid = cmd.Process.Pid
		job.pid = cmd.Process.Pid
		err, isTimeout := runCmdWithTimeout(cmd, timeout)
		log.Printf("start %d ,pid: %d\n", job.id, job.pid)
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

func (j *Job) GetPid() int {
	return j.pid
}

func (j *Job) Run() {

	if !j.Concurrent && j.status > 0 {
		log.Printf("任务[%d]上一次执行尚未结束，本次被忽略。\n", j.id)
		return
	}

	defer func() {
		if err := recover(); err != nil {
			log.Println(err, "\n", string(debug.Stack()))
		}
	}()



	if workPool != nil {
		workPool <- true
		defer func() {
			<-workPool
		}()
	}

	j.status++
	defer func() {
		j.status--
	}()

	t := time.Now()
	timeout := time.Duration(time.Hour * 24)
	if j.task.Timeout > 0 {
		timeout = time.Second * time.Duration(j.task.Timeout)
	}

	cmdOut, cmdErr, err, isTimeout := j.runFunc(timeout)

	ut := time.Now().Sub(t) / time.Millisecond

	//TODO add file mode
	// add logs
	tasklogs := new(models.TaskLog)
	tasklogs.TaskId = j.id
	tasklogs.Server = GetLocalIp()
	tasklogs.Output = cmdOut
	tasklogs.Error = cmdErr
	tasklogs.ProcessTime = int(ut)
	tasklogs.CreateTime = t.Unix()

	if isTimeout {
		tasklogs.Status = models.TASK_TIMEOUT
		tasklogs.Error = fmt.Sprintf("Task run over %d sec\n----------------------\n%s\n", int(timeout / time.Second), cmdErr)
	} else if err != nil {
		tasklogs.Status = models.TASK_ERROR
		tasklogs.Error = err.Error() + ":" + cmdErr
	}
	j.logId, _ = models.TaskLogAdd(tasklogs)

	// update prev run time
	j.task.PrevTime = t.Unix()
	j.task.ExecuteTimes = j.task.ExecuteTimes + 1
	j.task.RunServer = GetLocalIp()

	models.UpdateTask(j.task.Id, j.task)
}
