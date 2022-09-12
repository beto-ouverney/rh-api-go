package funcionariosusecase

import (
	"context"
	"errors"
	"github.com/beto-ouverney/rh-api/customerror"
	"github.com/beto-ouverney/rh-api/entity"
)

// GetByID check if funcionario exists and return it
func (u *funcionariosUseCase) GetByID(ctx context.Context, id string) (*entity.Funcionario, *customerror.CustomError) {

	user, err := u.r.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, customerror.NewError(customerror.ENOTFOUND, "Funcionário não encontrado", "usecase.GetByID", errors.New("funcionário não encontrado"), nil)
	}

	return user, nil
}
