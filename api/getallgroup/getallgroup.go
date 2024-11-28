package getallgroup

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/hulutech-web/goravel-socket/api"
	"github.com/hulutech-web/goravel-socket/define/retcode"
	"github.com/hulutech-web/goravel-socket/servers"
)

type GetAllGroupController struct {
	//Dependent services
}

func NewGetAllGroupController() *GetAllGroupController {
	return &GetAllGroupController{
		//Inject services
	}
}

type inputData struct {
	ClientId string `json:"clientId" validate:"required"`
	UserId   string `json:"userId"`
	Extend   string `json:"extend"` // 拓展字段，方便业务存储数据
}

func (c *GetAllGroupController) Run(ctx http.Context) http.Response {
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
	groups := servers.GetAllGroups()

	return ctx.Response().Success().Json(http.Json{
		"msg":  "退出成功",
		"data": groups,
	})
}
