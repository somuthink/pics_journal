package users

import (
	"github.com/somuthink/pics_journal/core/internal/db"
	"github.com/somuthink/pics_journal/core/internal/models"
)

func SelectUser(userID uint) (models.User, error) {
	var user models.User

	err := db.DB.Where("id=?", userID).First(&user).Error

	return user, err
}
