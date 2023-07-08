package service

import (
	"gorm.io/gorm"
	"hospital-backend-go/database"
	"hospital-backend-go/model"
)

type DoctorService struct {
}

func (s *DoctorService) GetDoctorListCount() (count int, err error) {
	db := database.GetDB()
	var t int64
	if err := db.Model(&model.Doctor{}).Count(&t).Error; err != nil {
		return 0, err
	}
	count = int(t)
	return count, nil
}

func (s *DoctorService) GetDoctorList(name string, pageNum int, pageSize int,
	id int, idNum int, phone string, fee string, departId int, role int, rooms *[]model.Doctor) error {
	db := database.GetDB()
	tx := db.Model(&model.Doctor{})
	if name != "0" {
		tx = tx.Where("name = ?", name)
	}
	if id != 0 {
		tx = tx.Where("id = ?", id)
	}
	if departId != 0 {
		tx = tx.Where("depart_id = ?", departId)
	}
	if idNum != 0 {
		tx = tx.Where("idNum = ?", idNum)
	}
	if phone != "0" {
		tx = tx.Where("phone = ?", phone)
	}
	if role != -1 {
		tx = tx.Where("role = ?", role)
	}
	if fee != "0" {
		tx = tx.Where("fee <= ?", fee)
	}
	if err := tx.Offset((pageNum - 1) * pageSize).
		Limit(pageSize).Find(rooms).Error; err != nil {
		return err
	}
	return nil
}

func (s *DoctorService) DeleteById(id int) error {
	db := database.GetDB()
	var doctor model.Doctor
	doctor.Id = int32(id)
	if err := db.Model(&model.Doctor{}).Where("id = ?", id).Delete(&doctor).Error; err != nil {
		return err
	}
	return nil
}

func (c *DoctorService) AddDoctor(doctor *model.Doctor) error {
	db := database.GetDB()
	if err := db.Model(&doctor).Create(doctor).Error; err != nil {
		return err
	}
	return nil
}

func (c *DoctorService) UpdateDoctor(doctor *model.Doctor) error {
	db := database.GetDB()
	if err := db.Model(&doctor).Where("id = ?", doctor.Id).Save(doctor).Error; err != nil {
		return err
	}
	return nil
}

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
