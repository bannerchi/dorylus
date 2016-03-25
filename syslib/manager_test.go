package syslib

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestNewSysinfo(t *testing.T) {
	sys := new(Sysinfo)
	Convey("new sysinfo", t, func() {
		So(newSysinfo("ls"), ShouldHaveSameTypeAs, sys)
	})
}

func TestGetLoadAverage(t *testing.T) {
	Convey("get loadaverage", t, func() {
		So(GetLoadAverage(), ShouldNotBeBlank)
	})
}

func TestGetProcStatusByPid(t *testing.T) {
	Convey("get process status by pid", t, func() {
		So(GetProcStatusByPid(1083), ShouldNotBeBlank)
	})
}
