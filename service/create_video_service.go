package service

import (
	"cilicili-go/model"
	"cilicili-go/serializer"
)

// CreateVideoService 视频投稿服务
type CreateVideoService struct {
	Title string `form:"title" json:"title" binding:"required,min=2,max=100"`
	Info  string `form:"info" json:"info" binding:"max=500"`
}

// Create 创建视频
func (service *CreateVideoService) Create() serializer.Response {
	video := model.Video{
		Title: service.Title,
		Info:  service.Info,
	}
	if err := model.DB.Create(&video).Error; err != nil {
		return serializer.Response{
			Code:  50001,
			Msg:   "视频保存失败",
			Error: err.Error(),
		}
	}

	return serializer.BuildVideoResponse(video)
	// return serializer.Response{
	// 	Data: serializer.BuildVideo(video),
	// }
}
