package initializer

import (
	"context"
	"fmt"
	"plbooking_go_structure1/global"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func InitRedis() {
	r := global.Config.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%v", r.Host, r.Port),
		Password: r.Password,
		DB:       0,
		PoolSize: 10,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Printf("failed to initialize redis-> %v", err)
		return
	}

	fmt.Println("inititialized redis sucessfully")
	global.Rdb = rdb
}
