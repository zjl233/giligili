package service

import (
	"giligili/model"
	"giligili/serializer"
)

// ListPhotoService 视频列表服务
type ListPhotoService struct {
	Limit int `form:"limit"`
	Start int `form:"start"`
}

// List 视频列表
func (service *ListPhotoService) List() serializer.Response {
	photos := []model.Photo{}
	total := 0

	if service.Limit == 0 {
		service.Limit = 6
	}

	if err := model.DB.Model(model.Photo{}).Count(&total).Error; err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}

	if err := model.DB.Limit(service.Limit).Offset(service.Start).Find(&photos).Error; err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}

	return serializer.BuildListResponse(serializer.BuildPhotos(photos), uint(total))
}
