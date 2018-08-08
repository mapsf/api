package app

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"time"
	"sync"
	"fmt"
)

func getRouter() *httprouter.Router {
	router := httprouter.New()

	router.GET("/", getRootPathHandler)
	router.GET("/ping", pingHandler)
	router.GET("/hello", helloHandler)
	router.GET("/version", versionHandler)
	router.GET("/high-load", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		var wg sync.WaitGroup
		var goroutines = 1000000
		wg.Add(goroutines)
		for i := 0; i < goroutines; i++ {
			go func() {
				defer wg.Done()
				time.Sleep(5 * time.Second)
			}()
		}
		wg.Wait()
		fmt.Fprintln(w, "Done.")
	})

	router.PanicHandler = panicHandler

	return router
}
