package util

import (
	"regexp"
	"strconv"

	"github.com/bannerchi/dorylus/syslib"
)

const (
	LOAD_AVERAGE     = "get_load_average"
	MEMORY           = "get_memory"
	PROC_STATUS      = `^get_proc_status_-?\d+$`
	NUMBER           = `-?\d+$`
	RUN_JOB          = `^run_task_-?\d+$`
	REMOVE_JOB       = `^rm_task_-?\d+$`
	READY_TO_RUN_JOB = `^ready_to_run_jobs_-?\d+$`
)

func ResponsFilter(req string) []byte {
	var resvMsg []byte
	//get load average
	regexpGetLoadAverage, _ := regexp.Compile(LOAD_AVERAGE)

	//get memory
	regexpGetMemory, _ := regexp.Compile(MEMORY)

	//get proc status by pid
	regexpGetProcStatus, _ := regexp.Compile(PROC_STATUS)

	regexpGetNumber, _ := regexp.Compile(NUMBER)

	// run job from task
	regexpRunJob, _ := regexp.Compile(RUN_JOB)

	// remove job by taskid
	regexpRmjob, _ := regexp.Compile(REMOVE_JOB)

	// remove job by taskid
	regexpGetReadToRunJobs, _ := regexp.Compile(READY_TO_RUN_JOB)

	if regexpGetLoadAverage.MatchString(req) {
		resvMsg = syslib.GetLoadAverage()
	}

	if regexpGetMemory.MatchString(req) {
		resvMsg = syslib.GetMemory()
	}

	if regexpGetProcStatus.MatchString(req) {
		var pidi32 int32
		strPid := regexpGetNumber.FindString(req)
		pid, _ := strconv.Atoi(strPid)
		pidi32 = int32(pid)
		resvMsg = []byte(syslib.GetProcStatusByPid(pidi32))
	}

	if regexpRunJob.MatchString(req) {
		var taskId int
		strTaskId := regexpGetNumber.FindString(req)
		taskId, _ = strconv.Atoi(strTaskId)
		resvMsg = []byte(syslib.RunTask(taskId))
	}

	if regexpRmjob.MatchString(req) {
		var taskId int
		strTaskId := regexpGetNumber.FindString(req)
		taskId, _ = strconv.Atoi(strTaskId)
		resvMsg = []byte(syslib.RmTaskById(taskId))
	}

	if regexpGetReadToRunJobs.MatchString(req) {
		var size int
		strSize := regexpGetNumber.FindString(req)
		size, _ = strconv.Atoi(strSize)
		resvMsg = syslib.GetReadToRunJobs(size)
	}

	return resvMsg
}
