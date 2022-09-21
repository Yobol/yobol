package config

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/viper"
)

var C *Config

// environment variables > config defaults

func Init() {
	C = new(Config)

	v := viper.New()
	v.SetConfigName("config") // name of config file (without extension)
	v.SetConfigType("yaml")   // REQUIRED if the config does not have the extension in the name
	v.AddConfigPath("./etc/")
	v.AddConfigPath("$HOME/.yobol/server/")
	v.AddConfigPath("/etc/yobol/server/")
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error read config file: %s", err))
	}
	if err := v.Unmarshal(C); err != nil {
		panic(fmt.Errorf("fatal error unmarshal config file: %s", err))
	}
	if data, err := json.MarshalIndent(C, "", "    "); err != nil {
		panic(fmt.Errorf("fatal error marshal config file with indent: %s", err))
	} else {
		fmt.Println(string(data))
	}
}
