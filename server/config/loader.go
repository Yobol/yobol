package config

import "github.com/spf13/viper"

// environment variables > config defaults

func Init() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config does not have the extension in the name

	viper.AddConfigPath("./etc/")
	viper.AddConfigPath("$HOME/.yobol/server/")
	viper.AddConfigPath("/etc/yobol/server/")
}
