package api

import (
	"giligili/service"
	"github.com/gin-gonic/gin"
)

func LikeVideo(c *gin.Context) {
	serv := service.LikeVideoService{}
	if err := c.ShouldBind(&serv); err == nil {
		res := serv.
			Like(CurrentUser(c))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
