package api

import (
	"cilicili-go/service"

	"github.com/gin-gonic/gin"
)

// CreateVideo 视频投稿
func CreateVideo(c *gin.Context) {
	service := service.CreateVideoService{}
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, ErrorResponse(err))
	} else {
		res := service.Create()
		c.JSON(200, res)
	}
}

// ShowVideo 视频详情接口
func ShowVideo(c *gin.Context) {
	service := service.ShowVideoService{}
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, ErrorResponse(err))
	} else {
		res := service.Show(c.Param("id"))
		c.JSON(200, res)
	}
}

// ListVideos 视频列表接口
func ListVideos(c *gin.Context) {
	service := service.ListVideoService{}
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, ErrorResponse(err))
	} else {
		res := service.List()
		c.JSON(200, res)
	}
}

// UpdateVideo 视频更新接口
func UpdateVideo(c *gin.Context) {
	service := service.UpdateVideoService{}
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, ErrorResponse(err))
	} else {
		res := service.Update(c.Param("id"))
		c.JSON(200, res)
	}
}

// DeleteVideo 删除视频接口
func DeleteVideo(c *gin.Context) {
	service := service.DeleteVideoService{}
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, ErrorResponse(err))
	} else {
		res := service.Delete(c.Param("id"))
		c.JSON(200, res)
	}
}
