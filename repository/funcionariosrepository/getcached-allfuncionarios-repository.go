package funcionariosrepository

import (
	"context"
	"github.com/beto-ouverney/rh-api/customerror"
	"github.com/beto-ouverney/rh-api/db"
)

// GetAllCache gets all funcionarios in cache
func (r *funcionariosRepository) GetAllCache() (*string, *customerror.CustomError) {
	rdb := db.ConnectCacheDB()

	cached, err := rdb.Get(context.Background(), cacheKeyGetAll).Result()
	if err != nil {
		if err.Error() == "redis: nil" {
			return nil, nil
		}
		return nil, customerror.NewError(customerror.EINTERNAL, "Erro ao buscar employeescontroller", "funcionariosrepository.GetAll", err, nil)
	}
	if cached == "null" {
		return nil, nil
	}
	return &cached, nil

}
