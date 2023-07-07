package model

var Avatar = "avatar: \"https://n.sinaimg.cn/sinakd10109/410/w868h1142/20200824/3045-iyaiihm5428976.jpg\""

type Doctor struct {
	Id        int32   `gorm:"column:id;type:INT;AUTO_INCREMENT;NOT NULL" json:"id"`
	Name      string  `gorm:"column:name;type:VARCHAR(255);NOT NULL" json:"name"`
	Sex       int32   `gorm:"column:sex;type:INT;NOT NULL" json:"sex"`
	Birthday  Time    `gorm:"column:birthday;type:DATE;" json:"birthday"`
	IdNum     string  `gorm:"column:id_num;type:VARCHAR(255);" json:"idNum"`
	Phone     string  `gorm:"column:phone;type:VARCHAR(255);NOT NULL" json:"phone"`
	Fee       float64 `gorm:"column:fee;type:DOUBLE;" json:"fee"`
	Introduce string  `gorm:"column:introduce;type:VARCHAR(255);" json:"introduce"`
	DepartId  int32   `gorm:"column:depart_id;type:INT;" json:"departId"`
	Avatar    string  `gorm:"column:avatar;type:VARCHAR(255);NOT NULL" json:"avatar"`
	Uuid      string  `gorm:"column:uuid;type:VARCHAR(255);NOT NULL" json:"uuid"`
	Role      int32   `gorm:"column:role;type:INT;NOT NULL" json:"role"`
}
