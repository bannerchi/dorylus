package models

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGetTaskById(t *testing.T) {
	//init mysql
	Init()
	var ts *Task

	Convey("Get task by id", t, func() {
		res1, _ := GetTaskById(1)
		So(res1, ShouldHaveSameTypeAs, ts)
	})
	Convey("Get task by no this id", t, func() {
		res2, err := GetTaskById(0)
		So(res2, ShouldBeNil)
		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldEqual, "Task does not exist")
	})
}

func TestGetTaskByPId(t *testing.T) {
	//init mysql
	Init()
	var ts *Task

	Convey("Get task by pid", t, func() {
		res1, _ := GetTaskById(20178)
		So(res1, ShouldHaveSameTypeAs, ts)
	})
	Convey("Get task by no this pid", t, func() {
		res2, err := GetTaskById(0)
		So(res2, ShouldBeNil)
		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldEqual, "Task does not exist")
	})
}

func TestUpdateTask(t *testing.T) {
	Init()
	ts := &Task{TaskName: "test_test_succ"}
	Convey("Update task success", t, func() {
		So(UpdateTask(1, ts), ShouldBeNil)
		res1, _ := GetTaskById(1)
		So(res1.TaskName, ShouldEqual, "test_test_succ")
	})
	ts2 := &Task{TaskName: "test_fail"}
	Convey("Update task faild", t, func() {
		err := UpdateTask(999, ts2)
		fmt.Println(err)
		So(err, ShouldNotBeNil)
	})
}
