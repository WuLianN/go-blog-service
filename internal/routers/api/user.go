package api

import (
	// "github.com/WuLianN/go-blog-service/global"
	"github.com/WuLianN/go-blog-service/internal/service"
	"github.com/WuLianN/go-blog-service/pkg/app"
	"github.com/WuLianN/go-blog-service/pkg/errcode"
	"github.com/gin-gonic/gin"
)

// 登录
func Login(c *gin.Context) {
	param := service.UserRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		// global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	bool := svc.CheckLogin(&param)

	if bool != true {
		return
	}

	token, err := app.GenerateToken(param.UserName, "")
	if err != nil {
		// global.Logger.Errorf(c, "app.GenerateToken err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		return
	}

	response.ToResponse(gin.H{
		"code": errcode.Success.Code(),
		"msg": errcode.Success.Msg(),
		"data": gin.H{ "token": token },
	})
}

// 注册
func Register(c *gin.Context) {
	param := service.UserRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		// global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	bool, err := svc.CheckRegister(&param)

	if bool == true {
		response.ToResponse(gin.H{
			"code": errcode.Success.Code(),
			"msg": "用户已注册",
		})
	} else if bool == false && err == nil {
		response.ToResponse(gin.H{
			"code": errcode.Success.Code(),
			"msg": "注册成功",
		})
	} else {
		response.ToResponse(gin.H{
			"code": errcode.Success.Code(),
			"msg": "注册失败",
		})
	}
}