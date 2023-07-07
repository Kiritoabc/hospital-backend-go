package dto

import "hospital-backend-go/model"

type UserRegisterDto struct {
	Name      string  `json:"name"`
	Sex       int32   `json:"sex"`
	Birthday  string  `json:"birthday"`
	IdNum     string  `json:"idNum"`
	Phone     string  `json:"phone"`
	Fee       float64 `json:"fee"`
	Introduce string  `json:"introduce"`
	DepartId  int32   `json:"departId"`
	Password  string  `json:"password"`
}

func UserRegisterDtoToUser(dto UserRegisterDto) model.User {
	return model.User{
		Phone:    dto.Phone,
		Password: dto.Password,
	}
}

func UserRegisterDtoToDoctor(dto UserRegisterDto) model.Doctor {
	return model.Doctor{
		Name:      dto.Name,
		Sex:       dto.Sex,
		Birthday:  model.TimeStringToGoTime(dto.Birthday),
		IdNum:     dto.IdNum,
		Phone:     dto.Phone,
		Fee:       dto.Fee,
		Introduce: dto.Introduce,
		DepartId:  dto.DepartId,
		//Avatar:    model.Avatar, // 获取默认的avatar
		//Role:      0,            // 设置默认的Role
	}
}
