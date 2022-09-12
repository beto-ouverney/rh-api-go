package funcionariosusecase

import (
	"context"
	"github.com/beto-ouverney/rh-api/customerror"
	"github.com/beto-ouverney/rh-api/entity"
	"github.com/beto-ouverney/rh-api/repository/funcionariosrepository"
)

// IFuncionariosUseCase presents the interface for the funcionarios use case
type IFuncionariosUseCase interface {
	GetByID(ctx context.Context, id string) (*entity.Funcionario, *customerror.CustomError)
	GetAll(ctx context.Context) (*[]entity.Funcionario, *customerror.CustomError)
	Register(ctx context.Context, funcionario entity.Funcionario) (*int64, *customerror.CustomError)
}

type funcionariosUseCase struct {
	r funcionariosrepository.IFuncionariosRepository
}

// New creates a new funcionarios use case
func New() *funcionariosUseCase {
	return &funcionariosUseCase{
		funcionariosrepository.New(),
	}
}
