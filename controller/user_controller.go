package controller

import (
	"github.com/gin-gonic/gin"
	"hospital-backend-go/database"
	"hospital-backend-go/dto"
	"hospital-backend-go/model"
	"hospital-backend-go/response"
	"hospital-backend-go/service"
	"hospital-backend-go/util"
	"net/http"
)

// Register 注册的接口
func Register(ctx *gin.Context) {
	//获取DB
	DB := database.GetDB()
	// 获取参数
	var registerUser = dto.UserRegisterDto{}
	err := ctx.ShouldBindJSON(&registerUser)
	if err != nil {
		return
	} // 解析 json -->UserRegisterDto
	phone := registerUser.Phone
	if service.IsPhoneExists(DB, phone) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, err.Error())
		return
	}

	var user = model.User{}
	user = dto.UserRegisterDtoToUser(registerUser)
	// 注册User
	err = service.AddUser(DB, user)
	if err != nil {
		response.Response(ctx, http.StatusBadRequest, 400, nil, err.Error())
		return
	}

	// 注册Doctor
	var doctor = model.Doctor{}
	doctor = dto.UserRegisterDtoToDoctor(registerUser)
	// 设置雪花算法的ID，默认头像和Role
	doctor.Uuid = util.GetID()
	doctor.Avatar = model.Avatar
	doctor.Role = 0

	err = service.AddDoctor(DB, doctor)
	if err != nil {
		response.Response(ctx, http.StatusBadRequest, 400, nil, err.Error())
		return
	}
	response.Success(ctx, nil, "注册成功")
	return
}
