package controllers

import (
	"github.com/mapsf/api/app/ws"
	"strconv"
	"github.com/mapsf/api/app/common"
	"github.com/gorilla/context"
	"github.com/mapsf/api/app/models"
)

func GetRootPathHandler(p common.Params) common.ResponseRenderer {
	user := context.Get(p.Request, "user").(models.Character)
	return common.JsonResponse{
		Data:   map[string]string{"welcome": "This is the RESTful API", "login": user.Login},
		Status: 200,
	}
}

func GetClients(p common.Params) common.ResponseRenderer {
	return common.JsonMessageResponse(strconv.Itoa(len(ws.Hub.Connections)))
}
