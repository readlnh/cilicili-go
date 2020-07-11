package service

import (
	"cilicili-go/cache"
	"cilicili-go/model"
	"cilicili-go/serializer"
	"fmt"
	"strings"
)

// DailyRankService 获取每日排名服务
type DailyRankService struct {
}

// Get 获取每日排名
func (service *DailyRankService) Get() serializer.Response {
	var videos []model.Video

	res := cache.RedisClient.ZRevRange(cache.RedisClient.Context(), cache.DailyRankKey, 0, 10)
	vids := res.Val()

	// 如果所有的浏览量都是0，那么不用查询了
	if len(vids) >= 1 {
		order := fmt.Sprintf("FIELD(id, %s)", strings.Join(vids, ","))
		println(order)
		err := model.DB.Where("id in (?)", vids).Order(order).Find(&videos).Error
		if err != nil {
			return serializer.Response{
				Code:  500000,
				Msg:   "数据库连接错误",
				Error: err.Error(),
			}
		}
	}

	return serializer.Response{
		Data: serializer.BuildVideos(videos),
	}
}
