package closeclient

import (
	"github.com/goravel/framework/contracts/http"
	"goravel/packages/socket/api"
	"goravel/packages/socket/define/retcode"
	"goravel/packages/socket/servers"
)

type CloseClientController struct {
	//Dependent services
}

func NewCloseClientController() *CloseClientController {
	return &CloseClientController{
		//Inject services
	}
}

type inputData struct {
	ClientId string `json:"clientId" validate:"required"`
}

func (c *CloseClientController) Run(ctx http.Context) http.Response {
	var inputData inputData
	if err1 := ctx.Request().Bind(&inputData); err1 != nil {
		ctx.Request().AbortWithStatusJson(retcode.FAIL, http.Json{
			"msg":  "error",
			"fail": err1.Error(),
		})
		return nil

	}
	err := api.Validate(inputData)
	if err != nil {
		ctx.Request().AbortWithStatusJson(retcode.FAIL, http.Json{
			"msg":  "error",
			"fail": err.Error(),
		})
		return nil

	}
	systemId := ctx.Request().Header("SystemId", "")

	//发送信息
	servers.CloseClient(inputData.ClientId, systemId)
	return ctx.Response().Success().Json(http.Json{
		"data": "success",
		"msg":  "退出成功",
	})
}
