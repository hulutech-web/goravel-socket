package send2clients

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/hulutech-web/goravel-socket/api"
	"github.com/hulutech-web/goravel-socket/define/retcode"
	"github.com/hulutech-web/goravel-socket/servers"
)

type Send2ClientsController struct {
	//Dependent services
}

func NewSend2ClientsController() *Send2ClientsController {
	return &Send2ClientsController{
		//Inject services
	}
}

type inputData struct {
	ClientIds  []string `json:"clientIds" validate:"required"`
	SendUserId string   `json:"sendUserId"`
	Code       int      `json:"code"`
	Msg        string   `json:"msg"`
	Data       string   `json:"data"`
}

func (c *Send2ClientsController) Run(ctx http.Context) http.Response {
	var inputData inputData
	if err1 := ctx.Request().Bind(&inputData); err1 != nil {
		ctx.Request().AbortWithStatusJson(http.StatusBadRequest, http.Json{
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

	for _, clientId := range inputData.ClientIds {
		//发送信息
		_ = servers.SendMessage2Client(clientId, inputData.SendUserId, inputData.Code, inputData.Msg, &inputData.Data)
	}

	return ctx.Response().Success().Json(http.Json{
		"msg":  "success",
		"data": "",
	})
}
