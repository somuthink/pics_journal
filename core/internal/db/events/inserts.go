package events

import (
	"github.com/somuthink/pics_journal/core/internal/db"
	"github.com/somuthink/pics_journal/core/internal/models"
)

func InsertEventsBatch(events []models.Event) error {
	return db.DB.Create(events).Error
}
