package funcionariosrepository

import (
	"context"
	"github.com/beto-ouverney/rh-api/customerror"
	"github.com/beto-ouverney/rh-api/db"
	"github.com/beto-ouverney/rh-api/entity"
	"github.com/jmoiron/sqlx"
)

// IFuncionariosRepository presents the interface for the funcionarios repository
type IFuncionariosRepository interface {
	Register(ctx context.Context, funcionario entity.Funcionario) (*int64, *customerror.CustomError)
	GetByID(ctx context.Context, id string) (*entity.Funcionario, *customerror.CustomError)
	GetAll(ctx context.Context) (*[]entity.Funcionario, *customerror.CustomError)
	GetByCPF(ctx context.Context, cpf string) (*entity.Funcionario, *customerror.CustomError)
	GetAllCache() (*string, *customerror.CustomError)
	SetCacheGetAllEmployees(funcionarios *[]entity.Funcionario) *customerror.CustomError
}

const cacheKeyGetAll = "employees-get-all"

type funcionariosRepository struct {
	sqlx *sqlx.DB
}

// New creates a new funcionarios repository
func New() *funcionariosRepository {
	return &funcionariosRepository{
		db.ConnectDB(),
	}
}
