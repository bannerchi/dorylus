package models

import (
	"errors"
)

type Task struct {
	Id           int
	UserId       int
	GroupId      int
	Pid          int
	RunServer    string
	TaskName     string
	TaskType     int
	Description  string
	CronSpec     string
	Concurrent   int
	Command      string
	Status       int
	Notify       int
	NotifyEmail  string
	Timeout      int
	ExecuteTimes int
	PrevTime     int64
	CreateTime   int64
}

const (
	TASK_SUCCESS = 0  // 任务执行成功
	TASK_ERROR   = -1 // 任务执行出错
	TASK_TIMEOUT = -2 // 任务执行超时
)

func (t *Task) TableName() string {
	return TableName("task")
}

func GetTaskById(id int) (*Task, error) {
	t := new(Task)
	has, err := engine.Id(id).Get(t)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, errors.New("Task does not exist")
	}

	return t, nil
}

func GetTaskByPid(pid int) (*Task, error) {
	t := new(Task)
	has, err := engine.Where("pid=?", pid).Get(t)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, errors.New("Task does not exist")
	}

	return t, nil
}

func UpdateTask(id int, t *Task) error {
	if _, err := engine.Id(id).Update(t); err != nil {
		return err
	}
	return nil
}
