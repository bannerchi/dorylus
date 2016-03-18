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

const (
	TASK_SUCCESS = 0  // 任务执行成功
	TASK_ERROR   = -1 // 任务执行出错
	TASK_TIMEOUT = -2 // 任务执行超时
)

func (t *Task) TableName() string {
	return TableName("task")
}

func UpdateTask(id int, t *Task) error {
	if _, err := engine.Id(id).Update(t); err != nil {
		return err
	}
	return nil
}
