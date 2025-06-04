package initialize

import (
	"github.com/go-redis/redis"
	"go-blog/global"
	"go.uber.org/zap"
	"os"
)

func ConnectRedis() redis.Client {
	redisConf := global.Config.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisConf.Address,
		Password: redisConf.Password,
		DB:       redisConf.DB,
	})

	_, err := client.Ping().Result()

	if err != nil {
		global.Log.Error("Failed to connect to Redis:", zap.Error(err))
		os.Exit(1)
	}

	return *client
}
