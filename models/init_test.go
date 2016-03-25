package models

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestTableName(t *testing.T) {
	Convey("get Table name right", t, func() {
		So(TableName("ss"), ShouldEqual, "t_ss")
	})
}
