package util

import (
	"log"
	"os"

	"github.com/astaxie/beego/config"
)

func GetConfig() config.Configer {
	env := os.Getenv("DORYLUS_ENV")
	if env == "dev" || env == "" {
		env = "dev"
	}

	conf, err := config.NewConfig("ini", "conf/"+env+".conf")
	if err != nil {
		conf2, err2 := config.NewConfig("ini", "../conf/"+env+".conf")
		if err2 != nil {
			log.Fatal("config err:" + err2.Error())
		} else {
			conf = conf2
		}
	}
	return conf
}
