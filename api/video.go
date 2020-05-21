package api

import (
	"cilicili-go/service"

	"github.com/gin-gonic/gin"
)

// CreateVideo 视频投稿
func CreateVideo(c *gin.Context) {
	var service service.CreateVideoService
	if err := c.ShouldBind(&service); err != nil {

	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
