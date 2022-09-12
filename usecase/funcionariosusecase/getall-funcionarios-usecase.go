package funcionariosusecase

import (
	"context"
	"encoding/json"
	"github.com/beto-ouverney/rh-api/customerror"
	"github.com/beto-ouverney/rh-api/entity"
)

func (u *funcionariosUseCase) GetAll(ctx context.Context) (*[]entity.Funcionario, *customerror.CustomError) {

	cache, err := u.r.GetAllCache()

	if err != nil {
		return nil, err
	}
	if cache != nil {
		employees := &[]entity.Funcionario{}
		errJ := json.Unmarshal([]byte(*cache), &employees)
		if errJ != nil {
			return nil, customerror.NewError(customerror.EINTERNAL, "Erro ao buscar employeescontroller", "employeescontroller-usecase.GetAll", errJ, nil)
		}
		return employees, nil
	}

	funcionarios, err := u.r.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	err = u.r.SetCacheGetAllEmployees(funcionarios)
	if err != nil {
		return nil, err
	}

	return funcionarios, nil
}
