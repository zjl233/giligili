package api

import (
	"giligili/service"

	"github.com/gin-gonic/gin"
)

// CreateAlbum 视频投稿
func CreateAlbum(c *gin.Context) {
	serv := service.CreateAlbumService{}
	if err := c.ShouldBind(&serv); err == nil {
		res := serv.Create(CurrentUser(c))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// ShowAlbum 视频详情接口
func ShowAlbum(c *gin.Context) {
	serv := service.ShowAlbumService{}
	res := serv.Show(c.Param("id"))
	c.JSON(200, res)
}

// ListAlbum 视频列表接口
func ListAlbum(c *gin.Context) {
	serv := service.ListAlbumService{}
	if err := c.ShouldBind(&serv); err == nil {
		res := serv.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// ListUserAlbum 视频列表接口
func ListUserAlbum(c *gin.Context) {
	serv := service.ListUserAlbumService{}
	if err := c.ShouldBind(&serv); err == nil {
		res := serv.List(CurrentUser(c))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}


// SearchAlbum 更新视频的接口
func SearchAlbum(c *gin.Context) {
	serv := service.SearchAlbumService{}
	if err := c.ShouldBind(&serv); err == nil {
		res := serv.Search()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

//// DeleteAlbum 删除视频的接口
//func DeleteAlbum(c *gin.Context) {
//	serv := service.DeleteAlbumService{}
//	res := serv.Delete(c.Param("id"))
//	c.JSON(200, res)
//}
