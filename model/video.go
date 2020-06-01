package model

import (
	"os"
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
