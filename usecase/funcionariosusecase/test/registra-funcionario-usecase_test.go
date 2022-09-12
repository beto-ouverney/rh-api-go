package test

import (
	"context"
	"errors"
	"github.com/beto-ouverney/rh-api/customerror"
	"github.com/beto-ouverney/rh-api/entity"
	mocksusecase "github.com/beto-ouverney/rh-api/usecase/funcionariosusecase/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func Test_funcionariosUseCase_Register(t *testing.T) {
	assertions := assert.New(t)

	var responseID int64 = 10

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
		{
			name: "Nao deve ser possível registrar um funcionário com data de admissão com ano inválido",
			args: args{
				funcionario: entity.Funcionario{
					Nome:         "Beto",
					Sobrenome:    "Ouverney",
					Documento:    "495.369.367-12",
					Setor:        "TI",
					SalarioBruto: 7000.00,
					DataAdmissao: "2021-31-31",
					Saude:        false,
					Transporte:   false,
					Dental:       false,
					Dependente:   0,
					Pensao:       0,
				},
			},
			want:  nil,
			want1: customerror.NewError(customerror.ECONFLICT, "Ano inválido", "usecase.validaData", errors.New("ano inválido"), nil),
			msg:   "Funcionário não deve ser registrado",
			msg1:  "O erro deve ser de ano inválido",
		},
		{
			name: "Não deve ser possível registrar um funcionário com nome inválido",
			args: args{
				funcionario: entity.Funcionario{
					Nome:         "",
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
			want:  nil,
			want1: customerror.NewError(customerror.ECONFLICT, "Nome inválido", "usecase.validaNome", errors.New("nome inválido"), nil),
			msg:   "Funcionário não deve ser registrado",
			msg1:  "O erro deve ser de nome inválido",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()

			m := new(mocksusecase.IFuncionariosUseCase)
			m.On("Register", mock.AnythingOfType("*context.emptyCtx"), tt.args.funcionario).Return(tt.want, tt.want1)

			got, got1 := m.Register(ctx, tt.args.funcionario)
			assertions.Equal(tt.want, got, tt.msg)
			assertions.Equal(tt.want1, got1, tt.msg1)
		})
	}
}
