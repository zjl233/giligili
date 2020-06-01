package service

import (
	"giligili/model"
	"giligili/serializer"
)

// ShowAlbumService 投稿详情的服务
type ShowAlbumService struct {
}

// Show 视频
func (service *ShowAlbumService) Show(id string) serializer.Response {
	var album model.Album
	err := model.DB.First(&album, id).Error
	if err != nil {
		return serializer.Response{
			Status: 404,
			Msg:    "视频不存在",
			Error:  err.Error(),
		}
	}

	//处理视频被观看的一系问题
	//album.AddView()

	return serializer.Response{
		Data: serializer.BuildAlbum(album),
	}
}
