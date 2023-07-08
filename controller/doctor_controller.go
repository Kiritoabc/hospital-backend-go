package controller

import (
	"github.com/gin-gonic/gin"
	"hospital-backend-go/model"
	"hospital-backend-go/response"
	"hospital-backend-go/service"
	"strconv"
)

type DoctorController struct {
}

func (c *DoctorController) DoctorList(ctx *gin.Context) {
	// 获取参数
	name, err := strconv.Atoi(ctx.DefaultQuery("name", "0"))
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}
	pageNum, err := strconv.Atoi(ctx.DefaultQuery("pageNum", "0"))
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}
	if pageNum < 0 {
		pageNum = 1
	}
	pageSize, err := strconv.Atoi(ctx.DefaultQuery("pageSize", "0"))
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}
	if pageSize < 0 {
		pageSize = 5
	}
	id, err := strconv.Atoi(ctx.DefaultQuery("id", "0"))
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}
	idNum, err := strconv.Atoi(ctx.DefaultQuery("idNum", "0"))
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}
	phone, err := strconv.Atoi(ctx.DefaultQuery("phone", "0"))
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}
	fee, err := strconv.Atoi(ctx.DefaultQuery("fee", "0"))
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}
	departId, err := strconv.Atoi(ctx.DefaultQuery("departId", "0"))
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}

	role, err := strconv.Atoi(ctx.DefaultQuery("role", "-1"))
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}

	doctorService := service.DoctorService{}
	total, err := doctorService.GetDoctorListCount()
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}
	var doctors []model.Doctor
	err = doctorService.GetDoctorList(strconv.Itoa(name), pageNum, pageSize, id, idNum, strconv.Itoa(phone), strconv.Itoa(fee), departId, role, &doctors)
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
	}
	response.Success(ctx, gin.H{
		"items": gin.H{
			"records": doctors,
			"total":   total,
		},
	}, "查询成功")
	return
}

func (c *DoctorController) DeleteById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.DefaultQuery("id", "0"))
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}
	doctorService := service.RoomService{}
	err = doctorService.DeleteById(id)
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}
	response.Success(ctx, nil, "删除成功")
	return
}

func (c *DoctorController) AddDoctor(ctx *gin.Context) {
	var doctor model.Doctor
	err := ctx.ShouldBindJSON(doctor)
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}
	doctorService := service.DoctorService{}
	err = doctorService.AddDoctor(&doctor)
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}
	response.Success(ctx, nil, "删除成功")
	return
}

func (c *DoctorController) Update(ctx *gin.Context) {
	var doctor model.Doctor
	err := ctx.ShouldBindJSON(doctor)
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}
	doctorService := service.DoctorService{}
	err = doctorService.UpdateDoctor(&doctor)
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}
	response.Success(ctx, nil, "更新成功")
	return
}
