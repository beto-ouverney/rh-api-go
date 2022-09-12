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

func Test_funcionariosRepository_Register(t *testing.T) {
	assertions := assert.New(t)

	var responseID int64 = 4

	type args struct {
		funcionario entity.Funcionario
	}
	tests := []struct {
		name  string
		args  args
		want  *int64
		want1 *customerror.CustomError
		msg   string
		msg1  string
	}{
		{
			name: "Deve ser possível registrar um funcionário retornando o ID do mesmo",
			args: args{
				funcionario: entity.Funcionario{
					Nome:         "Beto",
					Sobrenome:    "Ouverney",
					Documento:    "495.369.367-12",
					Setor:        "TI",
					SalarioBruto: 7000.00,
					DataAdmissao: "2021-01-01",
					Saude:        false,
					Transporte:   false,
					Dental:       false,
					Dependente:   0,
					Pensao:       0,
				},
			},
			want:  &responseID,
			want1: nil,
			msg:   "Funcionário deve ser registrado com sucesso",
			msg1:  "O erro deve ser nulo",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()

			m := new(mocks.IFuncionariosRepository)
			m.On("Register", mock.AnythingOfType("*context.emptyCtx"), tt.args.funcionario).Return(tt.want, tt.want1)

			got, got1 := m.Register(ctx, tt.args.funcionario)
			assertions.Equal(tt.want, got, tt.msg)
			assertions.Equal(tt.want1, got1, tt.msg1)
		})
	}

}
