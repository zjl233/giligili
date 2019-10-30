package api

import (
	"giligili/service"

	"github.com/gin-gonic/gin"
)

// UploadAvatar 上传授权
func UploadAvatar(c *gin.Context) {
	serv := service.UploadAvatarService{}
	if err := c.ShouldBind(&serv); err == nil {
		res := serv.Post()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UploadVideo 上传授权
func UploadVideo(c *gin.Context) {
	serv := service.UploadVideoService{}
	if err := c.ShouldBind(&serv); err == nil {
		res := serv.Post()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
