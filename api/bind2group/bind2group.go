package bind2group

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/hulutech-web/goravel-socket/api"
	"github.com/hulutech-web/goravel-socket/define/retcode"
	"github.com/hulutech-web/goravel-socket/servers"
)

type Bind2GroupController struct {
	//Dependent services
}

func NewBind2GroupController() *Bind2GroupController {
	return &Bind2GroupController{
		//Inject services
	}
}

type inputData struct {
	ClientId  string `json:"clientId" validate:"required"`
	GroupName string `json:"groupName" validate:"required"`
	UserId    string `json:"userId"`
	Extend    string `json:"extend"` // 拓展字段，方便业务存储数据
}

func (c *Bind2GroupController) Run(ctx http.Context) http.Response {
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
	servers.AddClient2Group(systemId, inputData.GroupName, inputData.ClientId, inputData.UserId, inputData.Extend)

	return ctx.Response().Success().Json(http.Json{
		"data": "success",
		"msg":  "发送成功",
	})
}
