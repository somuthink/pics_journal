package events

import (
	"fmt"
	"time"

	"github.com/somuthink/pics_journal/core/internal/db"
	"github.com/somuthink/pics_journal/core/internal/models"
)

type Stat struct {
	Emotions   map[string]int
	Categories map[string]int
}

func SelectEvents(userID uint, limit int, todayOnly bool) ([]models.Event, error) {
	var events []models.Event

	query := db.DB.Where("user_id = ?", userID).Order("created_at desc")

	if todayOnly {
		today := time.Now().Truncate(24 * time.Hour)
		query = query.Where("created_at >= ?", today)
	}

	res := query.Limit(limit).Find(&events)

	for i, j := 0, len(events)-1; i < j; i, j = i+1, j-1 {
		events[i], events[j] = events[j], events[i]
	}

	return events, res.Error
}

func SelectEventsAgenda(userID uint) (map[string]map[string][]models.Event, map[string]map[string]Stat, error) {
	var events []models.Event
	if err := db.DB.Where("user_id = ?", userID).Order("created_at asc").Find(&events).Error; err != nil {
		return nil, nil, err
	}

	groupedEvents := make(map[string]map[string][]models.Event)

	for _, event := range events {
		weekStart := event.CreatedAt
		for weekStart.Weekday() != time.Monday {
			weekStart = weekStart.AddDate(0, 0, -1)
		}
		year, week := weekStart.ISOWeek()
		weekKey := fmt.Sprintf("%d-W%02d", year, week)

		dayKey := event.CreatedAt.Weekday().String()

		if _, ok := groupedEvents[weekKey]; !ok {
			groupedEvents[weekKey] = make(map[string][]models.Event)
		}

		groupedEvents[weekKey][dayKey] = append(groupedEvents[weekKey][dayKey], event)
	}

	stats := make(map[string]map[string]Stat)

	for weekKey, weekData := range groupedEvents {
		if _, ok := stats[weekKey]; !ok {
			stats[weekKey] = make(map[string]Stat)
		}

		for dayKey, dayEvents := range weekData {
			dayStat := Stat{
				Emotions:   make(map[string]int),
				Categories: make(map[string]int),
			}

			for _, event := range dayEvents {
				dayStat.Emotions[event.Emotion]++
				dayStat.Categories[event.Category]++
			}

			stats[weekKey][dayKey] = dayStat
		}
	}

	return groupedEvents, stats, nil
}
