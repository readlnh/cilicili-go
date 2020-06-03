package cache

import (
	"cilicili-go/util"
	"context"
	"os"
	"strconv"

	"github.com/go-redis/redis/v8"
)

// RedisClient Redis缓存客户端单例
var RedisClient *redis.Client

var ctx = context.Background()

// Redis 在中间件中初始化redis链接
func Redis() {
	db, _ := strconv.ParseUint(os.Getenv("REDIS_DB"), 10, 64)
	client := redis.NewClient(&redis.Options{
		Addr:       os.Getenv("REDIS_ADDR"),
		Password:   os.Getenv("REDIS_PW"),
		DB:         int(db),
		MaxRetries: 1,
	})

	pong, err := client.Ping(ctx).Result()
	println(pong, err)

	if err != nil {
		util.Log().Panic("连接Redis不成功", err)
	}

	RedisClient = client
}
