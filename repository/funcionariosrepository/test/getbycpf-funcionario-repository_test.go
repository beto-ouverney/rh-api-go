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

func Test_funcionariosRepository_GetByCPF(t *testing.T) {
	assertions := assert.New(t)

	type args struct {
		cpf string
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
			name: "É possivel buscar um funcionario pelo CPF",
			args: args{
				cpf: "12345678910",
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
			msg:   "Funcionario deve ser encontrado com o CPF informado",
			msg1:  "O erro deve ser nulo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()

			m := new(mocks.IFuncionariosRepository)
			m.On("GetByCPF", mock.AnythingOfType("*context.emptyCtx"), tt.args.cpf).Return(tt.want, tt.want1)

			got, got1 := m.GetByCPF(ctx, tt.args.cpf)
			assertions.Equal(tt.want, got, tt.msg)
			assertions.Equal(tt.want1, got1, tt.msg1)
		})
	}
}
