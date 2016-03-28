package jobs

import (
	"fmt"
	"log"
	"sync"

	"github.com/bannerchi/dorylus/antcron"
	Config "github.com/bannerchi/dorylus/util/config"
)

var (
	mainCron *antcron.Cron
	workPool chan bool
	lock     sync.Mutex
)

func init() {
	if size, _ := Config.GetConfig().Int("WorkPollSize"); size > 0 {
		workPool = make(chan bool, size)
	}
	mainCron = antcron.New()
	mainCron.Start()
}

func AddJob(spec string, job *Job) string {
	lock.Lock()
	defer lock.Unlock()

	if GetEntryById(job.id) != nil {
		return "job is already exsit"
	}
	err := mainCron.AddJob(spec, job)
	if err != nil {
		log.Println("AddJob: ", err.Error())
		return fmt.Sprintf("AddJob error: %s", err)
	}
	return fmt.Sprintf("AddJob %s success", job.GetName())
}

func RemoveJob(id int) bool {
	var isSuccess bool
	mainCron.RemoveJob(func(e *antcron.Entry) bool {
		if v, ok := e.Job.(*Job); ok {
			if v.id == id {
				isSuccess = true
				return true
			}
		}
		return false
	})
	return isSuccess
}

func GetEntryById(id int) *antcron.Entry {
	entries := mainCron.Entries()
	for _, e := range entries {
		if v, ok := e.Job.(*Job); ok {
			if v.id == id {
				return e
			}
		}
	}
	return nil
}

func GetEntries(size int) []*antcron.Entry {
	ret := mainCron.Entries()
	if len(ret) > size {
		return ret[:size]
	}
	return ret
}
