package funcionariosusecase

import (
	"context"
	"errors"
	"github.com/beto-ouverney/rh-api/customerror"
	"github.com/beto-ouverney/rh-api/entity"
	"github.com/beto-ouverney/rh-api/helpers/cpf"
	"strings"
)

// validaData validate data
func validaData(dataAdmissao string) *customerror.CustomError {
	data := strings.Split(dataAdmissao, "/")
	if len(data[0]) < 1 || len(data[0]) > 30 {
		return customerror.NewError(customerror.ECONFLICT, "Dia inválido", "usecase.validaData", errors.New("dia inválido"), nil)
	}
	if len(data[1]) < 1 || len(data[1]) > 12 {
		return customerror.NewError(customerror.ECONFLICT, "Mês inválido", "usecase.validaData", errors.New("mês inválido"), nil)
	}

	if len(data[2]) < 4 || len(data[2]) > 4 {
		return customerror.NewError(customerror.ECONFLICT, "Ano inválido", "usecase.validaData", errors.New("ano inválido"), nil)
	}

	return nil
}

// validaFuncionario validate funcionario
func validaFuncionario(f entity.Funcionario) *customerror.CustomError {
	if f.Nome == "" {
		return customerror.NewError(customerror.ECONFLICT, "Nome não pode ser vazio", "usecase.validaFuncionario", errors.New("nome não pode ser vazio"), nil)
	}

	if f.Sobrenome == "" {
		return customerror.NewError(customerror.ECONFLICT, "Sobrenome não pode ser vazio", "usecase.validaFuncionario", errors.New("sobrenome não pode ser vazio"), nil)
	}

	if f.Setor == "" {
		return customerror.NewError(customerror.ECONFLICT, "Setor não pode ser vazio", "usecase.validaFuncionario", errors.New("setor não pode ser vazio"), nil)
	}

	if !cpf.IsValid(f.Documento) {
		return customerror.NewError(customerror.ECONFLICT, "CPF inválido", "usecase.validaFuncionario", errors.New("CPF inválido"), nil)
	}

	if f.SalarioBruto <= 0 {
		return customerror.NewError(customerror.ECONFLICT, "Salário bruto deve ser maior que zero", "usecase.validaFuncionario", errors.New("salário bruto deve ser maior que zero"), nil)
	}

	if f.DataAdmissao == "" {
		return customerror.NewError(customerror.ECONFLICT, "Data de admissão não pode ser vazio", "usecase.validaFuncionario", errors.New("data de admissão não pode ser vazio"), nil)
	}

	return nil
}

// convertData convert data from dd/mm/yyyy to yyyy-mm-dd
func convertData(dataBR string) (dataEN string) {
	data := strings.Split(dataBR, "/")
	dataEN = data[2] + "-" + data[1] + "-" + data[0]
	return dataEN
}

// Register validate funcionario and register it in repository
func (u *funcionariosUseCase) Register(ctx context.Context, funcionario entity.Funcionario) (*int64, *customerror.CustomError) {

	funcionario.Documento = cpf.Unmask(funcionario.Documento)

	err := validaFuncionario(funcionario)
	if err != nil {
		return nil, err
	}
	err = validaData(funcionario.DataAdmissao)
	if err != nil {
		return nil, err
	}

	employeeExist, err := u.r.GetByCPF(ctx, funcionario.Documento)
	if err != nil {
		return nil, err
	}

	if employeeExist != nil {
		return nil, customerror.NewError(customerror.ENOTFOUND, "Funcionário ja registrado", "usecase.GetByID", errors.New("funcionário já esta registrado"), nil)
	}

	funcionario.DataAdmissao = convertData(funcionario.DataAdmissao)

	return u.r.Register(ctx, funcionario)
}
