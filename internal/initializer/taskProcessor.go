package initializer

import (
	"fmt"
	"plbooking_go_structure1/global"
	worker "plbooking_go_structure1/internal/redis_workers"

	"github.com/hibiken/asynq"
)

func InitRedisTaskProccessor() worker.TaskProcessor {
	redisOpt := asynq.RedisClientOpt{
		Addr: fmt.Sprintf("%s:%d", global.Config.Redis.Host, global.Config.Redis.Port),
	}
	w := worker.NewRedisTaskProcessor(redisOpt, global.Pgdbc)
	fmt.Println("initialized redis background task processor successfully")
	return w
}
