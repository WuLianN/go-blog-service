package model

import (
	"gorm.io/gorm"
	"github.com/WuLianN/go-blog/global"
)

type Picture struct {
	*Model
	Name string `json:"name"`
	State uint8 `json:"state"`
	Url string `json:"url"`
}

func (p Picture) TableName() string {
	return "blog_picture"
}

func (p Picture) Count(db *gorm.DB) (int, error) {
	var count int64
	if p.Name != "" {
		db = db.Where("name = ?", p.Name)
	}
	db = db.Where("state = ?", p.State)

	if err := db.Model(&p).Where("is_del = ?", 0).Count(&count).Error; err != nil {
		return 0, err
	}

	convertCount := int(count)

	return convertCount, nil
}

func (p Picture) List(db *gorm.DB, pageOffset, pageSize int) ([]*Picture, error) {
	var pictures []*Picture
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if p.Name != "" {
		db = db.Where("name = ?", p.Name)
	}
	db = db.Where("state = ?", p.State)
	if err = db.Where("is_del = ?", 0).Find(&pictures).Error; err != nil {
		return nil, err
	}

	for _, v := range pictures {
		v.Url = global.AppSetting.UploadServerUrl + "/" + v.Name
	}

	return pictures, nil
}