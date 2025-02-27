package redis_worker

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"plbooking_go_structure1/global"
	"plbooking_go_structure1/internal/mail"

	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
)

const (
	TaskSendOrderSuccess = "task:send_order_success"
)

type PayloadSendOrderSuccess struct {
	OrderID int32  `json:"order_id"`
	Email   string `json:"email"`
}

func (distributor *RedisTaskDistributor) DistributeTaskSendOrderSuccess(ctx context.Context,
	payload *PayloadSendOrderSuccess,
	opts ...asynq.Option) error {

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal task payload: %s", err)
	}
	task := asynq.NewTask(TaskSendOrderSuccess, jsonPayload, opts...)
	infor, err := distributor.client.EnqueueContext(ctx, task)
	if err != nil {
		return fmt.Errorf("failed to enqueue task : %s", err)
	}

	log.Info().Str("type", task.Type()).Bytes("payload", task.Payload()).
		Str("queue", infor.Queue).Int("max_entry", infor.MaxRetry).Msg("enqueued task")
	return nil
}

func (processor *RedisTaskProcessor) ProcessTaskSendOrderSuccess(ctx context.Context, task *asynq.Task) error {
	var payload PayloadSendOrderSuccess
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
	// token, err := processor.store.GetEVTokenByUserID(ctx, user.ID)
	// if err != nil {
	// 	if err == sql.ErrNoRows {
	// 		return fmt.Errorf("verification token doesn't exist: %w", asynq.SkipRetry)
	// 	}
	// 	return fmt.Errorf("failed to get verification token: %w", err)
	// }

	subject := "Welcome to Premier League Booking"
	content := fmt.Sprintf(`Hello %s,<br/><br/>
		Thank you for booking with us today! We’re excited to confirm your ticket for the Premier League match.<br/><br/>
		**Booking Details:**<br/>
		- **Order ID:** [Order ID]<br/>
		- **Match:** [Team A] vs [Team B]<br/>
		- **Date:** [Match Date]<br/>
		- **Time:** [Match Time]<br/>
		- **Venue:** [Stadium Name, Location]<br/>
		- **Seat:** [Seat/Block Details]<br/><br/>
		Please keep your Order ID for future reference. Present this email at the venue for entry. If you have any questions or need assistance, feel free to contact us.<br/><br/>
		Enjoy the game!<br/><br/>
		`, user.FirstName.String)
	to := []string{user.Email}
	mailSender := mail.NewGmailSender(
		global.Config.EmailSender.EmailSenderName,
		global.Config.EmailSender.EmailSenderAddress,
		global.Config.EmailSender.EmailSenderPassword,
	)

	err = mailSender.SendEmail(subject, content, to, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to send order success: %w", err)
	}

	log.Info().Str("type", task.Type()).Bytes("payload", task.Payload()).
		Str("emai", user.Email).Msg("processed task")
	return nil
}
