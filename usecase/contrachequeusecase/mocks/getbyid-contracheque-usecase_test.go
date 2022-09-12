package usecasemocks

import (
	"context"
	"github.com/beto-ouverney/rh-api/customerror"
	"github.com/beto-ouverney/rh-api/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func Test_contrachequeUseCase_GetByFuncionarioID(t *testing.T) {
	assertions := assert.New(t)

	type args struct {
		employeeID string
	}
	tests := []struct {
		name  string
		args  args
		want  *entity.Contracheque
		want1 *customerror.CustomError
		msg   string
		msg1  string
	}{
		{
			name: "Deve ser possível buscar o contracheque de um funcionário",
			args: args{
				employeeID: "1",
			},
			want: &entity.Contracheque{
				MesReferencia: "01/2021",
				FuncionarioID: 2,
				Nome:          "Beto",
				Documento:     "12345678910",
				Setor:         "TI",
				SalarioBruto:  7000.00,
				Lancamentos: []entity.Lancamento{
					{
						Tipo:      "D",
						Valor:     560.00,
						Descricao: "FGTS",
					},
					{
						Tipo:      "D",
						Valor:     816.12,
						Descricao: "INSS",
					},
					{
						Tipo:      "D",
						Valor:     831.19,
						Descricao: "IR",
					},
				},
			},
			want1: nil,
			msg:   "Contracheque deve ser encontrado com o ID do funcionário informado",
			msg1:  "O erro deve ser nulo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()

			m := new(IContrachequeUseCase)
			m.On("GetByFuncionarioID", mock.AnythingOfType("*context.emptyCtx"), tt.args.employeeID).Return(tt.want, tt.want1)

			got, got1 := m.GetByFuncionarioID(ctx, tt.args.employeeID)
			assertions.Equal(tt.want, got, tt.msg)
			assertions.Equal(tt.want1, got1, tt.msg1)
		})
	}
}
