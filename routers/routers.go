package routers

import (
	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"
	"github.com/hulutech-web/goravel-socket/api/bind2group"
	"github.com/hulutech-web/goravel-socket/api/closeclient"
	"github.com/hulutech-web/goravel-socket/api/getallgroup"
	"github.com/hulutech-web/goravel-socket/api/getonlinelist"
	"github.com/hulutech-web/goravel-socket/api/register"
	"github.com/hulutech-web/goravel-socket/api/send2client"
	"github.com/hulutech-web/goravel-socket/api/send2clients"
	"github.com/hulutech-web/goravel-socket/api/send2group"
)

func Init() {
	facades.Route().Prefix("/api").Middleware(Jwt()).Group(func(router route.Router) {
		registerController := register.NewRegisterController()
		sendToClientController := send2client.NewRegisterController()
		sendToClientsController := send2clients.NewSend2ClientsController()
		sendToGroupController := send2group.NewSend2GroupController()
		bindToGroupController := bind2group.NewBind2GroupController()
		getOnlinelistController := getonlinelist.NewGetOnlineController()
		closeClientController := closeclient.NewCloseClientController()
		getAllGroupHandler := getallgroup.NewGetAllGroupController()

		router.Post("/register", registerController.Run)
		router.Post("/send_to_client", sendToClientController.Run)
		router.Post("/send_to_clients", sendToClientsController.Run)
		router.Post("/send_to_group", sendToGroupController.Run)
		router.Post("/bind_to_group", bindToGroupController.Run)
		router.Post("/get_online_list", getOnlinelistController.Run)
		router.Post("/close_client", closeClientController.Run)
		//新增：获取所有分组
		router.Post("/get_all_groups", getAllGroupHandler.Run)
	})
}
