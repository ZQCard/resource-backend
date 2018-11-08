package config

import (
	"github.com/Unknwon/goconfig"
	"resource-backend/pkg/logging"
	"strings"
	"time"
)

type Server struct {
	RunMode      string
	HTTPPort     int
	ReadTimeOut  time.Duration
	WriteTimeOut time.Duration
}

var ServerSettings = &Server{}

type Database struct {
	Type     string
	User     string
	Password string
	Host     string
	Name     string
}

var DatabaseSetting = &Database{}

type App struct {
	BaseUrl  string
	PageSize string
	PageNum  string
	JwtSecret string
	ImageSavePath string
	ImageAllowExts []string
	ImageMaxSize int
}

var AppSetting = &App{}

var cfg *goconfig.ConfigFile

func init() {
	var err error
	cfg, err = goconfig.LoadConfigFile("conf/app.ini", "conf/local.ini")
	if err != nil {
		logging.Fatal("加载配置文件出错 : %s", err)
	}
	// 读取服务器配置
	ServerSettings.RunMode, err = cfg.GetValue("server", "RunMode")
	if err != nil {
		logging.Fatal("读取键值出错(%s) : %s", "RunMode", err)
	}
	ServerSettings.HTTPPort = cfg.MustInt("server", "HTTPPort")
	ServerSettings.ReadTimeOut = time.Duration(cfg.MustInt("server", "ReadTimeOut", 60)) * time.Second
	ServerSettings.WriteTimeOut = time.Duration(cfg.MustInt("server", "WriteTimeOut", 60)) * time.Second

	// 读取数据库配置
	DatabaseSetting.Type, err = cfg.GetValue("database", "Type")
	if err != nil {
		logging.Fatal("读取键值出错(%s) : %s", "database.Type", err)
	}
	DatabaseSetting.User, err = cfg.GetValue("database", "User")
	if err != nil {
		logging.Fatal("读取键值出错(%s) : %s", "database.User", err)
	}
	DatabaseSetting.Password, err = cfg.GetValue("database", "Password")
	if err != nil {
		logging.Fatal("读取键值出错(%s) : %s", "database.Password", err)
	}
	DatabaseSetting.Host, err = cfg.GetValue("database", "Host")
	if err != nil {
		logging.Fatal("读取键值出错(%s) : %s", "database.Host", err)
	}
	DatabaseSetting.Name, err = cfg.GetValue("database", "Name")
	if err != nil {
		logging.Fatal("读取键值出错(%s) : %s", "database.Name", err)
	}

	// 读取应用配置
	AppSetting.PageSize, err = cfg.GetValue("app", "PageSize")
	if err != nil {
		logging.Fatal("读取键值出错(%s) : %s", "app.PageSize", err)
	}

	AppSetting.PageNum, err = cfg.GetValue("app", "PageNum")
	if err != nil {
		logging.Fatal("读取键值出错(%s) : %s", "app.PageNum", err)
	}

	AppSetting.BaseUrl, err = cfg.GetValue("app", "BaseUrl")
	if err != nil {
		logging.Fatal("读取键值出错(%s) : %s", "app.PageNum", err)
	}

	AppSetting.JwtSecret, err = cfg.GetValue("app", "JwtSecret")
	if err != nil {
		logging.Fatal("读取键值出错(%s) : %s", "app.JwtSecret", err)
	}

	AppSetting.ImageSavePath, err = cfg.GetValue("app", "ImageSavePath")
	if err != nil {
		logging.Fatal("读取键值出错(%s) : %s", "app.ImageSavePath", err)
	}

	allowExts, err := cfg.GetValue("app", "ImageAllowExts")
	if err != nil {
		logging.Fatal("读取键值出错(%s) : %s", "app.ImageAllowExts", err)
	}
	AppSetting.ImageAllowExts = strings.Split(allowExts, ",")

	AppSetting.ImageMaxSize = cfg.MustInt("app", "ImageMaxSize")
	AppSetting.ImageMaxSize = AppSetting.ImageMaxSize * 1024 * 1024
}

func GetConfigParam(section string, key string) (param string) {
	param, err  := cfg.GetValue(section, key)
	if err != nil {
		logging.Fatal("读取键值出错(%s) : %s", key, err)
	}
	return
}
