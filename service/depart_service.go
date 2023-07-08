package service

import (
	"hospital-backend-go/database"
	"hospital-backend-go/model"
)

type DepartService struct {
}

// GetIds Ids 获取depart的所有id
func (s *DepartService) GetIds(ids *[]string) error {
	db := database.GetDB()
	if err := db.Model(&model.Depart{}).
		Select("depart_id").
		First(ids).Error; err != nil {
		return err
	}
	return nil
}

func (s *DepartService) GetDepartListCount() (count int, err error) {
	db := database.GetDB()
	var t int64
	if err := db.Model(&model.Depart{}).Count(&t).Error; err != nil {
		return 0, err
	}
	count = int(t)
	return count, nil
}

func (s *DepartService) GetDepartList(departName string, pageNum int, pageSize int, departId int, departs *[]model.Depart) error {
	db := database.GetDB()

	tx := db.Model(&model.Depart{})

	if departName != "0" {
		tx = tx.Where("depart_name = ?", departName)
	}

	if departId != 0 {
		tx = tx.Where("depart_id = ?", departId)
	}

	if err := tx.Offset((pageNum - 1) * pageSize).
		Limit(pageSize).Find(departs).Error; err != nil {
		return err
	}
	return nil
}

func (s *DepartService) DeleteById(id int) error {
	db := database.GetDB()
	var depart model.Depart
	depart.DepartId = int32(id)
	if err := db.Model(&model.Depart{}).Where("depart_id = ?", id).Delete(&depart).Error; err != nil {
		return err
	}
	return nil
}

func (c *DepartService) AddDepart(depart *model.Depart) error {
	db := database.GetDB()
	if err := db.Model(&depart).Create(depart).Error; err != nil {
		return err
	}
	return nil
}

func (c *DepartService) UpdateDepart(depart *model.Depart) error {
	db := database.GetDB()
	if err := db.Model(&depart).Where("depart_id = ?", depart.DepartId).Save(depart).Error; err != nil {
		return err
	}
	return nil
}
