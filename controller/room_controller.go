package controller

import (
	"github.com/gin-gonic/gin"
	"hospital-backend-go/model"
	"hospital-backend-go/response"
	"hospital-backend-go/service"
	"strconv"
)

type RoomController struct {
}

func (c *RoomController) RoomList(ctx *gin.Context) {
	// 获取参数
	roomName, err := strconv.Atoi(ctx.DefaultQuery("searchName", "0"))
	pageNum, err := strconv.Atoi(ctx.DefaultQuery("pageNum", "1"))
	if err != nil || pageNum < 1 {
		pageNum = 1
	}
	pageSize, err := strconv.Atoi(ctx.DefaultQuery("pageSize", "5"))
	if err != nil || pageSize < 1 {
		pageSize = 5
	}
	roomId, _ := strconv.Atoi(ctx.DefaultQuery("roomId", "0"))
	departId, _ := strconv.Atoi(ctx.DefaultQuery("departId", "0"))
	roomservice := service.RoomService{}
	total, err := roomservice.GetRoomListCount()
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}
	var rooms []model.Room
	err = roomservice.GetRoomList(strconv.Itoa(roomName), pageNum, pageSize, roomId, departId, &rooms)
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}
	response.Success(ctx, gin.H{
		"items": gin.H{
			"records": rooms,
			"total":   total,
		},
	}, "查询成功")
	return
}

func (c *RoomController) DeleteById(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.DefaultQuery("id", "0"))
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}
	roomservice := service.RoomService{}
	if err = roomservice.DeleteById(id); err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}

	response.Success(ctx, nil, "删除成功")
	return
}

func (c *RoomController) AddRoom(ctx *gin.Context) {
	var room = model.Room{}
	err := ctx.ShouldBindJSON(&room)
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}
	roomService := service.RoomService{}

	err = roomService.AddRoom(&room)
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}

	response.Success(ctx, nil, "添加成功")
	return
}

func (c *RoomController) Update(ctx *gin.Context) {
	var room = model.Room{}

	err := ctx.ShouldBindJSON(&room)
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}

	roomService := service.RoomService{}
	err = roomService.UpdateRoom(&room)
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}
	response.Success(ctx, nil, "更新")
	return
}
