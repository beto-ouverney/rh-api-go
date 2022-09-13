package contrachequecontroller

import (
	"github.com/beto-ouverney/rh-api/customerror"
	"github.com/beto-ouverney/rh-api/usecase/contrachequeusecase"
)

// IContrachequeController presents the interface for the contracheque controller
type IContrachequeController interface {
	GetByFuncionarioID(ctx, funcionaarioID string) ([]byte, *customerror.CustomError)
}

type contrachequeController struct {
	u contrachequeusecase.IContrachequeUseCase
}

// New creates a new contracheque controller
func New() *contrachequeController {
	return &contrachequeController{
		u: contrachequeusecase.New(),
	}
}
