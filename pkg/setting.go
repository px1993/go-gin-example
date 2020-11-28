package setting

import (
	"log"
	"time"

	"gopkg.in/ini.v1"
)

var (
	Cfg     *ini.File
	RunMode string
	err     error

	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	PageSize  int
	JwtSecret string
)

func init() {
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}
	LoadBase()
	LoadServer()
	LoadApp()
}

//load base
func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

//load server
func LoadServer() {
	serv, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to parse 'server': %v", err)
	}
	HttpPort = serv.Key("HTTP_PORT").MustInt(8080)
	ReadTimeout = time.Duration(serv.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(serv.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

//load app
func LoadApp() {
	app, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to parse 'app': %v", err)
	}
	PageSize = app.Key("PAGE_SIZE").MustInt(10)
	JwtSecret = app.Key("JWT_SECRETE").MustString("!@)*#)!@U#@*!@!)")
}
