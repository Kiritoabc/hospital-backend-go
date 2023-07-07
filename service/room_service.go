package service

import (
	"hospital-backend-go/database"
	"hospital-backend-go/model"
)

type RoomService struct {
}

func (s *RoomService) GetRoomListCount() (count int, err error) {
	db := database.GetDB()
	var t int64
	if err := db.Model(&model.Room{}).Count(&t).Error; err != nil {
		return 0, err
	}
	count = int(t)
	return count, nil
}

func (s *RoomService) GetRoomList(roomName string, pageNum int, pageSize int, roomId int, departId int, rooms *[]model.Room) error {
	db := database.GetDB()

	tx := db.Model(&model.Room{})

	if roomName != "0" {
		tx = tx.Where("room_name = ?", roomName)
	}

	if roomId != 0 {
		tx = tx.Where("room_id = ?", roomId)
	}

	if departId != 0 {
		tx = tx.Where("depart_id = ?", departId)
	}

	if err := tx.Offset((pageNum - 1) * pageSize).
		Limit(pageSize).Find(rooms).Error; err != nil {
		return err
	}
	return nil
}

func (s *RoomService) DeleteById(id int) error {
	db := database.GetDB()
	var room model.Room
	room.RoomId = int32(id)
	if err := db.Model(&model.Room{}).Where("room_id = ?", id).Delete(&room).Error; err != nil {
		return err
	}
	return nil
}

func (c *RoomService) AddRoom(room *model.Room) error {
	db := database.GetDB()
	if err := db.Model(&room).Create(room).Error; err != nil {
		return err
	}
	return nil
}

func (c *RoomService) UpdateRoom(room *model.Room) error {
	db := database.GetDB()
	if err := db.Model(&room).Where("room_id = ?", room.RoomId).Save(room).Error; err != nil {
		return err
	}
	return nil
}
