package service

import (
	"cilicili-go/model"
	"cilicili-go/serializer"
)

// ShowVideoService 视频详情服务
type ShowVideoService struct {
}

// Show 视频详情
func (service ShowVideoService) Show(id string) serializer.Response {
	var video model.Video
	if err := model.DB.First(&video, id).Error; err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "视频不存在",
			Error: err.Error(),
		}
	}

	video.AddView()

	return serializer.BuildVideoResponse(video)
}
