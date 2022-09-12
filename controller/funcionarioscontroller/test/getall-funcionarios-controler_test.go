package test_test

import (
	"context"
	"encoding/json"
	mockscontroller "github.com/beto-ouverney/rh-api/controller/funcionarioscontroller/mocks"
	"github.com/beto-ouverney/rh-api/customerror"
	"github.com/beto-ouverney/rh-api/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func Test_funcionariosController_GetAll(t *testing.T) {
	assertions := assert.New(t)

	response := []entity.Funcionario{
		{
			ID:           1,
			Nome:         "Beto",
			Sobrenome:    "Ouverney",
			Documento:    "12345678910",
			Setor:        "Desenvolvedor",
			SalarioBruto: 10000.00,
			Saude:        true,
			Dental:       true,
			Transporte:   true,
		},
		{
			ID:           2,
			Nome:         "Maria",
			Sobrenome:    "Joaquina de Amaral Pereira Goes",
			Documento:    "12345678911",
			Setor:        "Desenvolvedor",
			SalarioBruto: 10000.00,
			Saude:        true,
			Dental:       true,
			Transporte:   true,
		},
	}

	responseJson, errj := json.MarshalIndent(response, "", "  ")
	if errj != nil {
		t.Errorf("Erro ao converter para json: %v", errj)
	}

	tests := []struct {
		name  string
		want  []byte
		want1 *customerror.CustomError
		msg   string
		msg1  string
	}{
		{
			name:  "Teste se é possível pegar todos os funcionarios em json",
			want:  responseJson,
			want1: nil,
			msg:   "Deve ser possivel pegar todos os funcionarios em json",
			msg1:  "Erro tem que ser nulo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()

			m := new(mockscontroller.IFuncionariosController)
			m.On("GetAll", mock.AnythingOfType("*context.emptyCtx")).Return(tt.want, tt.want1)

			got, got1 := m.GetAll(ctx)
			assertions.Equal(tt.want, got, tt.msg)
			assertions.Equal(tt.want1, got1, tt.msg1)
		})
	}
}
