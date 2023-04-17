package dao

import (
	"github.com/WuLianN/go-blog-service/internal/model"
)

// 用户是否已存在
func (d *Dao) IsExsitUser (UserName string) (bool, string) {
	user := model.User{ UserName: UserName }
	err := d.engine.Where("user_name = ?", UserName).First(&user).Error
	if err != nil {
		return false, ""
	}
	return true, user.Password
}

// 用户是否已注册
func (d *Dao) IsRegistered(UserName string) (bool) {
	bool, _ := d.IsExsitUser(UserName)

	return bool
}

// 注册
func (d *Dao) Register(UserName string, Password string) (error) {
	user := model.User{ UserName: UserName, Password: Password }
	err := d.engine.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}