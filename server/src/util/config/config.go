package config

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

var Conf *Config

// 系统告警邮箱信息
type systemEmail struct {
	Host     string `yaml:"Host"`
	Port     string `yaml:"Port"`
	Username string `yaml:"Username"`
	Password string `yaml:"Password"`
}

type Config struct {
	Env     string `yaml:"Env"`  // 环境：prod、dev
	Port    string `yaml:"Port"` // 端口
	AppName string `yaml:"AppName"`
	AppMode string `yaml:"AppMode"` // debug or release

	// MD5 密钥
	MD5 struct {
		SignSecret string `yaml:"SignSecret"`
		SignExpiry int    `yaml:"SignExpiry"`
	} `yaml:"MD5"`

	// 数据库连接配置
	MySql struct {
		Dialect   string `yaml:"Dialect"`
		Host      string `yaml:"Host"`
		Port      string `yaml:"Port"`
		Username  string `yaml:"Username"`
		Password  string `yaml:"Password"`
		Database  string `yaml:"Database"`
		MaxIdle   int    `yaml:"MaxIdle"`
		MaxActive int    `yaml:"MaxActive"`
		Charset   string `yaml:"Charset"`
		ShowSql   bool   `yaml:"ShowSql"`
	} `yaml:"MySql"`

	// 日志配置
	Logrus struct {
		AccessLogName string `yaml:"AccessLogName"`
		ErrorLogName  string `yaml:"ErrorLogName"`
		GrpcLogName   string `yaml:"GrpcLogName"`
	} `yaml:"Logrus"`

	// 告警通知
	ErrorNotify struct {
		Open        bool        `yaml:"Open"`
		NotifyUser  string      `yaml:"NotifyUser"`
		SystemEmail systemEmail `yaml:"SystemEmail"`
	} `yaml:"ErrorNotify"`

	// 链路追踪
	Jaeger struct {
		Open     bool   `yaml:"Open"`
		HostPort string `yaml:"HostPort"`
	} `yaml:"Jaeger"`

	// 超时时间
	Timeout struct {
		AppRead  int `yaml:"AppRead"`
		AppWrite int `yaml:"AppWrite"`
	} `yaml:"Timeout"`
}

func InitConfig(filename string) {
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		logrus.Error(err)
		return
	}

	Conf = &Config{}
	err = yaml.Unmarshal(yamlFile, Conf)
	if err != nil {
		logrus.Error(err)
	}
}
