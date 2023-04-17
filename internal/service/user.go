package service

import (
	"golang.org/x/crypto/bcrypt"
)

type UserRequest struct {
	UserName string `form:"user_name" binding:"required"`
	Password string `form:"password" binding:"required"`
}

// 检查登录
func (svc *Service) CheckLogin(param *UserRequest) bool {
	_, hash := svc.dao.IsExsitUser(param.UserName)
	if hash != "" {
		bool := ComparePassword(param.Password, hash)
		return bool
	}

	return false
}

// 检查注册
func (svc *Service) CheckRegister(param *UserRequest) (bool, error) {
	bool := svc.dao.IsRegistered(param.UserName)

	if (bool == false) {
		hash, err := GeneratePassword(param.Password)

		if err != nil {
			return false, err // 注册失败
		}

		err = svc.dao.Register(param.UserName, string(hash))

		if err != nil {
			return false, err // 注册失败
		}

		return false, nil // 注册成功
	} else {
		return true, nil // 用户已注册
	}
}

// 生成密码
func GeneratePassword(password string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return hash, err
}

// 比较密码
func ComparePassword(password string, hash string) (bool) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false
	} else {
		return true
	}
}