package test_test

import (
	"github.com/beto-ouverney/rh-api/customerror"
	"github.com/beto-ouverney/rh-api/entity"
	"github.com/beto-ouverney/rh-api/repository/funcionariosrepository/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_funcionariosRepository_SetCacheGetAllEmployees(t *testing.T) {
	assertions := assert.New(t)

	type args struct {
		funcionarios *[]entity.Funcionario
	}
	tests := []struct {
		name string
		args args
		want *customerror.CustomError
		msg  string
	}{
		{
			name: "Deve ser possível salvar os funcionários no cache",
			args: args{
				funcionarios: &[]entity.Funcionario{
					{
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
					{
						ID:           2,
						Nome:         "Renata",
						Sobrenome:    "Targaryen",
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
				},
			},
			want: nil,
			msg:  "Funcionários devem ser salvos no cache e o retorno deve ser nulo",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			m := new(mocks.IFuncionariosRepository)
			m.On("SetCacheGetAllEmployees", tt.args.funcionarios).Return(tt.want, nil)

			got := m.SetCacheGetAllEmployees(tt.args.funcionarios)
			assertions.Equal(tt.want, got, tt.msg)
		})
	}
}
