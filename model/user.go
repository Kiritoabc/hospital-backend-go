package model

type User struct {
	Id       int32  `gorm:"column:id;type:INT;AUTO_INCREMENT;NOT NULL" json:"id"`
	Phone    string `gorm:"column:phone;type:VARCHAR(255);NOT NULL" json:"phone"`
	Password string `gorm:"column:password;type:VARCHAR(255);NOT NULL" json:"password"`
}
