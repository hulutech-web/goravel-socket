package socket

import (
	"github.com/goravel/framework/contracts/foundation"
	"goravel/packages/socket/routers"
	"goravel/packages/socket/servers"
)

const Binding = "socket"

var App foundation.Application

type ServiceProvider struct {
}

func (receiver *ServiceProvider) Register(app foundation.Application) {
	App = app

	app.Bind(Binding, func(app foundation.Application) (any, error) {
		return nil, nil
	})
}

func (receiver *ServiceProvider) Boot(app foundation.Application) {
	routers.Init()
	go servers.WriteMessage()
}
