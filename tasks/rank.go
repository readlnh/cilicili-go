package tasks

import (
	"cilicili-go/cache"
)

// RestartDailyRank 重置每日排行
func RestartDailyRank() error {
	return cache.RedisClient.Del(cache.RedisClient.Context(), "rank:daily").Err()
}
