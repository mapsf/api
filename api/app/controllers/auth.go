package controllers

import (
	"github.com/mapsf/api/api/app/auth"
	"github.com/mapsf/api/api/app/repositories"
	"github.com/mapsf/api/api/app/common"
	"net/http"
	"encoding/json"
)

type authRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func parseRequest(r *http.Request, body interface{}) (error) {
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		return err
	}
	return nil
}

func AuthHandler(p common.Params) common.ResponseRenderer {

	body := &authRequest{}
	err := parseRequest(p.Request, body)
	if err != nil {
		return common.JsonErrorMessageResponse(err.Error(), 500)
	}

	if len(body.Login) == 0 || len(body.Password) == 0 {
		return common.JsonErrorMessageResponse("введите логин и пароль", 400)
	}

	user, err := repositories.GetUserByLoginAndPassword(body.Login, body.Password)
	if err != nil {
		return common.JsonErrorMessageResponse("неверный логин или пароль", 400)
	}

	token, err := auth.GenerateUserBasedJwtToken(user)
	if err != nil {
		return common.JsonErrorMessageResponse(err.Error(), 400)
	}

	return common.JsonMapResponse(map[string]interface{}{
		"token": token,
	})
}

func Me(p common.Params) common.ResponseRenderer {
	return common.JsonMapResponse(p.User)
}

func ValidateToken(p common.Params) common.ResponseRenderer {
	return common.JsonErrorMessageResponse("Token is live", 200)
}
