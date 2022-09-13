package handler

import (
	"github.com/beto-ouverney/rh-api/customerror"
	"log"
	"net/http"
)

// errorHandler handles the error and return a json to the client
func errorHandler(err *customerror.CustomError, w http.ResponseWriter) {
	var status int
	if err.Code == customerror.ENOTFOUND {
		status = 404
	} else if err.Code == customerror.ECONFLICT {
		status = 400
	} else {
		log.Println(err)
		status = 500
	}
	response := []byte("{\"message\":\"" + err.Error() + "\"}")

	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	_, errW := w.Write(response)
	if errW != nil {
		errorReturn(w, 500, errW.Error())
	}

}
