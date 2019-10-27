package service

import (
	"giligili/model"
	"giligili/serializer"
)

// ListCommentService 评论列表服务
type ListCommentService struct {
	VideoID string `form:"video_id"`
}

// List 评论列表
func (service *ListCommentService) List() serializer.Response {
	var comments []model.Comment
	total := 0
	// list only top level comments
	if err := model.DB.Where("video_id = ? and parent_id is null", service.VideoID).Find(&comments).Count(&total).Error; err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}

	return serializer.BuildListResponse(serializer.BuildComments(comments), uint(total))
}
