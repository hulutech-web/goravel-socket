package facades

import (
	"goravel/packages/socket"
	"goravel/packages/socket/contracts"
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
