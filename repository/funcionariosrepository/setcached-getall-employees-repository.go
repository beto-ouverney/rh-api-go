package funcionariosrepository

import (
	"context"
	"encoding/json"
	"github.com/beto-ouverney/rh-api/customerror"
	"github.com/beto-ouverney/rh-api/db"
	"github.com/beto-ouverney/rh-api/entity"
	"os"
	"strconv"
	"time"
)

func (r *funcionariosRepository) SetCacheGetAllEmployees(funcionarios *[]entity.Funcionario) *customerror.CustomError {
	rdb := db.ConnectCacheDB()

	json, err := json.Marshal(funcionarios)
	if err != nil {
		return customerror.NewError(customerror.EINTERNAL, "Erro ao buscar employeescontroller", "funcionariosrepository.GetAll", err, nil)
	}

	number, err := strconv.Atoi(os.Getenv("CACHETIME"))
	if err != nil {
		return customerror.NewError(customerror.EINTERNAL, "Erro ao buscar employeescontroller", "funcionariosrepository.GetAll", err, nil)
	}

	exp := time.Duration(number) * time.Minute

	err = rdb.Set(context.Background(), cacheKeyGetAll, json, exp).Err()
	if err != nil {
		return customerror.NewError(customerror.EINTERNAL, "Erro ao buscar employeescontroller", "funcionariosrepository.GetAll", err, nil)
	}
	return nil
}
