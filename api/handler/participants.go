package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/GDGVIT/dsc-events-registration/api/views"
	"github.com/GDGVIT/dsc-events-registration/pkg/participants"
	"github.com/julienschmidt/httprouter"
)

func register(svc participants.Service) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

		p := &participants.Participant{}
		if err := json.NewDecoder(r.Body).Decode(p); err != nil {
			views.Wrap(err, w)
			return
		}
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		id, err := svc.Save(ctx, p)
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

func MakeParticipantHandler(r *httprouter.Router, svc participants.Service) {
	r.GET("/api/v1/participants/ping", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.WriteHeader(http.StatusOK)
		return
	})
	r.POST("/api/v1/participants/register", register(svc))
}
