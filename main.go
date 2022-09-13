package main

import (
	"github.com/beto-ouverney/rh-api/handler"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"os"
)

func main() {
	r := chi.NewRouter()

	r.Route("/funcionarios", func(r chi.Router) {
		r.Get("/{id}", handler.GetById)
		r.Get("/", handler.GetAllEmployees)
		r.Post("/", handler.RegistraFuncionario)
	})

	r.Route("/contracheque", func(r chi.Router) {
		r.Get("/{id}", handler.GetByFuncIDContrachequeHandler)
	})

	port := os.Getenv("PORT")
	log.Println("Server running on port " + port)
	http.ListenAndServe(port, r)
}
