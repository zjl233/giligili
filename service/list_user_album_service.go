package service

import (
	"giligili/model"
	"giligili/serializer"
)

// ListAlbumService 视频列表服务
type ListUserAlbumService struct {
	Limit int `form:"limit"`
	Start int `form:"start"`
}

// List 视频列表
func (service *ListUserAlbumService) List(user *model.User) serializer.Response {
	albums := []model.Album{}
	total := 0

	if service.Limit == 0 {
		service.Limit = 6
	}

	if err := model.DB.Model(model.Album{}).Count(&total).Error; err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}

	if err := model.DB.Where("user_id = ?", user.ID).Limit(service.Limit).Offset(service.Start).Find(&albums).Error; err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}

	return serializer.BuildListResponse(serializer.BuildAlbums(albums), uint(total))
}
