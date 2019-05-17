package setting

import (
	"log"
	"os"
	"time"

	"github.com/go-ini/ini"
)

type App struct {
	JwtSecret string
	PageSize  int
	PrefixUrl string

	RuntimeRootPath string

	ImageSavePath  string
	ImageMaxSize   int
	ImageAllowExts []string

	ExportSavePath string
	QrCodeSavePath string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string
}

// AppSetting app configration
var AppSetting = &App{}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// ServerSetting server configuration
var ServerSetting = &Server{}

type Database struct {
	Dialect     string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

// DatabaseSetting database configuration
var DatabaseSetting = &Database{}

type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

// RedisSetting redis configuration
var RedisSetting = &Redis{}

var cfg *ini.File

// Setup initialize the configuration
func Setup() {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("[setting.Setup] Fail to parse 'conf/app.ini': %v", err)
	}

	mapTo("app", AppSetting)
	mapTo("server", ServerSetting)
	mapTo("database", DatabaseSetting)
	mapTo("redis", RedisSetting)

	AppSetting.ImageMaxSize = AppSetting.ImageMaxSize * 1024 * 1024 // MB
	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
	host := os.Getenv("DB_HOST") // Get Value from env
	if host != "" {
		DatabaseSetting.Host = host
	}
	RedisSetting.IdleTimeout = RedisSetting.IdleTimeout * time.Second
}

// mapTo maps data sources to given struct with error log.Fatalf().
func mapTo(sec string, v interface{}) {
	err := cfg.Section(sec).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %v section err: %v", sec, err)
	}
}
