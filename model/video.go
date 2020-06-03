package model

import (
	"cilicili-go/cache"
	"os"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/qiniu/api.v7/v7/auth"
	"github.com/qiniu/api.v7/v7/storage"
)

// Video 视频模型
type Video struct {
	gorm.Model
	Title  string
	Info   string
	Avatar string
	Video  string
}

// AvatarURL 封面地址
func (video *Video) AvatarURL() string {
	mac := auth.New(os.Getenv("AK"), os.Getenv("SK"))
	domain := os.Getenv("DOMAIN")
	key := video.Avatar
	deadline := time.Now().Add(time.Second * 3600).Unix() //1小时有效期
	privateAccessURL := storage.MakePrivateURL(mac, domain, key, deadline)

	return privateAccessURL
}

// VideoURL 封面地址
func (video *Video) VideoURL() string {
	mac := auth.New(os.Getenv("AK"), os.Getenv("SK"))
	domain := os.Getenv("DOMAIN")
	key := video.Video
	deadline := time.Now().Add(time.Second * 3600).Unix() //1小时有效期
	privateAccessURL := storage.MakePrivateURL(mac, domain, key, deadline)

	return privateAccessURL
}

// View 访问量
func (video *Video) View() uint64 {
	count, _ := cache.RedisClient.Get(cache.RedisClient.Context(), cache.VideoViewKey(video.ID)).Uint64()
	return count
}

// AddView 每次点击增加访问量
func (video *Video) AddView() {
	// 增加点击数
	cache.RedisClient.Incr(cache.RedisClient.Context(), cache.VideoViewKey(video.ID))
	// 增加排行点击数
	cache.RedisClient.ZIncrBy(cache.RedisClient.Context(), cache.DailyRankKey, 1, strconv.Itoa(int(video.ID)))
}
