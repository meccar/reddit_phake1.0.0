package worker

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"


// 	"github.com/hibiken/asynq"
// 	"github.com/rs/zerolog/log"
// )
// // mail "mail"

// // db "sqlc"
// // util "util"

// const TaskSendVerifyEmail = "task:send_verify_email"

// type PayloadSendVerifyEmail struct {
// 	Username string `json:"username"`
// }

// func (distributor *RedisTaskDistributor) DistributeTaskSendVerifyEmail(
// 	ctx context.Context,
// 	payload *PayloadSendVerifyEmail,
// 	opts ...asynq.Option,
// ) error {
// 	jsonPayload, err := json.Marshal(payload)
// 	if err != nil {
// 		return fmt.Errorf("failed to marshal task payload: %w", err)
// 	}

// 	task := asynq.NewTask(TaskSendVerifyEmail, jsonPayload, opts...)
// 	info, err := distributor.client.EnqueueContext(ctx, task)
// 	if err != nil {
// 		return fmt.Errorf("failed to enqueue task: %w", err)
// 	}

// 	log.Info().Str("type", task.Type()).Bytes("payload", task.Payload()).
// 		Str("queue", info.Queue).Int("max_retry", info.MaxRetry).Msg("enqueued task")
// 	return nil
// }

// func (processor *RedisTaskProcessor) ProcessTaskSendVerifyEmail(ctx context.Context, task *asynq.Task) error {
// 	var payload PayloadSendVerifyEmail
// 	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
// 		return fmt.Errorf("failed to unmarshal payload: %w", asynq.SkipRetry)
// 	}

// 	user := processor.handler.VerifyUsername(ctx, payload.Username)
// 	if !user {
//     	return fmt.Errorf("failed to verify username: %w", asynq.SkipRetry)
//   	}
	  
// 	fmt.Printf("ProcessTaskSendVerifyEmail Username: %s\n", payload.Username)
// 	fmt.Printf("ProcessTaskSendVerifyEmail user: %s\n", user)

// 	// verifyEmail, err := processor.handler.VerifyEmail(ctx, db.CreateVerifyEmailParams{
// 	// 	Username:   payload.Username,
// 	// 	SecretCode: util.RandomString(6),
// 	// })
// 	// if err != nil {
// 	// 	return fmt.Errorf("failed to create verify email: %w", err)
// 	// }

// 	// subject := "Welcome to Simple Bank"
// 	// // TODO: replace this URL with an environment variable that points to a front-end page
// 	// verifyUrl := fmt.Sprintf("http://localhost:8080/v1/verify_email?email_id=%d&secret_code=%s",
// 	// 	verifyEmail.ID, verifyEmail.SecretCode)
// 	// content := fmt.Sprintf(`Hello %s,<br/>
// 	// Thank you for registering with us!<br/>
// 	// Please <a href="%s">click here</a> to verify your email address.<br/>
// 	// `, payload.Username, verifyUrl)
// 	// to := []string{payload.Username}

// 	// err = processor.mailer.SendEmail(subject, content, to, nil, nil, nil)
// 	// if err != nil {
// 	// 	return fmt.Errorf("failed to send verify email: %w", err)
// 	// }

// 	// mail.Send()

// 	log.Info().Str("type", task.Type()).Bytes("payload", task.Payload()).
// 		Str("email", payload.Username).Msg("processed task")
// 	return nil
// }
