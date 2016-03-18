package jobs

import (
	"github.com/bannerchi/dorylus/antcron"
	"log"
	"sync"
)

var (
	mainCron *antcron.Cron
	workPool chan bool
	lock     sync.Mutex
)

func init() {
	if size, _ := GetConfig().Int("WorkPollSize"); size > 0 {
		workPool = make(chan bool, size)
	}
	mainCron = antcron.New()
	mainCron.Start()
}

func AddJob(spec string, job *Job) bool {
	lock.Lock()
	defer lock.Unlock()

	if GetEntryById(job.id) != nil {
		return false
	}
	err := mainCron.AddJob(spec, job)
	if err != nil {
		log.Println("AddJob: ", err.Error())
		return false
	}
	return true
}

func RemoveJob(id int) {
	mainCron.RemoveJob(func(e *antcron.Entry) bool {
		if v, ok := e.Job.(*Job); ok {
			if v.id == id {
				return true
			}
		}
		return false
	})
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
