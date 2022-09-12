package funcionariosrepository

import (
	"context"
	"github.com/beto-ouverney/rh-api/customerror"
	"github.com/beto-ouverney/rh-api/entity"
)

// GetAll gets all funcionarios in database
func (r *funcionariosRepository) GetAll(ctx context.Context) (*[]entity.Funcionario, *customerror.CustomError) {
	var funcionarios []entity.Funcionario
	err := r.sqlx.SelectContext(ctx, &funcionarios, "SELECT * FROM Funcionarios")
	if err != nil {
		return nil, customerror.NewError(customerror.EINTERNAL, "Erro ao buscar funcionarios", "funcionariosrepository.GetAll", err, nil)
	}

	return &funcionarios, nil
}
