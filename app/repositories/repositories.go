package repositories

import (
	"strings"
	"github.com/mapsf/api/app/db"
	"github.com/mapsf/api/app/models"
	"github.com/jinzhu/gorm"
)

func GetUserByLoginAndPassword(login, password string) (*models.Player, error) {

	var user = &models.Player{}

	err := db.Conn.First(user, "LOWER(login) = ? AND password = ?", strings.ToLower(login), password).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func FindOnlinePlayers() ([]models.Player, error) {
	var players = make([]models.Player, 0)
	err := db.Conn.Find(&players, "online = ?", true).Error
	if err != nil {
		return nil, err
	}

	return players, nil
}

func FindNearPlayers(player *models.Player, distanceMeters float64) ([]models.Player, error) {
	players := make([]models.Player, 0)
	err := db.Conn.Model(&models.Player{}).
		Where("ST_DWithin(location, ?, ?) AND id <> ?", player.Location, distanceMeters, player.ID).
		Find(&players).
		Error
	if err != nil {
		return players, err
	}
	return players, nil
}

func GetPlayerByID(id uint) (*models.Player, error) {
	player := new(models.Player)
	err := db.Conn.Model(&models.Player{}).
		Where("id = ?", id).
		First(player).
		Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return player, err
}

// distance in meters
func GetDistanceBetweenTwoPlayers(from *models.Player, to *models.Player) (float64, error) {

	type result struct {
		Distance float64 `json:"distance"`
	}

	res := new(result)
	err := db.Conn.Raw("SELECT st_distance(?, ?) distance", from.Location, to.Location).Scan(res).Error
	if err != nil {
		return 0, err
	}
	return res.Distance, nil
}
