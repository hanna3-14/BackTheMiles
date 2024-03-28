package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hanna3-14/BackTheMiles/pkg/middleware"
)

func Router(audience, domain string) http.Handler {

	router := mux.NewRouter()

	router.HandleFunc("/", middleware.NotFoundHandler)
	router.Handle("/api/data/results", middleware.ValidateJWT(audience, domain)(http.HandlerFunc(middleware.ResultsHandler)))
	router.Handle("/api/data/result/{id:[0-9]+}", middleware.ValidateJWT(audience, domain)(http.HandlerFunc(middleware.ResultHandler)))
	router.Handle("/api/data/goals", middleware.ValidateJWT(audience, domain)(http.HandlerFunc(middleware.GoalsHandler)))
	router.Handle("/api/data/goal/{id:[0-9]+}", middleware.ValidateJWT(audience, domain)(http.HandlerFunc(middleware.GoalHandler)))
	router.Handle("/api/data/events", middleware.ValidateJWT(audience, domain)(http.HandlerFunc(middleware.EventsHandler)))
	router.Handle("/api/data/event/{id:[0-9]+}", middleware.ValidateJWT(audience, domain)(http.HandlerFunc(middleware.EventHandler)))
	router.Handle("/api/data/distances", middleware.ValidateJWT(audience, domain)(http.HandlerFunc(middleware.DistancesHandler)))
	router.Handle("/api/data/distance/{id:[0-9]+}", middleware.ValidateJWT(audience, domain)(http.HandlerFunc(middleware.DistanceHandler)))

	return middleware.HandleCacheControl(router)
}
