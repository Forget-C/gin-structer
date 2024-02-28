package config

import (
	"os"
	"strings"

	"github.com/spf13/viper"

	"github.com/Forget-C/http-structer/internal/dao/sqldb"
)

const defaultConfigPath = "configs"

var Cfg = new(Config)

type Config struct {
	Listen string `yaml:"listen"`
	Mysql  *sqldb.ConOptions
}

func Parse(filepath string) {
	if filepath != "" {
		viper.SetConfigFile(filepath)
	} else {
		viper.AddConfigPath(defaultConfigPath)
		if os.Getenv("SERVER_ENV") != "" {
			viper.SetConfigName(os.Getenv("SERVER_ENV"))
		} else {
			viper.SetConfigName("config")
		}
		viper.SetConfigType("yaml")
	}
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(Cfg)
	if err != nil {
		panic(err)
	}
}

func ModifyHttpHost(s string) string {
	if strings.HasPrefix(s, "http://") || strings.HasPrefix(s, "https://") {
		return s
	}
	return "http://" + s
}
