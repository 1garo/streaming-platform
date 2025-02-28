package queue

import (
	"github.com/hibiken/asynq"
	"github.com/learn-video/streaming-platform/internal/config"
)

func NewClient(cfg *config.Config) *asynq.Client {
	return asynq.NewClient(asynq.RedisClientOpt{Addr: cfg.RedisAddr})
}
