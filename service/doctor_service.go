package service

import (
	"gorm.io/gorm"
	"hospital-backend-go/model"
)

// AddDoctor 添加医生
func AddDoctor(db *gorm.DB, doctor model.Doctor) error {
	result := db.Create(&doctor)

	return result.Error
}
