package conf

/**
    @date: 2024/7/13
**/

import (
	"fmt"
	"strings"

	"github.com/txxzx/goMemorandum/user/model"
	"github.com/usthooz/gconf"
)

type Config struct {
	Mysql struct {
		DbUser     string `yaml:"db_user"`
		DbPassword string `yaml:"db_password"`
		DbName     string `yaml:"db_name"`
		DbHost     string `yaml:"db_host"`
		DbPort     string `yaml:"db_port"`
	} `yaml:"mysql"`

	Service struct {
		AppModel string `yaml:"app_model"`
		HttpPort string `yaml:"http_port"`
	} `yaml:"service"`
}

func Init() {
	var (
		cfg Config
	)

	// 读取配置
	ozconf := gconf.NewConf(&gconf.Gconf{
		ConfPath: "./conf.yaml",
	})
	if err := ozconf.GetConf(&cfg); err != nil {
		fmt.Errorf("配置文件读取错误，请检查文件路径: %v", err)
		return
	}

	path := strings.Join([]string{cfg.Mysql.DbUser, ":", cfg.Mysql.DbPassword, "@tcp(", cfg.Mysql.DbHost, ":", cfg.Mysql.DbPort, ")/", cfg.Mysql.DbName, "?charset=utf8&parseTime=true"}, "")
	model.Database(path)
}
