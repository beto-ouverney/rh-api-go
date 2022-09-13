package test

import (
	"bytes"
	"encoding/json"
	"github.com/beto-ouverney/rh-api/entity"
	"github.com/beto-ouverney/rh-api/handler"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestRegistraFuncionario(t *testing.T) {
	assertions := assert.New(t)

	t.Setenv("POSTGRES_USER", "root")
	t.Setenv("POSTGRES_PASSWORD", "password")
	t.Setenv("POSTGRES_DB", "rh_db_test")
	t.Setenv("DB_CONNECTION", "user=root password=password dbname=rh_db_test sslmode=disable")

	// initialize the database if in the test environment
	if strings.Contains(os.Getenv("POSTGRES_DB"), "test") {
		t.Log("Initializing the database for testing")
		initDBTest(t)
		defer dropDBTEST(t)
	} else {
		t.Skip("Skipping test because it is not a test environment, the port number is not 6306")
	}

	router := chi.NewRouter()

	router.Route("/funcionarios", func(r chi.Router) {
		r.Get("/{id}", handler.GetById)
		r.Get("/", handler.GetAllEmployees)
		r.Post("/", handler.RegistraFuncionario)
	})
	var responseStruct struct {
		ID int64 `json:"id"`
	}

	responseStruct.ID = 3

	tests := []struct {
		describe        string
		requestBody     interface{}
		expectedStatus  int
		expectedMessage interface{}
		msg             string
		msg1            string
	}{
		{
			describe: "Deve ser capaz de registrar um funcionário",
			requestBody: entity.Funcionario{
				ID:           1,
				Nome:         "Jessica",
				Sobrenome:    "Ouverney Paz",
				Documento:    "849.918.097-38",
				Setor:        "TI",
				SalarioBruto: 7000,
				DataAdmissao: "30/10/2020",
				Saude:        false,
				Dental:       false,
				Transporte:   false,
				Dependente:   0,
				Pensao:       0,
			},
			expectedStatus:  201,
			expectedMessage: responseStruct,
			msg:             "Status code deve ser igual",
			msg1:            "Deve ser capaz de retornar os dados de uma funcionario de acordo com o id",
		},
	}

	for _, tt := range tests {
		t.Run(tt.describe, func(t *testing.T) {

			data, err := json.Marshal(tt.requestBody)
			if err != nil {
				t.Fatal(err)
			}
			req := httptest.NewRequest(http.MethodPost, "/funcionarios/", bytes.NewBuffer(data))
			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)

			assertions.Equal(tt.expectedStatus, rr.Code, tt.msg)

			var actual struct {
				ID int64 `json:"id"`
			}

			errJ := json.Unmarshal(rr.Body.Bytes(), &actual)
			if errJ != nil {
				t.Errorf("Error unmarshalling response: %v", err)
			}
			assertions.Equal(tt.expectedMessage, actual, tt.msg1)
		})
	}
}
