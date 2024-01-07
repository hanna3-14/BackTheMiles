package middleware

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

func ResultsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		results, err := data.GetResults()
		if err != nil {
			ServerError(w, err)
		}
		err = helpers.WriteJSON(w, http.StatusOK, results)
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

func ResultHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		id := mux.Vars(r)["id"]
		result, err := data.GetResultById(id)
		if err != nil {
			ServerError(w, err)
		}
		err = helpers.WriteJSON(w, http.StatusOK, result)
		if err != nil {
			ServerError(w, err)
		}
	} else if r.Method == http.MethodPatch {
		id := mux.Vars(r)["id"]
		body, err := io.ReadAll(r.Body)
		if err != nil {
			ServerError(w, err)
		}
		var result models.Result
		err = json.Unmarshal(body, &result)
		if err != nil {
			ServerError(w, err)
		}
		result.ID, err = strconv.Atoi(id)
		if err != nil {
			ServerError(w, err)
		}
		err = data.PatchResult(id, result)
		if err != nil {
			ServerError(w, err)
		}
	} else if r.Method == http.MethodDelete {
		id := mux.Vars(r)["id"]
		err := data.DeleteResult(id)
		if err != nil {
			ServerError(w, err)
		}
	}
}

func GoalsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		goals, err := data.GetGoals()
		if err != nil {
			ServerError(w, err)
		}
		err = helpers.WriteJSON(w, http.StatusOK, goals)
		if err != nil {
			ServerError(w, err)
		}
	} else if r.Method == http.MethodPost {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			ServerError(w, err)
		}
		var goal models.Goal
		err = json.Unmarshal(body, &goal)
		if err != nil {
			ServerError(w, err)
		}
		err = data.PostGoal(goal)
		if err != nil {
			ServerError(w, err)
		}
	} else {
		NotFoundHandler(w, r)
	}
}

func GoalHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		id := mux.Vars(r)["id"]
		goal, err := data.GetGoalById(id)
		if err != nil {
			ServerError(w, err)
		}
		err = helpers.WriteJSON(w, http.StatusOK, goal)
		if err != nil {
			ServerError(w, err)
		}
	} else if r.Method == http.MethodPatch {
		id := mux.Vars(r)["id"]
		body, err := io.ReadAll(r.Body)
		if err != nil {
			ServerError(w, err)
		}
		var goal models.Goal
		err = json.Unmarshal(body, &goal)
		if err != nil {
			ServerError(w, err)
		}
		goal.ID, err = strconv.Atoi(id)
		if err != nil {
			ServerError(w, err)
		}
		err = data.PatchGoal(id, goal)
		if err != nil {
			ServerError(w, err)
		}
	} else if r.Method == http.MethodDelete {
		id := mux.Vars(r)["id"]
		err := data.DeleteGoal(id)
		if err != nil {
			ServerError(w, err)
		}
	}
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
