package llm

import (
	"context"
	"time"

	"github.com/charmbracelet/log"
	"github.com/somuthink/pics_journal/core/internal/db/events"
	"github.com/somuthink/pics_journal/core/internal/handlers/socket"
	"github.com/somuthink/pics_journal/core/internal/models"
	"github.com/somuthink/pics_journal/core/internal/queue"
)

func StartWorker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			log.Info("Worker shutting down...")
			return
		default:

			job, err := queue.DequeueLlmJob(ctx)
			log.Info("worker started new job", "job", job)

			if err != nil {
				if ctx.Err() != nil {
					log.Info("Worker exiting due to context cancel")
					return
				}
				log.Warn("Error in job dequeue", "err", err)
				time.Sleep(100 * time.Millisecond)
				continue
			}

			// cancelStatus, err := queue.GetCancel(ctx, job.JobID)
			// if err != nil {
			// 	log.Warn("Error checking cancel status", "jobID", job.JobID, "err", err)
			// }
			// if cancelStatus {
			// 	log.Info("Skipping job", "jobID", job.JobID)
			// 	continue
			// }

			time.Sleep(2 * time.Second)

			var evs []models.Event

			evs = append(evs, models.Event{Content: job.Prompt, UserID: job.UserID})

			events.InsertEventsBatch(evs)

			select {
			case socket.HUB.Output <- socket.LlmOutput{UserID: job.UserID, Error: err, Result: "подумал", LlmJob: job}:
			case <-ctx.Done():
				log.Info("Worker shutting down while sending output...")
				return
			}
		}
	}
}
