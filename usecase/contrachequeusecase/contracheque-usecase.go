package contrachequeusecase

import (
	"context"
	"github.com/beto-ouverney/rh-api/customerror"
	"github.com/beto-ouverney/rh-api/entity"
	"github.com/beto-ouverney/rh-api/repository/funcionariosrepository"
)

// IContrachequeUseCase presents the interface for the contracheque use case
type IContrachequeUseCase interface {
	GetByFuncionarioID(ctx context.Context, employeeID string) (*entity.Contracheque, *customerror.CustomError)
}

type contrachequeUseCase struct {
	r funcionariosrepository.IFuncionariosRepository
}

// New creates a new contracheque use case
func New() *contrachequeUseCase {
	return &contrachequeUseCase{r: funcionariosrepository.New()}
}
