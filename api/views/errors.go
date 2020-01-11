package views

import (
	"encoding/json"
	"errors"
	"net/http"

	pkg "github.com/GDGVIT/dsc-events-registration/pkg"
	log "github.com/sirupsen/logrus"
)

type ErrView struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

var (
	ErrMethodNotAllowed = errors.New("Error: Method is not allowed")
	ErrInvalidSlug      = errors.New("Error: Invalid Slug")
)

var ErrHTTPStatusMap = map[string]int{
	ErrMethodNotAllowed.Error(): http.StatusMethodNotAllowed,
	pkg.ErrNotFound.Error():     http.StatusNotFound,
	pkg.ErrInvalidSlug.Error():  http.StatusBadRequest,
	ErrInvalidSlug.Error():      http.StatusBadRequest,
	pkg.ErrExists.Error():       http.StatusConflict,
	pkg.ErrNoContent.Error():    http.StatusNotFound,
	pkg.ErrDatabase.Error():     http.StatusInternalServerError,
	pkg.ErrUnauthorized.Error(): http.StatusUnauthorized,
	pkg.ErrForbidden.Error():    http.StatusForbidden,
}

func Wrap(err error, w http.ResponseWriter) {
	msg := err.Error()
	code := ErrHTTPStatusMap[msg]

	// If error code is not found
	// like a default case
	if code == 0 {
		code = http.StatusInternalServerError
	}

	w.WriteHeader(code)

	errView := ErrView{
		Message: msg,
		Status:  code,
	}
	log.WithFields(log.Fields{
		"message": msg,
		"code":    code,
	}).Error("Error occurred")

	json.NewEncoder(w).Encode(errView)
}
