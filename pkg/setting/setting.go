package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

var (
	// Cfg Config File
	Cfg *ini.File

	// RunMode app run mode
	RunMode string

	// HTTPPort http port
	HTTPPort int
	// ReadTimeout read timeout duration
	ReadTimeout time.Duration
	// WriteTimeout write timeout duration
	WriteTimeout time.Duration
	// PageSize defautl pagesize
	PageSize int
	// JwtSecret jwt verify secret
	JwtSecret string
)

// load config
func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatal(2, "Fail to parse 'conf/app.ini': %v", err)
	}

	loadBase()
	loadServer()
	loadApp()
}

func loadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func loadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatal(2, " Fail to get section 'server': ", err)
	}

	HTTPPort = sec.Key("HTTP_PORT").MustInt(3000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func loadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatal(2, "Fail to get section 'app': %v", err)
	}

	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}
