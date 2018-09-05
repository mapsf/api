package repositories

import (
	"strings"
	"github.com/mapsf/api/api/app/db"
	"github.com/mapsf/api/api/app/models"
)

func GetUserByLoginAndPassword(login, password string) (*models.Character, error) {

	var user = &models.Character{}

	err := db.Conn.First(user, "LOWER(login) = ? AND password = ?", strings.ToLower(login), password).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func FindOnlinePlayers() ([]models.Character, error) {
	var players = make([]models.Character, 0)
	err := db.Conn.Find(&players, "online = ?", true).Error
	if err != nil {
		return nil, err
	}

	return players, nil
}
