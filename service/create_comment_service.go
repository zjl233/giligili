package service

import (
	"giligili/model"
	"giligili/serializer"
)

// CreateCommentService 评论提交的服务
type CreateCommentService struct {
	VideoId uint   `form:"video_id" json:"video_id"`
	Info    string `form:"info" json:"info" binding:"max=3000"`
}

// Create 创建评论
func (service *CreateCommentService) Create() serializer.Response {
	comment := model.Comment{
		VideoId: service.VideoId,
		Info:    service.Info,
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
