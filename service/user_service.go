package service

import (
	"gorm.io/gorm"
	"hospital-backend-go/dto"
	"hospital-backend-go/model"
)

func Register(registerUser dto.UserRegisterDto) {

}

// IsPhoneExists 检查电话是否已经注册过了
func IsPhoneExists(db *gorm.DB, phone string) bool {
	var doctor model.Doctor
	result := db.Where("phone = ?", phone).First(&doctor)
	return result.Error == nil
}

// AddUser 添加用户
func AddUser(db *gorm.DB, user model.User) error {
	result := db.Create(&user)

	return result.Error
}
