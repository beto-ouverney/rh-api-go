package employeescontroller

import (
	"context"
	"encoding/json"
	"github.com/beto-ouverney/rh-api/customerror"
)

// GetByID returns a funcionario by id from service and return a json to handler
func (c *funcionariosController) GetByID(ctx context.Context, id string) ([]byte, *customerror.CustomError) {
	user, err := c.u.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	userJson, errJ := json.MarshalIndent(user, "", "  ")
	if errJ != nil {
		return nil, customerror.NewError(customerror.EINTERNAL, "Erro ao converter para json", "controller.GetByID", errJ, nil)
	}

	return userJson, nil
}
