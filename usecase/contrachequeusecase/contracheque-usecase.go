package contrachequeusecase

import (
	"context"
	"github.com/beto-ouverney/rh-api/customerror"
	"github.com/beto-ouverney/rh-api/entity"
	"github.com/beto-ouverney/rh-api/repository/funcionariosrepository"
)

type IContrachequeUseCase interface {
	GetByFuncionarioID(ctx context.Context, employeeID string) (*entity.Contracheque, *customerror.CustomError)
}

type contrachequeUseCase struct {
	r funcionariosrepository.IFuncionariosRepository
}

func New() *contrachequeUseCase {
	return &contrachequeUseCase{r: funcionariosrepository.New()}
}
