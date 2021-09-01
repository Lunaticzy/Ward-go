package common

import (
	"fmt"
	"gopkg.in/ini.v1"
	"log"
)

func init() {
	cfg, err := ini.Load("ward.ini")
	if err != nil {
		log.Fatalln(err)
	}
	section, err := cfg.GetSection("")
	if err != nil {
		log.Fatalln(err)
	}
	Config.ServerName = section.Key("server_name").String()
	Config.Theme = section.Key("theme").String()
	Config.Version = section.Key("version").String()
	Config.Port, _ = section.Key("port").Int()
	runLocal, _ := section.Key("run_local").Bool()
	if runLocal {
		Config.Host = "127.0.0.1"
	} else {
		Config.Host = "0.0.0.0"
	}

	Config.Addr = fmt.Sprintf("%s:%d", Config.Host, Config.Port)

}

type ConfigStruct struct {
	Version    string
	Theme      string
	ServerName string
	Host       string
	Port       int
	Addr       string
}

var Config = ConfigStruct{}
