package core

import (
	"bons-server/src/initialize"
	"github.com/fvbock/endless"
)

func Server()  {
	Router := initialize.Routers()
	server := endless.NewServer("ss",Router)
	server.ReadTimeout = 10000
	server.WriteTimeout = 10000
	server.MaxHeaderBytes = 1 << 20
	server.ListenAndServe().Error()
}