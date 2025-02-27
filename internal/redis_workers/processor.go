package redis_worker

import (
	"context"
	"fmt"
	"log"

	db "plbooking_go_structure1/internal/db/sqlc"

	"github.com/hibiken/asynq"
)

type TaskProcessor interface {
	Start() error
	ProcessTaskSendVerifyEmail(ctx context.Context, task *asynq.Task) error
	ProcessTaskSendOrderSuccess(ctx context.Context, task *asynq.Task) error
}

type RedisTaskProcessor struct {
	server *asynq.Server
	store  db.Store
}

func NewRedisTaskProcessor(redisOpt asynq.RedisClientOpt, store db.Store) TaskProcessor {
	server := asynq.NewServer(redisOpt, asynq.Config{})
	return &RedisTaskProcessor{
		server: server,
		store:  store,
	}
}

func (processor *RedisTaskProcessor) Start() error {
	mux := asynq.NewServeMux()
	mux.HandleFunc(TaskSendVerifyEmail, processor.ProcessTaskSendVerifyEmail)
	mux.HandleFunc(TaskSendOrderSuccess, processor.ProcessTaskSendOrderSuccess)

	if err := processor.server.Start(mux); err != nil {
		log.Fatalf("failed to run redis background task processor %v", err)
	}
	fmt.Println("run redis background task processor successfully")
	return nil
}
