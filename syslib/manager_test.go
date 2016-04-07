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
	tmpProcStatus := GetProcStatusByPid(108300)
	Convey("get error when process not exsit", t, func() {
		So(tmpProcStatus, ShouldEqual, "Process pid:108300 is not exsit")
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