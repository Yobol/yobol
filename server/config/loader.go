package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var Config *Server

// environment variables > config defaults

func Init() {
	v := viper.New()
	v.SetConfigName("config") // name of config file (without extension)
	v.SetConfigType("yaml")   // REQUIRED if the config does not have the extension in the name

	v.AddConfigPath("./etc/")
	v.AddConfigPath("$HOME/.yobol/server/")
	v.AddConfigPath("/etc/yobol/server/")

	if err := v.ReadConfig(); err != nil {
		panic(fmt.Errorf("Fatal error read config file: %s\n", err))
	}
	if err := v.Unmarshal(Config); err != nil {
		panic(fmt.Errorf("Fatal error unmarshal config file: %s\n", err))
	}
}
