package redis_worker

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"plbooking_go_structure1/global"
	"plbooking_go_structure1/internal/mail"

	// "github.com/HuyTanVan/soccer_booking_ticket/util"

	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
)

const (
	TaskSendVerifyEmail = "task:send_verify_email"
)

type PayloadSendVerifyEmail struct {
	Email string `json:"email"`
}

// dequeue
func (distributor *RedisTaskDistributor) DistributeTaskSendVerifyEmail(ctx context.Context,
	payload *PayloadSendVerifyEmail,
	opts ...asynq.Option) error {

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal task payload: %s", err)
	}
	task := asynq.NewTask(TaskSendVerifyEmail, jsonPayload, opts...)
	infor, err := distributor.client.EnqueueContext(ctx, task)
	// fmt.Println("ENQUEUE THE TASK", infor)
	if err != nil {
		return fmt.Errorf("failed to enqueue task : %s", err)
	}

	log.Info().Str("type", task.Type()).Bytes("payload", task.Payload()).
		Str("queue", infor.Queue).Int("max_entry", infor.MaxRetry).Msg("enqueued task")
	return nil
}

func (processor *RedisTaskProcessor) ProcessTaskSendVerifyEmail(ctx context.Context, task *asynq.Task) error {
	var payload PayloadSendVerifyEmail
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return fmt.Errorf("failed to unmarshal payload: %w", asynq.SkipRetry)
	}

	user, err := processor.store.GetUserByEmail(ctx, payload.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("user doesn't exist: %w", asynq.SkipRetry)
		}
		return fmt.Errorf("failed to get user: %w", err)
	}
	token, err := processor.store.GetEVTokenByUserID(ctx, user.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("verification token doesn't exist: %w", asynq.SkipRetry)
		}
		return fmt.Errorf("failed to get verification token: %w", err)
	}

	// config, err := util.LoadConfig("..")
	// if err != nil {
	// 	return fmt.Errorf("can not load config in task_send_verify_email: %w", err)
	// }

	subject := "Welcome to Premier League Booking"
	// // TODO: replace this URL with an environment variable that points to a front-end page
	verifyUrl := fmt.Sprintf("http://localhost:8990/api/v1/users/verify_email?user_id=%d&token=%s",
		user.ID, token.Token)
	content := fmt.Sprintf(`Hello %s,<br/>
	Thank you for registering with us!<br/>
	Please <a href="%s">click here</a> to verify your email address.<br/>
	`, user.Email, verifyUrl)
	to := []string{user.Email}
	mailSender := mail.NewGmailSender(
		global.Config.EmailSender.EmailSenderName,
		global.Config.EmailSender.EmailSenderAddress,
		global.Config.EmailSender.EmailSenderPassword)

	err = mailSender.SendEmail(subject, content, to, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to send verify email: %w", err)
	}

	log.Info().Str("type", task.Type()).Bytes("payload", task.Payload()).
		Str("emai", user.Email).Msg("processed task")
	return nil
}
