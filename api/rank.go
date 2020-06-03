package api

import (
	"cilicili-go/service"

	"github.com/gin-gonic/gin"
)

// DailyRank 每日排行榜
func DailyRank(c *gin.Context) {
	service := service.DailyRankService{}
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, ErrorResponse(err))
	} else {
		res := service.Get()
		c.JSON(200, res)
	}
}
