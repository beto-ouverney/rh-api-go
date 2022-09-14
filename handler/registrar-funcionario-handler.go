package handler

import (
	"encoding/json"
	"github.com/beto-ouverney/rh-api/controller/funcionarioscontroller"
	"github.com/beto-ouverney/rh-api/entity"
	"net/http"
)

// RegistraFuncionario receive a funcionario from client and send to controller
func RegistraFuncionario(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var funcionario entity.Funcionario

	errJ := json.NewDecoder(r.Body).Decode(&funcionario)
	if errJ != nil {
		errorReturn(w, 500, errJ.Error())
	}

	c := funcionarioscontroller.New()

	response, err := c.Register(r.Context(), funcionario)
	status := http.StatusCreated
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
