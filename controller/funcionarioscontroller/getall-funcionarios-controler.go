package funcionarioscontroller

import (
	"context"
	"encoding/json"
	"github.com/beto-ouverney/rh-api/customerror"
)

// GetAll returns all funcionarios from service and return a json to handler
func (c *funcionariosController) GetAll(ctx context.Context) ([]byte, *customerror.CustomError) {
	funcionarios, err := c.u.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	jsonF, errJ := json.MarshalIndent(funcionarios, "", "  ")
	if errJ != nil {
		return nil, customerror.NewError(customerror.EINTERNAL, "Erro ao buscar employeescontroller", "funcionarioscontroller.GetAll", errJ, nil)
	}
	return jsonF, nil
}
