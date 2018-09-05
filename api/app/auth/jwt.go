package auth

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"github.com/mapsf/api/api/app/models"
	"os"
	"fmt"
	"github.com/mapsf/api/api/app/db"
)

var (
	secretKey             = []byte(os.Getenv("JWT_SECRET_KEY"))
	tokenLifetimeDuration = time.Duration(60 * time.Minute)
)

func findCharacterByClaims(c jwt.MapClaims) (*models.Character, error) {
	id := c["sub"]
	char := &models.Character{}
	if err := db.Conn.First(char, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return char, nil
}

func GenerateUserBasedJwtToken(character *models.Character) (string, error) {

	token := jwt.New(jwt.SigningMethodHS512)

	token.Claims = jwt.MapClaims{
		"iat":   time.Now().Unix(),
		"exp":   time.Now().Add(tokenLifetimeDuration).Unix(),
		"sub":   character.ID,
		"login": character.Login,
	}

	str, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return str, nil
}

func ValidateJwt(val string) (*models.Character, error) {
	token, err := jwt.Parse(val, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, err
	}

	character, err := findCharacterByClaims(claims)
	if err != nil {
		return nil, err
	}

	return character, nil
}
