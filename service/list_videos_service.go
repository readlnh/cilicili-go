package service

import (
	"cilicili-go/model"
	"cilicili-go/serializer"
)

// ListVideoService 视频列表服务
type ListVideoService struct {
}

// List 视频列表
func (service ListVideoService) List() serializer.Response {
	var videos []model.Video
	if err := model.DB.Find(&videos).Error; err != nil {
		return serializer.Response{
			Code:  50000,
			Msg:   "数据库查询错误",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildVideos(videos),
	}
}
