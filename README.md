# Socket
#### 使用说明
#### 1、在goravel项目中的config目录下的app.go文件中的providers数组中添加
```go
import Socket "goravel/packages/socket"
```
在providers数组中添加
```go
	&Socket.ServiceProvider{},
```
#### 2、在goravel项目中的router目录下的web.go文件中添加
```go
import 	"goravel/packages/socket/servers"
func Web() {
websocketHandler := &servers.Controller{}
facades.Route().Get("/ws", websocketHandler.Run)
go servers.Manager.Start()
}
```
#### 3、访问路径/ws
[//]: # 业务流程()
1、注册大区，systemId,不能重复，随机值，需缓存
2、连接ws,需要携带systemId头信息，连接成功后返回clientId，缓存至本地（这个时候会出现：xxx上线了，xxx下线了），每上线一个人将该人的clientId存入本地缓存，方便随时根据client进行发送消息
3、绑定客户端到分组（房间），对战双方都需要绑定，其他业务逻辑自行研究

#### 4、集成操作
#### 4.1、创建一个包
```go
go run . artisan make:package socket
```
#### 4.2、使用包
将扩展覆盖到扩展包中
#### 4.3、路由说明 routers/routers.go
```go
facades.Route().Prefix("/api").Middleware(middleware.Jwt()).Group(func(router route.Router) {
		registerController := register.NewRegisterController()
		sendToClientController := send2client.NewRegisterController()
		sendToClientsController := send2clients.NewSend2ClientsController()
		sendToGroupController := send2group.NewSend2GroupController()
		bindToGroupController := bind2group.NewBind2GroupController()
		getOnlinelistController := getonlinelist.NewGetOnlineController()
		closeClientController := closeclient.NewCloseClientController()
		getAllGroupHandler := getallgroup.NewGetAllGroupController()

		router.Post("/register", registerController.Run) //注册大区
		router.Post("/send_to_client", sendToClientController.Run) //发送消息给指定的客户端
		router.Post("/send_to_clients", sendToClientsController.Run) //发送消息给指定的客户端
		router.Post("/send_to_group", sendToGroupController.Run) //发送消息给指定的分组
		router.Post("/bind_to_group", bindToGroupController.Run) //绑定客户端到分组
		router.Post("/get_online_list", getOnlinelistController.Run) //获取在线列表
		router.Post("/close_client", closeClientController.Run) //关闭客户端
		router.Post("/get_all_groups", getAllGroupHandler.Run) //获取所有分组
	})
```