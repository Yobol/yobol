package server

import (
	"fmt"

	v1 "github.com/yobol/yobol/api/v1"
	"github.com/yobol/yobol/config"
)

type server interface {
	ListenAndServe() error
}

func Run() {
	address := fmt.Sprintf(":%d", config.C.Server.Port)
	fmt.Printf("server run on: %s\n", address)
	s := initServer(address, v1.MountApiGroups())
	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}
