package test_test

import (
	"context"
	"github.com/beto-ouverney/rh-api/customerror"
	"github.com/beto-ouverney/rh-api/entity"
	"github.com/beto-ouverney/rh-api/repository/funcionariosrepository/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func Test_funcionariosRepository_GetAll(t *testing.T) {
	assertions := assert.New(t)

	tests := []struct {
		name  string
		want  *[]entity.Funcionario
		want1 *customerror.CustomError
		msg   string
		msg1  string
	}{
		{
			name: "Teste se é possível pegar todos os funcionarios",
			want: &[]entity.Funcionario{
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
			},
			want1: nil,
			msg:   "Deve ser possivel pegar todos os funcionarios",
			msg1:  "Erro tem que ser nulo",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()

			m := new(mocks.IFuncionariosRepository)
			m.On("GetAll", mock.AnythingOfType("*context.emptyCtx")).Return(tt.want, nil)

			got, got1 := m.GetAll(ctx)
			assertions.Equal(tt.want, got, tt.msg)
			assertions.Equal(tt.want1, got1, tt.msg1)
		})
	}
}
