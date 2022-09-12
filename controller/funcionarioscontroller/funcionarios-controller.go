package employeescontroller

import (
	"context"
	"github.com/beto-ouverney/rh-api/customerror"
	"github.com/beto-ouverney/rh-api/entity"
	"github.com/beto-ouverney/rh-api/usecase/funcionariosusecase"
)

type IFuncionariosController interface {
	GetByID(ctx context.Context, id string) ([]byte, *customerror.CustomError)
	GetAll(ctx context.Context) ([]byte, *customerror.CustomError)
	Register(ctx context.Context, funcionario entity.Funcionario) ([]byte, *customerror.CustomError)
}

type funcionariosController struct {
	u funcionariosusecase.IFuncionariosUseCase
}

func New() *funcionariosController {
	return &funcionariosController{
		funcionariosusecase.New(),
	}
}
