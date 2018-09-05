package common

import (
	"net/http"
	"encoding/json"
	"github.com/mapsf/api/api/app/models"
)

type Params struct {
	*http.Request
	http.ResponseWriter
	User *models.Character
}

type responseError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (p Params) JSON(data interface{}, statusCode int) {
	w := p.ResponseWriter
	err, ok := data.(error)
	if ok {
		data = responseError{Message: err.Error(), Code: statusCode}
	}
	w.WriteHeader(statusCode)
	writeJson(w, data)
}

func writeJson(w http.ResponseWriter, data interface{}) {
	bytes, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
	} else {
		w.Write(bytes)
	}
}

type MyHandler func(Params) ResponseRenderer

// передаем в хендлер новые параметры и обрабатыает возвращаемый с функции json
func GetHandlerFunc(handler MyHandler) http.HandlerFunc {
	// возвращаем функцию что будет вызывается как при обработке запроса
	return func(w http.ResponseWriter, r *http.Request) {
		var params = Params{Request: r, ResponseWriter: w}
		var response = handler(params)
		if response != nil {
			response.Render(params)
		}
	}
}

type ResponseRenderer interface {
	Render(p Params)
}

type JsonResponse struct {
	Data   interface{}
	Status int
}

func (response JsonResponse) Render(p Params) {
	p.JSON(response.Data, response.Status)
}

type messageResponse struct {
	Message string `json:"message"`
}

func JsonMessageResponse(message string) JsonResponse {
	return JsonResponse{
		Data:   messageResponse{Message: message},
		Status: 200,
	}
}

func JsonErrorMessageResponse(message string, code int) JsonResponse {
	return JsonResponse{
		Data:   messageResponse{Message: message},
		Status: code,
	}
}

func JsonMapResponse(data interface{}) JsonResponse {
	return JsonResponse{
		Data:   data,
		Status: 200,
	}
}

func ServerError(err error) JsonResponse {
	return JsonErrorMessageResponse(err.Error(), 500)
}
