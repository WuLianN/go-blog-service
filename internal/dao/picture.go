package dao

import (
	"github.com/WuLianN/go-blog-service/internal/model"
	"github.com/WuLianN/go-blog-service/pkg/app"
)

func (d *Dao) GetPictureList(name string, state uint8, page, pageSize int) ([]*model.Picture, error) {
	picture := model.Picture{Name: name, State: state}
	pageOffset := app.GetPageOffset(page, pageSize)
	return picture.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) CountPicture(name string, state uint8) (int, error) {
	picture := model.Picture{Name: name, State: state}
	return picture.Count(d.engine)
}