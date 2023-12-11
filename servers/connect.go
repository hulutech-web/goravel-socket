package servers

import (
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"goravel/packages/socket/api"
	"goravel/packages/socket/define/retcode"
	"goravel/packages/socket/tools/util"
	nethttp "net/http"

	"github.com/goravel/framework/contracts/http"
)

const (
	// 最大的消息大小
	maxMessageSize = 8192
)

type Controller struct {
}

type renderData struct {
	ClientId string `json:"clientId"`
}

func (c *Controller) Run(ctx http.Context) http.Response {
	conn, err := (&websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		// 允许所有CORS跨域请求
		CheckOrigin: func(r *nethttp.Request) bool {
			return true
		},
	}).Upgrade(ctx.Response().Writer(), ctx.Request().Origin(), nil)
	if err != nil {
		log.Errorf("upgrade error: %v", err)
		nethttp.NotFound(ctx.Response().Writer(), ctx.Request().Origin())
		return nil
	}

	//设置读取消息大小上线
	conn.SetReadLimit(maxMessageSize)

	//解析参数
	//systemId := r.FormValue("systemId")
	systemId := ctx.Request().Input("systemId")
	if len(systemId) == 0 {
		_ = Render(conn, "", "", retcode.SYSTEM_ID_ERROR, "系统ID不能为空", []string{})
		_ = conn.Close()
		return nil
	}

	clientId := util.GenClientId()

	clientSocket := NewClient(clientId, systemId, conn)

	Manager.AddClient2SystemClient(systemId, clientSocket)

	//读取客户端消息
	clientSocket.Read()

	if err = api.ConnRender(conn, renderData{ClientId: clientId}); err != nil {
		_ = conn.Close()
		return nil
	}

	// 用户连接事件
	Manager.Connect <- clientSocket

	return nil

}
