package contrachequecontroller

import (
	"context"
	"encoding/json"
	"github.com/beto-ouverney/rh-api/customerror"
)

// GetByID returns a contracheque from service and return a json to handler
func (c *contrachequeController) GetByFuncionarioID(ctx context.Context, id string) ([]byte, *customerror.CustomError) {
	contracheque, err := c.u.GetByFuncionarioID(ctx, id)
	if err != nil {
		return nil, err
	}
	jsonCC, errJ := json.MarshalIndent(contracheque, "", "  ")
	if errJ != nil {
		return nil, customerror.NewError(customerror.EINTERNAL, "Erro ao buscar contracheque", "contrachequecontroller.GetByID", errJ, nil)
	}
	return jsonCC, nil
}
