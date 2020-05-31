package service

import (
	"cilicili-go/serializer"
	"os"

	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
)

// QiniuTokenService 获取七牛云token
type QiniuTokenService struct {
}

// GetToken 获取七牛云token
func (service *QiniuTokenService) GetToken() serializer.Response {
	putPolicy := storage.PutPolicy{
		Scope: os.Getenv("QINIU_BUCKET"),
	}
	putPolicy.Expires = 7200
	mac := qbox.NewMac(os.Getenv("AK"), os.Getenv("SK"))
	upToken := putPolicy.UploadToken(mac)

	return serializer.Response{
		Data: map[string]string{
			"token": upToken,
		},
	}
}
