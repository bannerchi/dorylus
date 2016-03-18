package dorylus

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
)

type Sysinfo struct {
	command string
	bufout string
	buferr string
	pid    int
}

func newSysinfo(command string) *Sysinfo {
	return &Sysinfo{
		command : command
	}
}

func (s *Sysinfo) runCommand() *Sysinfo {
	s.bufout = new(bytes.Buffer)
	s.bufout = new(bytes.Buffer)
	cmd := exec.Command("/bin/bash", "-c", s.command)
	cmd.Stdout = s.bufout
	cmd.Stderr = s.buferr
	cmd.Start()

	s.pid = cmd.Process.Pid
	return s
}

func GetLoadAverage() string {
	sys := newSysinfo("uptime")
	result := sys.runCommand()
	fmt.Println(result.bufout)

	return result.bufout.String()
}
