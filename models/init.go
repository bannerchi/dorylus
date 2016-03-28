package models

import (
	//"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"

	Config "github.com/bannerchi/dorylus/util/config"
)

var engine *xorm.Engine

func Init() {
	var err error
	conf := Config.GetConfig()

	engine, err = xorm.NewEngine("mysql", conf.String("mysql.conn"))

	if err != nil {
		log.Fatalf("Fail to create engine: %v\n", err)
	}

	// sync struct to table ,double way
	if err = engine.Sync(new(TaskLog), new(Task)); err != nil {
		log.Fatalf("Fail to sync database: %v\n", err)
	}
}

func TableName(name string) string {
	conf := Config.GetConfig()
	tablePrefix := conf.String("mysql.prefix")
	return tablePrefix + name
}
