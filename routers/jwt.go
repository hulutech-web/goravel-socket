package routers

import (
	"errors"
	"github.com/goravel/framework/auth"
	contractshttp "github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

func Jwt() contractshttp.Middleware {
	return func(ctx contractshttp.Context) {
		//获取header中的Authorization 为Bearer token
		token := ctx.Request().Header("Authorization", "")
		//如果token为空
		if token == "" {
			ctx.Request().AbortWithStatus(401)
			return
		}
		token = token[7:]
		//解析payload
		if _, err := facades.Auth(ctx).Parse(token); err != nil {
			is := errors.Is(err, auth.ErrorTokenExpired)
			if is {
				ctx.Request().AbortWithStatusJson(contractshttp.StatusUnauthorized, contractshttp.Json{
					"message": "登录已过期，请重新登录",
				})
				return
			}
		}
		ctx.Request().Next()
	}
}
