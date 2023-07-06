package model

import "time"

type Call struct {
	CallId     int32     `gorm:"column:call_id;type:INT;NOT NULL" json:"callId"`
	DoctorId   int32     `gorm:"column:doctor_id;type:INT;" json:"doctorId"`
	CallDate   time.Time `gorm:"column:call_date;type:DATE;" json:"callDate"`
	FromTime   time.Time `gorm:"column:from_time;type:DATETIME;" json:"fromTime"`
	ToTime     time.Time `gorm:"column:to_time;type:DATETIME;" json:"toTime"`
	DoctorName string    `gorm:"column:doctor_name;type:VARCHAR(255);" json:"doctorName"`
}
