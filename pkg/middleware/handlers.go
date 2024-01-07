package middleware

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/hanna3-14/BackTheMiles/pkg/data"
	"github.com/hanna3-14/BackTheMiles/pkg/helpers"
	"github.com/hanna3-14/BackTheMiles/pkg/models"
)

const (
	notFoundErrorMessage       = "Not Found"
	internalServerErrorMessage = "Internal Server Error"
)

type ErrorMessage struct {
	Message string `json:"message"`
}

func ResultsDataHandler(w http.ResponseWriter, r *http.Request) {
	results := data.ResultsData()
	if r.Method == http.MethodGet {
		err := helpers.WriteJSON(w, http.StatusOK, results)
		if err != nil {
			ServerError(w, err)
		}
	} else if r.Method == http.MethodPost {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			ServerError(w, err)
		}
		var result models.Result
		err = json.Unmarshal(body, &result)
		if err != nil {
			ServerError(w, err)
		}
		err = data.PostResult(result)
		if err != nil {
			ServerError(w, err)
		}
	} else {
		NotFoundHandler(w, r)
	}
}

func sendGoalsData(w http.ResponseWriter, r *http.Request, data []models.Goal) {
	if r.Method == http.MethodGet {
		err := helpers.WriteJSON(w, http.StatusOK, data)
		if err != nil {
			ServerError(w, err)
		}
	} else {
		NotFoundHandler(w, r)
	}
}

func GoalsDataHandler(w http.ResponseWriter, r *http.Request) {
	sendGoalsData(w, r, data.GoalsData())
}

func HandleCacheControl(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		headers := rw.Header()
		headers.Set("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate")
		headers.Set("Pragma", "no-cache")
		headers.Set("Expires", "0")
		next.ServeHTTP(rw, req)
	})
}

func NotFoundHandler(rw http.ResponseWriter, req *http.Request) {
	errorMessage := ErrorMessage{Message: notFoundErrorMessage}
	err := helpers.WriteJSON(rw, http.StatusNotFound, errorMessage)
	if err != nil {
		ServerError(rw, err)
	}
}

func ServerError(rw http.ResponseWriter, err error) {
	errorMessage := ErrorMessage{Message: internalServerErrorMessage}
	helpers.WriteJSON(rw, http.StatusInternalServerError, errorMessage)
	log.Print("Internal error server: ", err.Error())
}
