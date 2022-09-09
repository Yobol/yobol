package config

type Server struct {
	MySQL *MySQL `json:"mysql" yaml:"mysql"`
}

type MySQL struct {
	Host string `json:"host" yaml:"host"`
	Port uint16 `json:"port" yaml:"port"`
}
