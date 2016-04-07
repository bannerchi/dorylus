package syslib

import (
	"encoding/json"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"github.com/bannerchi/dorylus/antcron"
)


func TestGetLoadAverage(t *testing.T) {
	Convey("get loadaverage", t, func() {
		So(GetLoadAverage(), ShouldNotBeEmpty)
	})
}

func TestGetProcStatusByPid(t *testing.T) {
	var procStatusT = ProcessState{}
	var psr = new(ProcessState)
	tmpProcStatus := GetProcStatusByPid(1083)
	json.Unmarshal([]byte(tmpProcStatus), psr)
	Convey("get process status by pid", t, func() {
		So(tmpProcStatus, ShouldNotBeBlank)
		So(psr,ShouldHaveSameTypeAs, &procStatusT)
	})
}

func TestGetMemory(t *testing.T) {
	Convey("get memory", t, func() {
		So(GetMemory(), ShouldNotBeEmpty)
	})
}

func TestGetReadToRunJobs(t *testing.T) {
	var jobs = []*antcron.Entry{}
	json.Unmarshal(GetReadToRunJobs(1), jobs)

	Convey("get ready to run jobs", t, func() {
		So(jobs, ShouldBeEmpty)
	})
}