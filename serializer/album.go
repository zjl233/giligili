package serializer

import (
	"giligili/model"
	"log"
)

// Album 视频序列化器
type Album struct {
	ID         uint    `json:"id"`
	Title      string  `json:"title"`
	Category   string  `json:"category"`
	Visibility string  `json:"visibility"`
	Photos     []Photo `json:"photos"`
	CreatedAt  int64   `json:"created_at"`
	User       User    `json:"user"`
}

// BuildAlbum 序列化视频
func BuildAlbum(item model.Album) Album {
	user, _ := model.GetUser(item.UserID)
	return Album{
		ID:         item.ID,
		Title:      item.Title,
		Category:   item.Category,
		Visibility: item.Visibility,
		Photos:     BuildPhotosFromAlbumID(item.ID),
		CreatedAt:  item.CreatedAt.Unix(),
		User:       BuildUser(user),
	}
}

// BuildAlbums 序列化视频列表
func BuildAlbums(items []model.Album) (albums []Album) {
	for _, item := range items {
		album := BuildAlbum(item)
		albums = append(albums, album)
	}
	return albums
}

// BuildPhotosFromAlbumID 序列化评论列表，查表，获得与次 album 相关联的 photo
func BuildPhotosFromAlbumID(albumID uint) (photos []Photo) {

	items := []model.Photo{}

	if err := model.DB.Where("album_id = ?", albumID).Find(&items); err != nil {
		log.Println("数据库错误：", err) // 这里有点问题
	}

	for _, item := range items {
		photo := BuildPhoto(item)
		photos = append(photos, photo)
	}

	if photos == nil {
		photos = []Photo{}
	}

	return photos
}
