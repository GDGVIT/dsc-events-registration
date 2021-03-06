package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/GDGVIT/dsc-events-registration/api/views"
	"github.com/GDGVIT/dsc-events-registration/pkg/participants"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/validator.v2"
)

var AcceptedEvents = [2]string{
	"WomenTechies20",
	"SolutionsChallenge",
	"DevStack20",
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

		// verify recaptcha
		// take captcha token from the request header
		captchaSecret := os.Getenv("CAPTCHA_SECRET")
		captchaToken := r.Header.Get("g-recaptcha-response")
		remoteIP := r.RemoteAddr
		url := fmt.Sprintf("https://www.google.com/recaptcha/api/siteverify?secret=%s&response=%s&remoteip=%s", captchaSecret, captchaToken, remoteIP)

		log.Println(captchaToken)

		resp, err := http.Get(url)
		if err != nil {
			views.Wrap(err, w)
			return
		}
		defer resp.Body.Close()

		var i map[string]interface{}
		if err = json.NewDecoder(resp.Body).Decode(&i); err != nil {
			views.Wrap(err, w)
			return
		}
		if i["success"] == false {
			views.Wrap(views.ErrInvalidCaptcha, w)
			return
		}

		// captcha verified
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

		w.WriteHeader(http.StatusNotFound)
		return
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
