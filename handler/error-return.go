package handler

import "net/http"

// errorReturn is a function to return a error to the client
func errorReturn(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	_, _ = w.Write([]byte("{\"message\":\"" + message + "\"}"))
}
