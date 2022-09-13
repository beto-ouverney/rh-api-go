package test

import (
	"encoding/json"
	"fmt"
	"github.com/beto-ouverney/rh-api/entity"
	"github.com/beto-ouverney/rh-api/handler"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"
)

func GetByFuncIDContrachequeHandler(t *testing.T) {
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

	router.Route("/contracheque", func(r chi.Router) {
		r.Get("/{id}", handler.GetByFuncIDContrachequeHandler)
	})

	tests := []struct {
		describe        string
		id              string
		expectedStatus  int
		expectedMessage interface{}
		msg             string
		msg1            string
	}{
		{
			describe:        "Deve ser capaz de o contracheque do funcionario de acordo com o id",
			id:              "1",
			expectedStatus:  200,
			expectedMessage: allFuncMock[0],
			msg:             "Status code deve ser igual",
			msg1:            "Deve ser capaz de o contracheque do funcionario de acordo com o id",
		},
	}
	for _, tt := range tests {
		t.Run(tt.describe, func(t *testing.T) {
			path := fmt.Sprintf("/contracheque/%s", url.QueryEscape(tt.id))
			req := httptest.NewRequest(http.MethodGet, path, nil)
			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)

			assertions.Equal(tt.expectedStatus, rr.Code, tt.msg)
			var actual entity.Contracheque
			err := json.Unmarshal(rr.Body.Bytes(), &actual)
			if err != nil {
				t.Errorf("Error unmarshalling response: %v", err)
			}

			t.Log(actual)
			assertions.Equal(tt.expectedMessage, actual, tt.msg1)
		})
	}
}
