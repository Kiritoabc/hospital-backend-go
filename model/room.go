package model

type Room struct {
	RoomId        int32  `gorm:"column:room_id;type:INT;AUTO_INCREMENT;NOT NULL" json:"roomId"`
	RoomName      string `gorm:"column:room_name;type:VARCHAR(255);NOT NULL" json:"roomName"`
	RoomIntroduce string `gorm:"column:room_introduce;type:VARCHAR(255);NOT NULL" json:"roomIntroduce"`
	DepartId      int32  `gorm:"column:depart_id;type:INT;NOT NULL" json:"departId"`
}
