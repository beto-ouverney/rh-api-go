package handler

import (
	"github.com/beto-ouverney/rh-api/controller/funcionarioscontroller"
	"github.com/go-chi/chi/v5"
	"net/http"
	"net/url"
)

// GetById returns a funcionario by id from controller and send a response to client
func GetById(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	param, errQ := url.QueryUnescape(chi.URLParam(r, "id"))
	if errQ != nil {
		errorReturn(w, 500, errQ.Error())
	}

	c := funcionarioscontroller.New()

	response, err := c.GetByID(r.Context(), param)
	status := http.StatusOK
	if err != nil {
		response, status = errorHandler(err, w)
	}
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	_, errW := w.Write(response)
	if errW != nil {
		errorReturn(w, 500, errW.Error())
	}
}
