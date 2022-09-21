package config

import "time"

type Config struct {
	Server *Server `json:"server" yaml:"server"`

	MySQL *MySQL `json:"mysql" yaml:"mysql"`
}

type Server struct {
	Host string `json:"host" yaml:"host"`
	Port uint16 `json:"port" yaml:"port"`

	Auth *Auth `json:"auth" yaml:"auth"`
}

type Auth struct {
	TokenName   string        `mapstructure:"token_name" json:"token_name" yaml:"token_name"`
	TokenMaxAge time.Duration `mapstructure:"token_max_age" json:"token_max_age" yaml:"token_max_age"`
	JWTSecret   string        `mapstructure:"jwt_secret" json:"jwt_secret" yaml:"jwt_secret"`
}

type MySQL struct {
	Host string `json:"host" yaml:"host"`
	Port uint16 `json:"port" yaml:"port"`
}
