package tasks

import (
	"encoding/json"

	"github.com/hibiken/asynq"
)

const TypePing = "background:ping"

type PingPayload struct {
	Message string `json:"message"`
}

func NewPingTask(message string) *asynq.Task {
	payload, _ := json.Marshal(PingPayload{Message: message})
	return asynq.NewTask(TypePing, payload)
}
