package models

import (
	//"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
)

var engine *xorm.Engine

func Init() {
	var err error
	//TODO add config
	engine, err = xorm.NewEngine("mysql", "root:556213@/webcron?charset=utf8")

	if err != nil {
		log.Fatalf("Fail to create engine: %v\n", err)
	}

	// 同步结构体与数据表
	if err = engine.Sync(new(TaskLog), new(Task)); err != nil {
		log.Fatalf("Fail to sync database: %v\n", err)
	}

}

func TableName(name string) string {
	//TODO add config
	return "t_" + name
}
