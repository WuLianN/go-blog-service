package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/WuLianN/go-blog/global"
	"github.com/WuLianN/go-blog/internal/service"
	"github.com/WuLianN/go-blog/pkg/app"
	// "github.com/WuLianN/go-blog/pkg/convert"
	"github.com/WuLianN/go-blog/pkg/errcode"
)

type Picture struct{}

func NewPicture() Picture {
	return Picture{}
}

func (p Picture) List(c *gin.Context) {
	param := service.PictureListRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	totalRows, err := svc.CountPicture(&service.CountPictureRequest{Name: param.Name, State: param.State})
	if err != nil {
		global.Logger.Errorf(c, "svc.CountPicture err: %v", err)
		response.ToErrorResponse(errcode.ErrorCountPictureFail)
		return
	}
	pictures, err := svc.GetPictureList(&param, &pager)
	if err != nil {
		global.Logger.Errorf(c, "svc.GetPictureList err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetPictureListFail)
		return
	}

	response.ToResponseList(pictures, totalRows)
	return
}