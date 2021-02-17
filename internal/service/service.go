package service

import (
	"context"

	// otgorm "github.com/eddycjy/opentracing-gorm"

	"github.com/WuLianN/go-blog/global"
	"github.com/WuLianN/go-blog/internal/dao"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}

	// 链路追踪之 SQL 追踪 -> otgorm gorm 版本冲突
	// svc.dao = dao.New(otgorm.WithContext(svc.ctx, global.DBEngine))

	svc.dao = dao.New(global.DBEngine)
	return svc
}
