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

////ShowComment 评论详情接口
//func ShowComment(c *gin.Context) {
//	service := service.ShowCommentService{}
//	res := service.Show(c.Param("id"))
//	c.JSON(200, res)
//}
//
//// ListComment 评论列表接口
//func ListComment(c *gin.Context) {
//	service := service.ListCommentService{}
//	if err := c.ShouldBind(&service); err == nil {
//		res := service.List()
//		c.JSON(200, res)
//	} else {
//		c.JSON(200, ErrorResponse(err))
//	}
//}
//
//// UpdateComment 更新评论的接口
//func UpdateComment(c *gin.Context) {
//	service := service.UpdateCommentService{}
//	if err := c.ShouldBind(&service); err == nil {
//		res := service.Update(c.Param("id"))
//		c.JSON(200, res)
//	} else {
//		c.JSON(200, ErrorResponse(err))
//	}
//}
//
//// DeleteComment 删除评论的接口
//func DeleteComment(c *gin.Context) {
//	service := service.DeleteCommentService{}
//	res := service.Delete(c.Param("id"))
//	c.JSON(200, res)
//}
