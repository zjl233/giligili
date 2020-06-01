package service

import (
	"giligili/model"
	"giligili/serializer"
)

// SearchAlbumService 视频列表服务
type SearchAlbumService struct {
	Title    string `form:"title"`
	Category string `form:"category"`
}

// Search 视频列表
func (service *SearchAlbumService) Search() serializer.Response {
	albums := []model.Album{}
	total := 0

	if err := model.DB.Model(model.Album{}).Count(&total).Error; err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}

	// LIKE
	//db.Where("name LIKE ?", "%jin%").Find(&users)
	//// SELECT * FROM users WHERE name LIKE '%jin%';

	if err := model.DB.Where("title LIKE ? AND category LIKE ?", "%"+service.Title+"%", "%"+service.Category+"%").Find(&albums).Error; err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}

	return serializer.BuildListResponse(serializer.BuildAlbums(albums), uint(total))
}
