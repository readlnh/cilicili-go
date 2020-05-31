package api

import (
	"cilicili-go/service"

	"github.com/gin-gonic/gin"
)

// UploadToken 上传授权
func UploadToken(c *gin.Context) {
	service := service.QiniuTokenService{}
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, ErrorResponse(err))
	} else {
		res := service.GetToken()
		c.JSON(200, res)
	}
}
