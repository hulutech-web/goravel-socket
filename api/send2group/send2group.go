package send2group

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/hulutech-web/goravel-socket/api"
	"github.com/hulutech-web/goravel-socket/define/retcode"
	"github.com/hulutech-web/goravel-socket/servers"
)

type Send2GroupController struct {
	//Dependent services
}

func NewSend2GroupController() *Send2GroupController {
	return &Send2GroupController{
		//Inject services
	}
}

type inputData struct {
	SendUserId string `json:"sendUserId"`
	GroupName  string `json:"groupName" validate:"required"`
	Code       int    `json:"code"`
	Msg        string `json:"msg"`
	Data       string `json:"data"`
}

func (c *Send2GroupController) Run(ctx http.Context) http.Response {
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
	messageId := servers.SendMessage2Group(systemId, inputData.SendUserId, inputData.GroupName, inputData.Code, inputData.Msg, &inputData.Data)

	return ctx.Response().Success().Json(http.Json{
		"msg": "发送成功",
		"data": map[string]string{
			"messageId": messageId,
		},
	})
}
