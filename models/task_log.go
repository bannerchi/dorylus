package models

type TaskLog struct {
	Id          int
	TaskId      int
	Server      string
	Output      string
	Error       string
	Status      int
	ProcessTime int
	CreateTime  int64
}

func (tlog *TaskLog) TableName() string {
	return TableName("task_log")
}

func TaskLogAdd(tlog *TaskLog) (int64, error) {
	return engine.Insert(tlog)
}

func TaskLogDelByTaskId(taskId int) (int64, error) {
	return engine.Delete(&TaskLog{TaskId: taskId})
}
