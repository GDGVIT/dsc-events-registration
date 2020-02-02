package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/GDGVIT/dsc-events-registration/api/views"
	"github.com/GDGVIT/dsc-events-registration/pkg/participants"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/validator.v2"
)

var AcceptedEvents = [1]string{
	"WomenTechies20",
	"SolutionsChallenge",
}

func isEventAccepted(event string) bool {
	for _, v := range AcceptedEvents {
		if v == event {
			return true
		}
	}
	return false
}

func register(svc participants.Service) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

		p := &participants.Participant{}
		if err := json.NewDecoder(r.Body).Decode(p); err != nil {
			views.Wrap(err, w)
			return
		}

		if errs := validator.Validate(p); errs != nil {
			views.Wrap(views.ErrInvalidSlug, w)
			return
		}
		if !isEventAccepted(p.EventName) {
			log.Println("Not accepted event")
			views.Wrap(views.ErrInvalidSlug, w)
			return
		}

		id, err := svc.Save(r.Context(), p)
		if err != nil {
			views.Wrap(err, w)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"ID": id,
		})
	}
}

func viewCount(svc participants.Service) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		eventName := r.URL.Query().Get("event")

		if eventName == "" {
			data, err := svc.CountParticipantsByEvents(r.Context())
			if err != nil {
				views.Wrap(err, w)
				return
			}
			if err = json.NewEncoder(w).Encode(data); err != nil {
				views.Wrap(err, w)
				return
			}
			return
		}
		count, err := svc.CountParticipantsByEvent(r.Context(), eventName)
		if err != nil {
			views.Wrap(err, w)
			return
		}
		if err = json.NewEncoder(w).Encode(map[string]interface{}{
			"eventName":         eventName,
			"registrationCount": count,
		}); err != nil {
			views.Wrap(err, w)
			return
		}
	}
}

func MakeParticipantHandler(r *httprouter.Router, svc participants.Service) {
	r.GET("/api/v1/participants/ping", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.WriteHeader(http.StatusOK)
		return
	})
	r.POST("/api/v1/participants/register", register(svc))
	r.GET("/api/v1/participants/count", viewCount(svc))
}
