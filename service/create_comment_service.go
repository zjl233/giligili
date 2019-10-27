package service

import (
	"giligili/model"
	"giligili/serializer"
)

// CreateCommentService 评论提交的服务
type CreateCommentService struct {
	VideoID  string `form:"video_id" json:"video_id"`
	Info     string `form:"info" json:"info" binding:"max=3000"`
	ParentID string `form:"parent_id" json:"parent_id"`
}

// Create 创建评论
func (service *CreateCommentService) Create() serializer.Response {
	comment := model.Comment{
		VideoID:  service.VideoID,
		Info:     service.Info,
		ParentID: service.ParentID,
	}

	err := model.DB.Create(&comment).Error
	if err != nil {
		return serializer.Response{
			Status: 50001,
			Msg:    "评论保存失败",
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildComment(comment),
	}
}
