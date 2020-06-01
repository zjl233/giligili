package service

import (
	"giligili/model"
	"giligili/serializer"
)

// CreateAlbumService 视频投稿的服务
type CreateAlbumService struct {
	Title      string   `form:"title" json:"title" binding:"required,min=2,max=100"`
	Category   string   `form:"category" json:"category"`
	Visibility string   `form:"visibility" json:"visibility"`
	Photos     []string `form:"photos" json:"photos"`
	Password   string   `form:"password" json:"password"`
}

// Create 创建视频
func (service *CreateAlbumService) Create(user *model.User) serializer.Response {
	// 创建 album
	album := model.Album{
		Title:        service.Title,
		Category:     service.Category,
		Visibility:   service.Visibility,
		PasswordHash: service.Password,
		UserID:       user.ID,
	}
	if err := model.DB.Create(&album).Error; err != nil {
		return serializer.Response{
			Status: 50001,
			Msg:    "相册创建失败",
			Error:  err.Error(),
		}
	}

	// album 在保存到数据库里以后，得到自增 id
	// 现在，创建与 album 相关联的 photos
	for _, p := range service.Photos {
		// p 是保存在 oss 里的，照片的 path + name
		photo := model.Photo{
			AlbumID: album.ID,
			Name:    p,
		}
		if err := model.DB.Create(&photo).Error; err != nil {
			return serializer.Response{
				Status: 50001,
				Msg:    "照片保存失败",
				Error:  err.Error(),
			}
		}
	}

	return serializer.Response{
		Data: serializer.BuildAlbum(album),
	}
}
