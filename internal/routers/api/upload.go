package api

import (
	"github.com/gin-gonic/gin"
	// "github.com/WuLianN/go-blog/global"
	"github.com/WuLianN/go-blog/internal/service"
	"github.com/WuLianN/go-blog/pkg/app"
	"github.com/WuLianN/go-blog/pkg/convert"
	"github.com/WuLianN/go-blog/pkg/errcode"
	"github.com/WuLianN/go-blog/pkg/upload"
)

type Upload struct{}

func NewUpload() Upload {
	return Upload{}
}

func (u Upload) UploadFile(c *gin.Context) {
	response := app.NewResponse(c)
	file, fileHeader, err := c.Request.FormFile("file")
	fileType := convert.StrTo(c.PostForm("type")).MustInt()
	
	if err != nil {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Error()))
		return 
	}

	if fileHeader == nil || fileType <= 0 {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}

	svc := service.New(c.Request.Context())
	fileInfo, err := svc.UploadFile(upload.FileType(fileType), file, fileHeader)

	if err != nil {
		// global.Logger.Errorf("svc.UploadFile err: %v", err)
		response.ToErrorResponse(errcode.ERROR_UPLOAD_FILE_FAIL.WithDetails(err.Error()))
		return
	}

	response.ToResponse(gin.H{
		"file_access_url": fileInfo.AccessUrl,
	})
}