package routers

import (
	"net/http"
	"time"

	"github.com/WuLianN/go-blog-service/global"
	"github.com/WuLianN/go-blog-service/internal/middleware"
	"github.com/WuLianN/go-blog-service/internal/routers/api"
	v1 "github.com/WuLianN/go-blog-service/internal/routers/api/v1"
	"github.com/WuLianN/go-blog-service/pkg/limiter"
	"github.com/gin-gonic/gin"
)

var methodLimiters = limiter.NewMethodLimiter().AddBuckets(
	limiter.LimiterBucketRule{
		Key:          "/auth", // 自定义键值对名称
		FillInterval: time.Second, // 间隔多久时间放 Quantum 个令牌
		Capacity:     10, // 令牌桶的容量
		Quantum:      10, // 每次到达间隔时间后所放的具体令牌数量
	},
)


func SetupRouter() *gin.Engine {
  r := gin.Default()

	// 跨域
	r.Use(middleware.Cors())

	// 访问日志
	// r.Use(middleware.AccessLog())
	// 链路追踪
	// r.Use(middleware.Tracing())
	// 接口限流控制
	r.Use(middleware.RateLimiter(methodLimiters))
	// 统一超时管理
	r.Use(middleware.ContextTimeout(global.AppSetting.DefaultContextTimeout))
     
	upload := api.NewUpload()
  picture := v1.NewPicture()

  r.POST("/upload/file", upload.UploadFile)
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))

	r.POST("/login", api.Login)
	r.POST("/register", api.Register)
    
	apiv1 := r.Group("api/v1")
	apiv1.Use(middleware.JWT())
	{
		apiv1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})

		apiv1.GET("/pictures", picture.List)
	}
	
  return r
}