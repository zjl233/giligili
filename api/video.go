package api

import (
	"giligili/service"

	"github.com/gin-gonic/gin"
)

// CreateVideo 视频投稿
func CreateVideo(c *gin.Context) {
	serv := service.CreateVideoService{}
	if err := c.ShouldBind(&serv); err == nil {
		res := serv.Create(CurrentUser(c))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// ShowVideo 视频详情接口
func ShowVideo(c *gin.Context) {
	serv := service.ShowVideoService{}
	res := serv.Show(c.Param("id"))
	c.JSON(200, res)
}

// ListVideo 视频列表接口
func ListVideo(c *gin.Context) {
	serv := service.ListVideoService{}
	if err := c.ShouldBind(&serv); err == nil {
		res := serv.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UpdateVideo 更新视频的接口
func UpdateVideo(c *gin.Context) {
	serv := service.UpdateVideoService{}
	if err := c.ShouldBind(&serv); err == nil {
		res := serv.Update(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// DeleteVideo 删除视频的接口
func DeleteVideo(c *gin.Context) {
	serv := service.DeleteVideoService{}
	res := serv.Delete(c.Param("id"))
	c.JSON(200, res)
}
