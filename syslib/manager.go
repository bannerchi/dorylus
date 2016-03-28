package syslib

import (
	"encoding/json"
	. "fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/bannerchi/dorylus/jobs"
)

type Sysinfo struct {
	command string
	stdout  string
}

func newSysinfo(comm string) *Sysinfo {
	return &Sysinfo{
		command: comm,
	}
}

func (s *Sysinfo) runCommand() *Sysinfo {
	out, err := exec.Command("/bin/bash", "-c", s.command).Output()
	if err != nil {
		log.Fatal(err)
	}
	s.stdout = string(out)
	return s
}

func GetLoadAverage() string {
	sys := newSysinfo("uptime")
	result := sys.runCommand()
	spliceRes := strings.Split(result.stdout, "load average:")
	return spliceRes[1]
}

func GetProcStatusByPid(pid int) string {
	sys := newSysinfo(Sprintf("cat /proc/%d/status", pid))
	result := sys.runCommand()

	return result.stdout
}

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

// get entries
func GetReadToRunJobs(size int) []byte {
	arrEntry := jobs.GetEntries(size)
	jsonArrEntry, _ := json.Marshal(arrEntry)
	return jsonArrEntry
}
