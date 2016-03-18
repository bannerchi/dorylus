package models

type Task struct {
	Id           int
	UserId       int
	GroupId      int
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

func (t *Task) TableName() string {
	return TableName("task")
}

func UpdateTask(id int, t *Task) error {
	if _, err := engine.Id(id).Update(t); err != nil {
		return err
	}
	return nil
}
