package test_test

import (
	"github.com/beto-ouverney/rh-api/customerror"
	"github.com/beto-ouverney/rh-api/repository/funcionariosrepository/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_funcionariosRepository_GetAllCache(t *testing.T) {
	assertions := assert.New(t)

	response := "MEMORY"

	tests := []struct {
		name  string
		want  *string
		want1 *customerror.CustomError
		msg   string
		msg1  string
	}{
		{
			name:  "Deve ser possível buscar todos os funcionários no cache",
			want:  &response,
			want1: nil,
			msg:   "Deve retornar dados do cache",
			msg1:  "O erro deve ser nulo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			m := new(mocks.IFuncionariosRepository)
			m.On("GetAllCache").Return(tt.want, tt.want1)

			got, got1 := m.GetAllCache()
			assertions.Equal(tt.want, got, tt.msg)
			assertions.Equal(tt.want1, got1, tt.msg1)
		})
	}
}
