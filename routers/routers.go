package routers

import (
	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"
	"goravel/app/http/middleware"
	"goravel/packages/socket/api/bind2group"
	"goravel/packages/socket/api/closeclient"
	"goravel/packages/socket/api/getallgroup"
	"goravel/packages/socket/api/getonlinelist"
	"goravel/packages/socket/api/register"
	"goravel/packages/socket/api/send2client"
	"goravel/packages/socket/api/send2clients"
	"goravel/packages/socket/api/send2group"
)

func Init() {
	facades.Route().Prefix("/api").Middleware(middleware.Jwt()).Group(func(router route.Router) {
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
		//@袁浩 新增：获取所有分组
		router.Post("/get_all_groups", getAllGroupHandler.Run)
	})
}
