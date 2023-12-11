package getonlinelist

import (
	"github.com/goravel/framework/contracts/http"
	"goravel/packages/socket/api"
	"goravel/packages/socket/define/retcode"
	"goravel/packages/socket/servers"
)

type GetOnlineController struct {
	//Dependent services
}

func NewGetOnlineController() *GetOnlineController {
	return &GetOnlineController{
		//Inject services
	}
}

type inputData struct {
	GroupName string      `json:"groupName" validate:"required"`
	Code      int         `json:"code"`
	Msg       string      `json:"msg"`
	Data      interface{} `json:"data"`
}

func (c *GetOnlineController) Run(ctx http.Context) http.Response {
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

	ret := servers.GetOnlineList(&systemId, &inputData.GroupName)
	return ctx.Response().Success().Json(http.Json{
		"msg":  "发送成功",
		"data": ret,
	})
}
