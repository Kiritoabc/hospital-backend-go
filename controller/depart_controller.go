package controller

import (
	"github.com/gin-gonic/gin"
	"hospital-backend-go/model"
	"hospital-backend-go/response"
	"hospital-backend-go/service"
	"strconv"
)

type DepartController struct {
}

func (c *DepartController) GetAllIds(ctx *gin.Context) {
	departService := service.DepartService{}
	var ids []string
	err := departService.GetIds(&ids)
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}
	response.Success(ctx, gin.H{
		"departIds": ids,
	}, "查询科室Ids成功")
	return
}

func (c *DepartController) DepartList(ctx *gin.Context) {
	departName, err := strconv.Atoi(ctx.DefaultQuery("searchName", "0"))
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}
	pageNum, err := strconv.Atoi(ctx.DefaultQuery("pageNum", "1"))
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}
	if pageNum <= 0 {
		pageNum = 1
	}
	pageSize, err := strconv.Atoi(ctx.DefaultQuery("pageSize", "5"))
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}
	if pageSize <= 0 {
		pageSize = 5
	}
	departId, err := strconv.Atoi(ctx.DefaultQuery("departId", "0"))
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}

	departService := service.DepartService{}
	total, err := departService.GetDepartListCount()
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}
	var departs []model.Depart
	err = departService.GetDepartList(strconv.Itoa(departName), pageNum, pageSize, departId, &departs)
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}
	response.Success(ctx, gin.H{
		"items": gin.H{
			"records": departs,
			"total":   total,
		},
	}, "查询成功")
	return

}
func (c *DepartController) DeleteById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.DefaultQuery("id", "0"))
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}
	departService := service.DepartService{}
	err = departService.DeleteById(id)
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}
	response.Success(ctx, nil, "删除成功")
	return
}
func (c *DepartController) AddDepart(ctx *gin.Context) {

}
func (c *DepartController) Update(ctx *gin.Context) {
	var depart model.Depart
	err := ctx.ShouldBindJSON(&depart)
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}
	departService := service.DepartService{}
	err = departService.UpdateDepart(&depart)
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}
	response.Success(ctx, nil, "更新成功")
	return
}
