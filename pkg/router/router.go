package router

import (
	"net/http"

	"github.com/hanna3-14/BackTheMiles/pkg/middleware"
)

func Router(audience, domain string) http.Handler {

	router := http.NewServeMux()

	router.HandleFunc("/", middleware.NotFoundHandler)
	router.Handle("/api/messages/results", middleware.ValidateJWT(audience, domain)(http.HandlerFunc(middleware.ResultsApiHandler)))

	return middleware.HandleCacheControl(router)
}
