package service

import (
	"cilicili-go/model"
	"cilicili-go/serializer"
)

// DeleteVideoService 删除视频服务
type DeleteVideoService struct {
}

// Delete 删除视频
func (service DeleteVideoService) Delete(id string) serializer.Response {
	var video model.Video
	if err := model.DB.First(&video, id).Error; err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "视频不存在",
			Error: err.Error(),
		}
	}

	if err := model.DB.Delete(&video).Error; err != nil {
		return serializer.Response{
			Code:  50000,
			Msg:   "视频删除失败",
			Error: err.Error(),
		}
	}

	return serializer.Response{}
}
