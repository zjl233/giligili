package api

import (
	"giligili/service"

	"github.com/gin-gonic/gin"
)

// CreateComment 评论提交
func CreateComment(c *gin.Context) {
	serv := service.CreateCommentService{}
	if err := c.ShouldBind(&serv); err == nil {
		res := serv.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// ListComment 评论列表接口
func ListComment(c *gin.Context) {
	serv := service.ListCommentService{}
	if err := c.ShouldBind(&serv); err == nil {
		res := serv.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
