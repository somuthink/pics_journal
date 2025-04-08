package users

import (
	"github.com/somuthink/pics_journal/core/internal/crypto"
	"github.com/somuthink/pics_journal/core/internal/db"
	"github.com/somuthink/pics_journal/core/internal/models"
)

func InsertOrSelect(name string, password string) (models.User, bool, error) {
	var user models.User
	result := db.DB.Where("name = ?", name).Attrs(models.User{Name: name, Password: crypto.GeneratePassword(password)}).FirstOrCreate(&user)

	if result.RowsAffected == 0 {
		return user, true, result.Error
	}

	return user, false, result.Error
}
