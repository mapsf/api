package controllers

import (
	"github.com/mapsf/api/app/common"
)

func AttackHandler(p common.Params) common.ResponseRenderer {

	var (
		_ = p.URL.Query().Get("id")
	)

	return common.JsonResponse{Data: p.User, Status: 200}
}
