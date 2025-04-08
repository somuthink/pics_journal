package queue

import (
	"context"

	"github.com/somuthink/pics_journal/core/internal/models"
)

func QueueLlmJob(ctx context.Context, job models.LlmJob) error {
	return RDB.LPush(ctx, "llm_queue", job).Err()
}
