package app

import (
	"net/http"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"os"
	"runtime"
)

func getRootPathHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "This is the RESTful API")
}

func pingHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "pong...")
}

func helloHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Hello %s", os.Getenv("NAME"))
}

func versionHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "version: %s", runtime.Version())
}

func panicHandler(w http.ResponseWriter, r *http.Request, err interface{}) {
	fmt.Fprintf(w, "Паника: %s", err)
}
