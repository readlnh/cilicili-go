package service

import (
	"cilicili-go/model"
	"cilicili-go/serializer"
)

// UpdateVideoService 更新视频服务
type UpdateVideoService struct {
	Title string `form:"title" json:"title" binding:"required,min=2,max=30"`
	Info  string `form:"info" json:"info" binding:"max=300"`
}

// Update 视频详情
func (service UpdateVideoService) Update(id string) serializer.Response {
	var video model.Video
	if err := model.DB.First(&video, id).Error; err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "视频不存在",
			Error: err.Error(),
		}
	}

	video.Title = service.Title
	video.Info = service.Info
	if err := model.DB.Save(&video).Error; err != nil {
		return serializer.Response{
			Code:  50001,
			Msg:   "视频更新失败",
			Error: err.Error(),
		}
	}

	return serializer.BuildVideoResponse(video)
}
