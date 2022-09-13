package handler

import (
	"github.com/beto-ouverney/rh-api/controller/contrachequecontroller"
	"github.com/go-chi/chi/v5"
	"net/http"
	"net/url"
)

// GetByFuncIDContrachequeHandler returns a payslip by funcionario id from controller and send a response to client
func GetByFuncIDContrachequeHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	param, errP := url.QueryUnescape(chi.URLParam(r, "id"))
	if errP != nil {
		errorReturn(w, 500, errP.Error())
	}

	c := contrachequecontroller.New()
	response, err := c.GetByFuncionarioID(r.Context(), param)
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
