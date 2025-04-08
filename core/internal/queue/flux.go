package queue

import (
	"context"
	"encoding/json"

	"github.com/somuthink/pics_journal/core/internal/models"
)

func QueueFluxJob(ctx context.Context, job models.FluxJob) error {
	return RDB.LPush(ctx, "flux_queue", job).Err()
}

func DequeueFluxJob(ctx context.Context) (models.FluxJob, error) {
	var job models.FluxJob
	data, err := RDB.BRPop(ctx, 0, "flux_queue").Result()

	err = json.Unmarshal([]byte(data[1]), &job)

	return job, err
}
