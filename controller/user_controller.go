package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/spf13/viper"
	"hospital-backend-go/database"
	"hospital-backend-go/dto"
	"hospital-backend-go/model"
	"hospital-backend-go/response"
	"hospital-backend-go/service"
	"hospital-backend-go/util"
	"log"
	"net/http"
	"strconv"
	"time"
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

// Login 登录接口
func Login(ctx *gin.Context) {
	DB := database.GetDB()
	var loginUser = dto.UserLoginDto{}
	err := ctx.ShouldBindBodyWith(&loginUser, binding.JSON)
	// 绑定数据出问题
	if err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, err.Error())
		return
	}
	// 用户不存在
	if !service.IsPhoneExists(DB, loginUser.Phone) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户不存在")
		return
	}
	ok, err := service.AuthUser(DB, loginUser)
	if !ok || err != nil {
		response.Response(ctx, 400, 400, nil, "密码错误")
		return
	}
	var token string
	token, err = util.GenerateTokenUsingHs256(loginUser.Phone)
	if err != nil {
		response.Response(ctx, 400, 400, nil, err.Error())
		return
	}

	response.Success(ctx, gin.H{
		"token": token,
	}, "登录成功")
	return
}

// GetUserInfo 获取用户信息的接口
func GetUserInfo(ctx *gin.Context) {
	doctor, ok := ctx.Get("doctorInfo")
	if !ok {
		response.Fail(ctx, "请重新登录", gin.H{})
		return
	}
	response.Success(ctx, gin.H{
		"doctorVo": doctor,
	}, "获取信息成功过")
}

// UpdateAvatar 更新头像
func UpdateAvatar(ctx *gin.Context) {
	doctorInfo, ok := ctx.Get("doctorInfo")
	if !ok {
		response.Fail(ctx, "请重新登录", gin.H{})
		return
	}
	doctor, ok := doctorInfo.(model.Doctor)
	form, _ := ctx.MultipartForm()
	file := form.File["file"][0]
	//var filePathMinio string
	log.Printf("upload file path: %s\n", file.Filename)
	objectName := time.Now().Format("20060102") + "/" + strconv.Itoa(int(doctor.Id)) + "/" + file.Filename
	err := util.UploadToMinio(objectName, file, "application/octet-stream")
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}
	filePathMinio := "http://" + viper.GetString("minio.endpoint") + "/" + viper.GetString("minio.bucketName") + "/" + objectName
	doctor.Avatar = filePathMinio
	doctorService := service.DoctorService{}
	err = doctorService.UpdateDoctor(&doctor)
	if err != nil {
		response.Fail(ctx, err.Error(), gin.H{})
		return
	}
	response.Success(ctx, gin.H{}, "上传成功")
	return
}
