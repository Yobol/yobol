package config

type Config struct {
	Server *Server `json:"server" yaml:"server"`
	MySQL  *MySQL  `json:"mysql" yaml:"mysql"`
}

type Server struct {
	Host string `json:"host" yaml:"host"`
	Port uint16 `json:"port" yaml:"port"`
}

type MySQL struct {
	Host string `json:"host" yaml:"host"`
	Port uint16 `json:"port" yaml:"port"`
}
