package facades

import (
	socket "github.com/hulutech-web/goravel-socket"
	"github.com/hulutech-web/goravel-socket/contracts"
	"log"
)

func Socket() contracts.Socket {
	instance, err := socket.App.Make(socket.Binding)
	if err != nil {
		log.Println(err)
		return nil
	}

	return instance.(contracts.Socket)
}
