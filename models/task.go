package models

import (
	"errors"
)

type Task struct {
	Id           int
	UserId       int
	GroupId      int
	Pid          int
	ServerId     int
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
	UpdatedAt    int64 `xorm:"updated"`
}

const (
	TASK_SUCCESS = 0
	TASK_ERROR   = -1
	TASK_TIMEOUT = -2
)

func (t *Task) TableName() string {
	return TableName("task")
}

func GetTaskById(id int) (*Task, error) {
	t := &Task{}
	has, err := engine.Where("id=?", id).Get(t)

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
	_, err := GetTaskById(id)
	if err != nil {
		return err
	}
	_, err2 := engine.Update(t)

	if err2 != nil {
		return err2
	}

	return nil
}
