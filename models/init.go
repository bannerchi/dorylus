package models

import (
	//"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
)

var engine *xorm.Engine

func Init(dbConf string) {
	var err error
	//TODO add config
	engine, err = xorm.NewEngine("mysql", dbConf)

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
