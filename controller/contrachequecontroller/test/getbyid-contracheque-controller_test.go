package test

import (
	"encoding/json"
	contrachequemocks "github.com/beto-ouverney/rh-api/controller/contrachequecontroller/mocks"
	"github.com/beto-ouverney/rh-api/customerror"
	"github.com/beto-ouverney/rh-api/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_contrachequeController_GetByID(t *testing.T) {
	assertions := assert.New(t)

	contracheque := &entity.Contracheque{
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
	}

	response, errJ := json.MarshalIndent(&contracheque, "", "  ")
	if errJ != nil {
		t.Errorf("Erro ao converter contracheque para json: %v", errJ)
	}

	type args struct {
		id string
	}
	tests := []struct {
		name  string
		args  args
		want  []byte
		want1 *customerror.CustomError
		msg   string
		msg1  string
	}{
		{
			name: "Deve ser possível buscar o contracheque de um funcionário e retorna-lo no formato json",
			args: args{
				id: "1",
			},
			want:  response,
			want1: nil,
			msg:   "Contracheque deve ser encontrado com o ID do funcionário informado",
			msg1:  "O erro deve ser nulo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			m := new(contrachequemocks.IContrachequeController)
			m.On("GetByFuncionarioID", "*context.emptyCtx", tt.args.id).Return(tt.want, tt.want1)

			got, got1 := m.GetByFuncionarioID("*context.emptyCtx", tt.args.id)
			assertions.Equal(tt.want, got, tt.msg)
			assertions.Equal(tt.want1, got1, tt.msg1)
		})
	}
}
