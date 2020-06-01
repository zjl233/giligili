package service

import (
	"giligili/model"
	"giligili/serializer"
)

// SearchPhotoService 视频列表服务
type SearchPhotoService struct {
	Name string `form:"name"`
}

// Search 视频列表
func (service *SearchPhotoService) Search() serializer.Response {
	photos := []model.Photo{}
	total := 0

	if err := model.DB.Model(model.Photo{}).Count(&total).Error; err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}

	// LIKE
	//db.Where("name LIKE ?", "%jin%").Find(&users)
	//// SELECT * FROM users WHERE name LIKE '%jin%';

	if err := model.DB.Where("name LIKE ?", "%"+service.Name+"%").Find(&photos).Error; err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}

	return serializer.BuildListResponse(serializer.BuildPhotos(photos), uint(total))
}
