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

// SelectDoctorByPhone 根据phone查询医生
func SelectDoctorByPhone(db *gorm.DB, phone string) (model.Doctor, error) {
	var doctor = model.Doctor{}
	result := db.Where("phone = ?", phone).First(&doctor)

	return doctor, result.Error
}
