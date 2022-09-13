package handler

import (
	"github.com/beto-ouverney/rh-api/controller/funcionarioscontroller"
	"net/http"
)

// GetAllEmployees returns all funcionarios from controller and send a response to client
func GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	c := funcionarioscontroller.New()
	response, err := c.GetAll(r.Context())
	if err != nil {
		errorHandler(err, w)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_, errW := w.Write(response)
	if errW != nil {
		errorReturn(w, 500, errW.Error())
	}
}
