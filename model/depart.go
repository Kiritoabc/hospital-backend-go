package model

type Depart struct {
	DepartId        int32  `gorm:"column:depart_id;type:INT;AUTO_INCREMENT;NOT NULL" json:"departId"`
	DepartName      string `gorm:"column:depart_name;type:VARCHAR(255);" json:"departName"`
	DepartIntroduce string `gorm:"column:depart_introduce;type:VARCHAR(255);" json:"departIntroduce"`
}
