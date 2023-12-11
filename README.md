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
