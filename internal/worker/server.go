package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/cakmakfatih/chattered-background/pkg/tasks"
	"github.com/hibiken/asynq"
)

func Run(redisAddr, redisPassword string) error {
	server := asynq.NewServer(
		asynq.RedisClientOpt{
			Addr:     redisAddr,
			Password: redisPassword,
		},
		asynq.Config{
			Concurrency: 10,
			Queues: map[string]int{
				"critical": 6,
				"default":  3,
				"low":      1,
			},
		},
	)

	mux := asynq.NewServeMux()
	mux.HandleFunc(tasks.TypePing, handlePing)

	return server.Run(mux)
}

func handlePing(ctx context.Context, task *asynq.Task) error {
	var payload tasks.PingPayload
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return fmt.Errorf("decode ping task payload: %w", err)
	}

	slog.InfoContext(ctx, "processed background task", "task_type", tasks.TypePing, "message", payload.Message)
	return nil
}
