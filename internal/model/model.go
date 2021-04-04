package model

import (
	"fmt"
	"time"
	"github.com/WuLianN/go-blog/global"
	"github.com/WuLianN/go-blog/pkg/setting"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

const (
	STATE_OPEN  = 1
	STATE_CLOSE = 0
)

type Model struct {
	ID         uint32 `gorm:"primary_key" json:"id"`
	// IsDel      uint8  `json:"is_del"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime,
	)
	
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		
	})


	if err != nil {
		return nil, err
	}

	if global.ServerSetting.RunMode == "debug" {
		
	}
	return db, nil
}

