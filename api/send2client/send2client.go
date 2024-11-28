package send2client

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/hulutech-web/goravel-socket/api"
	"github.com/hulutech-web/goravel-socket/define/retcode"
	"github.com/hulutech-web/goravel-socket/servers"
)

type Send2ClientController struct {
	//Dependent services
}

func NewRegisterController() *Send2ClientController {
	return &Send2ClientController{
		//Inject services
	}
}

type inputData struct {
	ClientId   string `json:"clientId" validate:"required"`
	SendUserId string `json:"sendUserId"`
	Code       int    `json:"code"`
	Msg        string `json:"msg"`
	Data       string `json:"data"`
}

func (c *Send2ClientController) Run(ctx http.Context) http.Response {
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

	//发送信息
	messageId := servers.SendMessage2Client(inputData.ClientId, inputData.SendUserId, inputData.Code, inputData.Msg, &inputData.Data)

	return ctx.Response().Success().Json(http.Json{
		"msg": "success",
		"data": []map[string]string{{
			"messageId": messageId,
		}},
	})
}
