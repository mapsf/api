package app

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/mapsf/api/api/app/controllers"
	"github.com/mapsf/api/api/app/ws"
	"github.com/mapsf/api/api/app/common"
	"github.com/mapsf/api/api/app/auth"
)

func tryFindToken(r *http.Request) (token string) {

	if len(r.Header.Get("Authorization")) > 0 {
		token = r.Header.Get("Authorization")
	} else if len(r.URL.Query().Get("Authorization")) > 0 {
		token = r.URL.Query().Get("Authorization")
	}

	return token
}

func validateJwtMiddleware(next common.MyHandler) http.HandlerFunc {
	return common.GetHandlerFunc(func(params common.Params) common.ResponseRenderer {

		var (
			token = tryFindToken(params.Request)
		)

		if len(token) == 0 {
			return common.JsonErrorMessageResponse("токен не найден", 401)
		}

		user, err := auth.ValidateJwt(token)
		if err != nil {
			return common.JsonErrorMessageResponse(err.Error(), 401)
		}

		params.User = user

		return next(params)
	})
}

func getRouter() *mux.Router {

	r := mux.NewRouter()

	r.HandleFunc("/", validateJwtMiddleware(controllers.GetRootPathHandler)).Methods("GET")

	r.HandleFunc("/auth", common.GetHandlerFunc(controllers.AuthHandler)).Methods("POST")
	r.HandleFunc("/clients", common.GetHandlerFunc(controllers.GetClients)).Methods("GET")

	r.HandleFunc("/attack", validateJwtMiddleware(controllers.AttackHandler)).Methods("POST")
	r.HandleFunc("/character", validateJwtMiddleware(controllers.GetCharacter)).Methods("GET")
	r.HandleFunc("/online-players", validateJwtMiddleware(controllers.GetOnlinePlayers)).Methods("GET")
	r.HandleFunc("/me", validateJwtMiddleware(controllers.Me)).Methods("GET")

	// роут проверки действительный токен или нет
	r.HandleFunc("/validate-token", validateJwtMiddleware(controllers.ValidateToken)).Methods("POST")

	r.HandleFunc("/io", validateJwtMiddleware(ws.Handler))

	http.Handle("/", r)

	return r
}
