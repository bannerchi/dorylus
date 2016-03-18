package syslib

import (
	//"bytes"
	//"errors"
	//"fmt"
	"log"
	"os/exec"
	"strings"
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
	//fmt.Printf("%#v", spliceRes[1])
	return spliceRes[1]
}
