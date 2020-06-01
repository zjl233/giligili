package serializer

import (
	"giligili/model"
)

// Photo 照片序列化器
type Photo struct {
	ID        uint   `json:"id"`
	AlbumID   uint   `json:"album_id"`
	Name      string `json:"name"`
	URL       string `json:"url"`
	CreatedAt int64  `json:"created_at"`
}

// BuildPhoto 序列化评论
func BuildPhoto(item model.Photo) Photo {
	return Photo{
		ID:        item.ID,
		AlbumID:   item.AlbumID,
		Name:      item.Name,
		URL:       item.PhotoURL(),
		CreatedAt: item.CreatedAt.Unix(),
	}
}

// BuildPhotosFromAlbumID 序列化评论列表
func BuildPhotos(items []model.Photo) (photos []Photo) {
	for _, item := range items {
		photo := BuildPhoto(item)
		photos = append(photos, photo)
	}

	return photos
}
