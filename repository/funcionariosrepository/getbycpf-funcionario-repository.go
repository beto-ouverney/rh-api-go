package funcionariosrepository

import (
	"context"
	"github.com/beto-ouverney/rh-api/customerror"
	"github.com/beto-ouverney/rh-api/entity"
)

// GetByCPF gets a funcionario by cpf in database
func (r *funcionariosRepository) GetByCPF(ctx context.Context, cpf string) (*entity.Funcionario, *customerror.CustomError) {
	var funcionario entity.Funcionario

	err := r.sqlx.GetContext(ctx, &funcionario, "SELECT * FROM Funcionarios WHERE documento = $1", cpf)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, nil
		}
		return nil, customerror.NewError(customerror.ENOTFOUND, "Funcionario não encontrado", "funcionariosrepository.GetByCPF", err, nil)
	}

	return &funcionario, nil
}
