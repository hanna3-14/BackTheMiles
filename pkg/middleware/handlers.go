package middleware

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hanna3-14/BackTheMiles/pkg/application"
	"github.com/hanna3-14/BackTheMiles/pkg/domain"
	"github.com/hanna3-14/BackTheMiles/pkg/helpers"
)

const (
	notFoundErrorMessage       = "Not Found"
	internalServerErrorMessage = "Internal Server Error"
)

type ErrorMessage struct {
	Message string `json:"message"`
}

type HTTPServer struct {
	ResultRequestService   application.ResultRequestService
	GoalRequestService     application.GoalRequestService
	EventRequestService    application.EventRequestService
	DistanceRequestService application.DistanceRequestService
}

func NewHTTPServer(
	resultRequestService application.ResultRequestService,
	goalRequestService application.GoalRequestService,
	eventRequestService application.EventRequestService,
	distanceRequestService application.DistanceRequestService) (*HTTPServer, error) {
	return &HTTPServer{
		ResultRequestService:   resultRequestService,
		GoalRequestService:     goalRequestService,
		EventRequestService:    eventRequestService,
		DistanceRequestService: distanceRequestService,
	}, nil
}

func (s *HTTPServer) ResultsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		results, err := s.ResultRequestService.GetResults()
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
		var result domain.Result
		err = json.Unmarshal(body, &result)
		if err != nil {
			ServerError(w, err)
		}
		err = s.ResultRequestService.PostResult(result)
		if err != nil {
			ServerError(w, err)
		}
	} else {
		NotFoundHandler(w, r)
	}
}

func (s *HTTPServer) ResultHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			ServerError(w, err)
		}
		result, err := s.ResultRequestService.GetResult(id)
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
		var result domain.Result
		err = json.Unmarshal(body, &result)
		if err != nil {
			ServerError(w, err)
		}
		result.ResultID, err = strconv.Atoi(id)
		if err != nil {
			ServerError(w, err)
		}
		err = s.ResultRequestService.PatchResult(result.ResultID, result)
		if err != nil {
			ServerError(w, err)
		}
	} else if r.Method == http.MethodDelete {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			ServerError(w, err)
		}
		err = s.ResultRequestService.DeleteResult(id)
		if err != nil {
			ServerError(w, err)
		}
	}
}

func (s *HTTPServer) GoalsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		goals, err := s.GoalRequestService.GetGoals()
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
		var goal domain.Goal
		err = json.Unmarshal(body, &goal)
		if err != nil {
			ServerError(w, err)
		}
		err = s.GoalRequestService.PostGoal(goal)
		if err != nil {
			ServerError(w, err)
		}
	} else {
		NotFoundHandler(w, r)
	}
}

func (s *HTTPServer) GoalHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			ServerError(w, err)
		}
		goal, err := s.GoalRequestService.GetGoal(id)
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
		var goal domain.Goal
		err = json.Unmarshal(body, &goal)
		if err != nil {
			ServerError(w, err)
		}
		goal.ID, err = strconv.Atoi(id)
		if err != nil {
			ServerError(w, err)
		}
		err = s.GoalRequestService.PatchGoal(goal.ID, goal)
		if err != nil {
			ServerError(w, err)
		}
	} else if r.Method == http.MethodDelete {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			ServerError(w, err)
		}
		err = s.GoalRequestService.DeleteGoal(id)
		if err != nil {
			ServerError(w, err)
		}
	}
}

func (s *HTTPServer) EventsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		events, err := s.EventRequestService.GetEvents()
		if err != nil {
			ServerError(w, err)
		}
		err = helpers.WriteJSON(w, http.StatusOK, events)
		if err != nil {
			ServerError(w, err)
		}
	} else if r.Method == http.MethodPost {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			ServerError(w, err)
		}
		var event domain.Event
		err = json.Unmarshal(body, &event)
		if err != nil {
			ServerError(w, err)
		}
		err = s.EventRequestService.PostEvent(event)
		if err != nil {
			ServerError(w, err)
		}
	} else {
		NotFoundHandler(w, r)
	}
}

func (s *HTTPServer) EventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			ServerError(w, err)
		}
		event, err := s.EventRequestService.GetEvent(id)
		if err != nil {
			ServerError(w, err)
		}
		err = helpers.WriteJSON(w, http.StatusOK, event)
		if err != nil {
			ServerError(w, err)
		}
	} else if r.Method == http.MethodPatch {
		id := mux.Vars(r)["id"]
		body, err := io.ReadAll(r.Body)
		if err != nil {
			ServerError(w, err)
		}
		var event domain.Event
		err = json.Unmarshal(body, &event)
		if err != nil {
			ServerError(w, err)
		}
		event.ID, err = strconv.Atoi(id)
		if err != nil {
			ServerError(w, err)
		}
		err = s.EventRequestService.PatchEvent(event.ID, event)
		if err != nil {
			ServerError(w, err)
		}
	} else if r.Method == http.MethodDelete {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			ServerError(w, err)
		}
		err = s.EventRequestService.DeleteEvent(id)
		if err != nil {
			ServerError(w, err)
		}
	}
}

func (s *HTTPServer) DistancesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		distances, err := s.DistanceRequestService.GetDistances()
		if err != nil {
			ServerError(w, err)
		}
		err = helpers.WriteJSON(w, http.StatusOK, distances)
		if err != nil {
			ServerError(w, err)
		}
	} else if r.Method == http.MethodPost {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			ServerError(w, err)
		}
		var distance domain.Distance
		err = json.Unmarshal(body, &distance)
		if err != nil {
			ServerError(w, err)
		}
		err = s.DistanceRequestService.PostDistance(distance)
		if err != nil {
			ServerError(w, err)
		}
	} else {
		NotFoundHandler(w, r)
	}
}

func (s *HTTPServer) DistanceHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			ServerError(w, err)
		}
		distance, err := s.DistanceRequestService.GetDistance(id)
		if err != nil {
			ServerError(w, err)
		}
		err = helpers.WriteJSON(w, http.StatusOK, distance)
		if err != nil {
			ServerError(w, err)
		}
	} else if r.Method == http.MethodPatch {
		id := mux.Vars(r)["id"]
		body, err := io.ReadAll(r.Body)
		if err != nil {
			ServerError(w, err)
		}
		var distance domain.Distance
		err = json.Unmarshal(body, &distance)
		if err != nil {
			ServerError(w, err)
		}
		distance.ID, err = strconv.Atoi(id)
		if err != nil {
			ServerError(w, err)
		}
		err = s.DistanceRequestService.PatchDistance(distance.ID, distance)
		if err != nil {
			ServerError(w, err)
		}
	} else if r.Method == http.MethodDelete {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			ServerError(w, err)
		}
		err = s.DistanceRequestService.DeleteDistance(id)
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
