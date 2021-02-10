package model

import (
	// "github.com/WuLianN/go-blog/pkg/app"
	"gorm.io/gorm"
)

type Picture struct {
	*Model
	Name string `json:"name"`
	State uint8 `json:"state"`
}

func (p Picture) TableName() string {
	return "blog_picture"
}

func (p Picture) Count(db *gorm.DB) (int, error) {
	var count int
	if p.Name != "" {
		db = db.Where("name = ?", p.Name)
	}
	db = db.Where("state = ?", p.State)
	if err := db.Model(&p).Where("is_del = ?", 0).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
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

	return pictures, nil
}