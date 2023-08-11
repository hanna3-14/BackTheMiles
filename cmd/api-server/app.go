package main

import (
	"log"
	"net/http"

	"github.com/hanna3-14/BackTheMiles/pkg/router"
	"github.com/rs/cors"
	"github.com/unrolled/secure"
)

type Config struct {
	Port          string
	SecureOptions secure.Options
	CorsOptions   cors.Options
	Audience      string
	Domain        string
}

type App struct {
	Config Config
}

func (app *App) RunServer() {
	router := router.Router(app.Config.Audience, app.Config.Domain)
	corsMiddleware := cors.New(app.Config.CorsOptions)
	routerWithCORS := corsMiddleware.Handler(router)

	secureMiddleware := secure.New(app.Config.SecureOptions)
	finalHandler := secureMiddleware.Handler(routerWithCORS)

	server := &http.Server{
		Addr:    ":" + app.Config.Port,
		Handler: finalHandler,
	}

	log.Printf("API server listening on %s", server.Addr)
	log.Fatal(server.ListenAndServe())
}
