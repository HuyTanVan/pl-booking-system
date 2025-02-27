package initializer

import (
	"fmt"
	"plbooking_go_structure1/global"
	worker "plbooking_go_structure1/internal/redis_workers"

	"github.com/hibiken/asynq"
)

func InitRedisTaskDistributer() *worker.TaskDistributor {
	redisOpt := asynq.RedisClientOpt{
		Addr: fmt.Sprintf("%s:%d", global.Config.Redis.Host, global.Config.Redis.Port),
	}
	d := worker.NewRedisTaskDistributor(redisOpt)
	fmt.Println("initialized redis background task distributer successfully")
	return &d
}
