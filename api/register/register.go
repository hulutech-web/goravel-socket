package register

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/hulutech-web/goravel-socket/api"
	"github.com/hulutech-web/goravel-socket/define/retcode"
	"github.com/hulutech-web/goravel-socket/servers"
)

type RegisterController struct {
	//Dependent services
}

func NewRegisterController() *RegisterController {
	return &RegisterController{
		//Inject services
	}
}

type inputData struct {
	SystemId string `json:"systemId" form:"systemId" validate:"required"`
}

func (c *RegisterController) Run(ctx http.Context) http.Response {
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

	if err = servers.Register(inputData.SystemId); err != nil {
		ctx.Request().AbortWithStatusJson(retcode.FAIL, http.Json{
			"msg":  "error",
			"fail": err.Error(),
		})
		return nil
	}

	return ctx.Response().Success().Json(http.Json{
		"msg":  "success",
		"data": []string{},
	})
}
