package config

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

var Conf *Config

type Config struct {
	Env     string `yaml:"Env"`     // 环境：prod、dev
	Port    string `yaml:"Port"`    // 端口
	LogFile string `yaml:"LogFile"` // 日志文件

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
		OutputFile string `yaml:"OutputFile"`
	} `yaml:"Logrus"`
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
