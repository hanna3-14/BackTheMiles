package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hanna3-14/BackTheMiles/pkg/middleware"
)

func Router(audience, domain string) http.Handler {

	router := mux.NewRouter()

	router.HandleFunc("/", middleware.NotFoundHandler)
	router.Handle("/api/data/results", middleware.ValidateJWT(audience, domain)(http.HandlerFunc(middleware.ResultsDataHandler)))
	router.Handle("/api/data/result/{id:[0-9]+}", middleware.ValidateJWT(audience, domain)(http.HandlerFunc(middleware.ResultHandler)))
	router.Handle("/api/data/goals", middleware.ValidateJWT(audience, domain)(http.HandlerFunc(middleware.GoalsDataHandler)))

	return middleware.HandleCacheControl(router)
}
