package api

import (
	"giligili/service"
	"github.com/gin-gonic/gin"
)

// ListPhoto 视频列表接口
func ListPhoto(c *gin.Context) {
	serv := service.ListPhotoService{}
	if err := c.ShouldBind(&serv); err == nil {
		res := serv.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// DeletePhoto 删除视频的接口
func DeletePhoto(c *gin.Context) {
	serv := service.DeletePhotoService{}
	res := serv.Delete(c.Param("id"))
	c.JSON(200, res)
}

// SearchPhoto 更新视频的接口
func SearchPhoto(c *gin.Context) {
	serv := service.SearchPhotoService{}
	if err := c.ShouldBind(&serv); err == nil {
		res := serv.Search()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
