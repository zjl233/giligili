package model

import (
	"github.com/jinzhu/gorm"
)

const (
	// 公开的相册
	Public string = "public"
	// 私有的相册
	Private string = "private"
	// 保护的相册
	Protect string = "protect"
)

// Album 视频模型
type Album struct {
	gorm.Model
	Title        string
	Category     string
	Visibility   string
	PasswordHash string
	UserID       uint
}

// Photos 封面地址
//func (album *Album) Photos() []string {
//
//}

// CheckPassword 校验密码
func (album *Album) CheckPassword(password string) bool {
	//err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return password == album.PasswordHash
}

//// AlbumURL 视频地址
//func (album *Album) AlbumURL() []string {
//	return SignedGetURL(album.URL)
//}
//
//// View 点击数
//func (album *Album) View() uint64 {
//	count, _ := cache.RedisClient.Get(cache.AlbumViewKey(album.ID)).Uint64()
//	return count
//}
//
//// AddView 视频游览
//func (album *Album) AddView() {
//	// 增加视频点击数
//	cache.RedisClient.Incr(cache.AlbumViewKey(album.ID))
//	// 增加排行点击数
//	cache.RedisClient.ZIncrBy(cache.DailyRankKey, 1, strconv.Itoa(int(album.ID)))
//}
