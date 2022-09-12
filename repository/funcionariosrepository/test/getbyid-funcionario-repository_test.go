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

func Test_funcionariosRepository_GetByID(t *testing.T) {
	assertions := assert.New(t)

	type args struct {
		id string
	}
	tests := []struct {
		name  string
		args  args
		want  *entity.Funcionario
		want1 *customerror.CustomError
		msg   string
		msg1  string
	}{
		{
			name: "Deve ser possível buscar um funcionário pelo ID",
			args: args{
				id: "1",
			},
			want: &entity.Funcionario{
				ID:           1,
				Nome:         "Beto",
				Sobrenome:    "Ouverney",
				Documento:    "12345678910",
				Setor:        "TI",
				SalarioBruto: 7000.00,
				DataAdmissao: "2021-01-01",
				Saude:        false,
				Transporte:   false,
				Dental:       false,
				Dependente:   0,
				Pensao:       0,
			},
			want1: nil,
			msg:   "Funcionário deve ser encontrado com o ID informado",
			msg1:  "O erro deve ser nulo",
		},
		{
			name: "Não deve encontrar um funcionário com ID inexistente",
			args: args{
				id: "2",
			},
			want:  nil,
			want1: nil,
			msg:   "Funcionário não deve ser encontrado",
			msg1:  "O erro deve ser nulo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()

			m := new(mocks.IFuncionariosRepository)
			m.On("GetByID", mock.AnythingOfType("*context.emptyCtx"), tt.args.id).Return(tt.want, nil)

			got, got1 := m.GetByID(ctx, tt.args.id)
			assertions.Equal(tt.want, got, tt.msg)
			assertions.Equal(tt.want1, got1, tt.msg1)
		})
	}
}
