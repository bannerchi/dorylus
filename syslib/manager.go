package syslib

import (
	"encoding/json"
	. "fmt"
	"time"

	"github.com/bannerchi/dorylus/jobs"
	UtilMem "github.com/shirou/gopsutil/mem"
	UtilLoad "github.com/shirou/gopsutil/load"
	UtilProc "github.com/shirou/gopsutil/process"
)

type ProcessState struct {
	Pid           int32 `json:"pid"`
	IsRunning     bool `json:"is_running"`
	MemoryPercent float32 `json:"memory_percent"`
	CpuPercent    float64 `json:"cpu_percent"`
}

type RetRunJob struct {
	Tid 		int `json:"tid"`
	Status 		int `json:"status"`
	Name        string `json:"name"`
	Pid 		int `json:"pid"`
}

func GetLoadAverage() []byte {
	v, _ := UtilLoad.Avg()

	jsonArr, _ := json.Marshal(v)
	return jsonArr
}
/**
	@return {"is_running":true,"memory_percent":2.2123966,"cpu_percent":0.999388968588172}
 */
func GetProcStatusByPid(pid int32) string {
	if isExsit, _ := UtilProc.PidExists(pid); isExsit == false {
		return Sprintf("Process pid:%d is not exsit", pid)
	}

	processInfo := new(ProcessState)
	process, _ := UtilProc.NewProcess(pid)

	isRunning, _ := process.IsRunning()
	memoryPercent, _ := process.MemoryPercent()
	cpuPercent, _ := process.Percent(1 * time.Second)

	processInfo.Pid = pid
	processInfo.IsRunning = isRunning
	processInfo.MemoryPercent = memoryPercent
	processInfo.CpuPercent = cpuPercent

	jsonArr, _ := json.Marshal(processInfo)

	return string(jsonArr)
}
/**
 return AddJob error: %s OR AddJob %s success
 */
func RunTask(taskId int) string {
	return jobs.InitOneJobByTaskId(taskId)
}

func RmTaskById(taskId int) string {
	status := jobs.RemoveJob(taskId)
	if status {
		return "success"
	} else {
		return "faild"
	}
}

func GetMemory() []byte {
	v, _ := UtilMem.VirtualMemory()
	jsonArr, _ := json.Marshal(v)
	return jsonArr
}

// get entries
func GetReadToRunJobs(size int) []byte {
	arrEntry := jobs.GetEntries(size)
	arrRet := []*RetRunJob{}
	for _, e := range arrEntry {
		if v, ok := e.Job.(*jobs.Job); ok {
			arrRet = append(arrRet, &RetRunJob{
				Name: v.GetName(),
				Tid: v.GetId(),
				Status: v.Status(),
				Pid: v.GetPid(),
			})
		}
	}
	jsonArrEntry, _ := json.Marshal(arrRet)

	return jsonArrEntry
}
