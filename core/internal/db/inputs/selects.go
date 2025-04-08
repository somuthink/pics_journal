package inputs

import (
	"fmt"
	"time"

	"github.com/somuthink/pics_journal/core/internal/db"
	"github.com/somuthink/pics_journal/core/internal/models"
)

func SelectInputs(userID uint, limit int, todayOnly bool) ([]models.Input, error) {
	user := &models.User{}

	query := db.DB.Where("id = ?", userID)

	if todayOnly {
		today := time.Now().Truncate(24 * time.Hour)
		query = query.Where("created_at >= ?", today)
	}

	res := query.Preload("Inputs").Limit(limit).Find(user)

	return user.Inputs, res.Error
}

func SelectInputsAgenda() (map[string]map[string][]models.Input, error) {
	var inputs []models.Input
	if err := db.DB.Order("created_at asc").Find(&inputs).Error; err != nil {
		return nil, err
	}

	result := make(map[string]map[string][]models.Input)

	for _, input := range inputs {
		weekStart := input.CreatedAt
		for weekStart.Weekday() != time.Monday {
			weekStart = weekStart.AddDate(0, 0, -1)
		}
		year, week := weekStart.ISOWeek()
		weekKey := fmt.Sprintf("%d-W%02d", year, week)

		dayKey := input.CreatedAt.Weekday().String()

		if _, ok := result[weekKey]; !ok {
			result[weekKey] = make(map[string][]models.Input)
		}

		result[weekKey][dayKey] = append(result[weekKey][dayKey], input)
	}

	return result, nil
}
