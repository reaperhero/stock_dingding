package config

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var Config config

type config struct {
	LogLevel string
	MySQL    MySQLConfig
	Web      WebConfig
}

type WebConfig struct {
	Port int
}

type MySQLConfig struct {
	IP       string
	Port     int
	User     string
	Password string
	Database string
}

func init() {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("toml")
	v.AddConfigPath(".")
	v.AddConfigPath("cmd/")
	err := v.ReadInConfig()
	if err != nil {
		log.Fatal("read config failed: %v", err)
	}

	err = v.Unmarshal(&Config)
	if err != nil {
		log.Fatal(err)
	}
	configValue, _ := json.Marshal(Config)
	fmt.Println("run config: ", string(configValue))
}
