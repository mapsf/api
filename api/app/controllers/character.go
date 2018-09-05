package controllers

import (
	"github.com/mapsf/api/api/app/common"
	"github.com/mapsf/api/api/app/repositories"
)

func GetCharacter(p common.Params) common.ResponseRenderer {
	return common.JsonResponse{
		Data:   p.User,
		Status: 200,
	}
}

func GetOnlinePlayers(p common.Params) common.ResponseRenderer {

	players, err := repositories.FindOnlinePlayers()
	if err != nil {
		return common.ServerError(err)
	}

	return common.JsonResponse{
		Data:   players,
		Status: 200,
	}
}
