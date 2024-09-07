package router

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/hanna3-14/BackTheMiles/pkg/adapters/db"
	"github.com/hanna3-14/BackTheMiles/pkg/application"
	"github.com/hanna3-14/BackTheMiles/pkg/middleware"
)

func Router(audience, domain string) http.Handler {

	dbAdapter, err := db.NewSQLDBAdapter("/marathon.db")
	if err != nil {
		log.Fatal("Database not found")
	}

	resultRequestService, err := application.NewResultRequestService(dbAdapter, dbAdapter)
	if err != nil {
		log.Fatal("Result request service not found")
	}
	goalRequestService, err := application.NewGoalRequestService(dbAdapter)
	if err != nil {
		log.Fatal("Goal request service not found")
	}
	eventRequestService, err := application.NewEventRequestService(dbAdapter, dbAdapter)
	if err != nil {
		log.Fatal("Event request service not found")
	}
	distanceRequestService, err := application.NewDistanceRequestService(dbAdapter)
	if err != nil {
		log.Fatal("Distance request service not found")
	}

	httpServer, err := middleware.NewHTTPServer(resultRequestService, goalRequestService, eventRequestService, distanceRequestService)
	if err != nil {
		log.Fatal("Error creating server")
	}

	router := mux.NewRouter()
	router.HandleFunc("/", middleware.NotFoundHandler)
	router.Handle("/api/data/results", middleware.ValidateJWT(audience, domain)(http.HandlerFunc(httpServer.ResultsHandler)))
	router.Handle("/api/data/result/{id:[0-9]+}", middleware.ValidateJWT(audience, domain)(http.HandlerFunc(httpServer.ResultHandler)))
	router.Handle("/api/data/goals", middleware.ValidateJWT(audience, domain)(http.HandlerFunc(httpServer.GoalsHandler)))
	router.Handle("/api/data/goal/{id:[0-9]+}", middleware.ValidateJWT(audience, domain)(http.HandlerFunc(httpServer.GoalHandler)))
	router.Handle("/api/data/events", middleware.ValidateJWT(audience, domain)(http.HandlerFunc(httpServer.EventsHandler)))
	router.Handle("/api/data/event/{id:[0-9]+}", middleware.ValidateJWT(audience, domain)(http.HandlerFunc(httpServer.EventHandler)))
	router.Handle("/api/data/distances", middleware.ValidateJWT(audience, domain)(http.HandlerFunc(httpServer.DistancesHandler)))
	router.Handle("/api/data/distance/{id:[0-9]+}", middleware.ValidateJWT(audience, domain)(http.HandlerFunc(httpServer.DistanceHandler)))

	return middleware.HandleCacheControl(router)
}
