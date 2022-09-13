package funcionarioscontroller

import (
	"context"
	"encoding/json"
	"github.com/beto-ouverney/rh-api/customerror"
	"github.com/beto-ouverney/rh-api/entity"
)

// Register send a funcionario to service and return a json to handler with the id of the funcionario
func (c *funcionariosController) Register(ctx context.Context, funcionario entity.Funcionario) ([]byte, *customerror.CustomError) {

	id, err := c.u.Register(ctx, funcionario)
	if err != nil {
		return nil, err
	}

	var newUserID struct {
		ID int64 `json:"id"`
	}

	newUserID.ID = *id

	idJson, errJ := json.MarshalIndent(newUserID, "", "  ")
	if errJ != nil {
		return nil, customerror.NewError(customerror.EINTERNAL, "Erro ao converter para json", "controller.Register", errJ, nil)
	}

	return idJson, nil
}
