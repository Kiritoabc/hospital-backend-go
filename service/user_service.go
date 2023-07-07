package service

import (
	"gorm.io/gorm"
	"hospital-backend-go/dto"
	"hospital-backend-go/model"
)

// IsPhoneExists 检查电话是否已经注册过了
func IsPhoneExists(db *gorm.DB, phone string) bool {
	var doctor model.Doctor
	db.Where("phone = ?", phone).First(&doctor)
	return doctor.Id != 0
}

// AddUser 添加用户
func AddUser(db *gorm.DB, user model.User) error {
	result := db.Create(&user)

	return result.Error
}

func AuthUser(db *gorm.DB, loginUser dto.UserLoginDto) (bool, error) {
	// 主要是判断密码书否正确
	var user model.User
	result := db.Where("phone = ?", loginUser.Phone).First(&user)
	if result.Error != nil {
		return false, result.Error
	}
	if user.Password != loginUser.Password {
		return false, nil
	}
	return true, nil
}
