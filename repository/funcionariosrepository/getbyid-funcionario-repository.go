package funcionariosrepository

import (
	"context"
	"github.com/beto-ouverney/rh-api/customerror"
	"github.com/beto-ouverney/rh-api/entity"
)

func (r *funcionariosRepository) GetByID(ctx context.Context, id string) (*entity.Funcionario, *customerror.CustomError) {
	var funcionario entity.Funcionario

	err := r.sqlx.GetContext(ctx, &funcionario, "SELECT * FROM Funcionarios WHERE id = $1", id)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, nil
		}
		return nil, customerror.NewError(customerror.EINTERNAL, "Internal error", "funcionariosrepository.GetByID", err, nil)
	}
	return &funcionario, nil
}
