package model

import (
	"giligili/cache"
	"strconv"

	"github.com/jinzhu/gorm"
)

// Video 视频模型
type Video struct {
	gorm.Model
	Title  string
	Info   string
	URL    string
	Avatar string
	UserID uint
}

// AvatarURL 封面地址
func (video *Video) AvatarURL() string {
	return SignedGetURL(video.Avatar)
}

// VideoURL 视频地址
func (video *Video) VideoURL() string {
	return SignedGetURL(video.URL)
}

// View 点击数
func (video *Video) View() uint64 {
	countStr, _ := cache.RedisClient.Get(cache.VideoViewKey(video.ID)).Result()
	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}

// AddView 视频游览
func (video *Video) AddView() {
	// 增加视频点击数
	cache.RedisClient.Incr(cache.VideoViewKey(video.ID))
	// 增加排行点击数
	cache.RedisClient.ZIncrBy(cache.DailyRankKey, 1, strconv.Itoa(int(video.ID)))
}
