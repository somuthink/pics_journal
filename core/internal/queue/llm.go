package queue

import (
	"context"
	"encoding/json"

	"github.com/somuthink/pics_journal/core/internal/models"
)

func QueueLlmJob(ctx context.Context, job models.LlmJob) error {
	return RDB.LPush(ctx, "llm_queue", job).Err()
}

func DequeueLlmJob(ctx context.Context) (models.LlmJob, error) {
	var job models.LlmJob
	data, err := RDB.BRPop(ctx, 0, "llm_queue").Result()

	err = json.Unmarshal([]byte(data[1]), &job)

	return job, err
}
