package funcionariosrepository

import (
	"context"
	"github.com/beto-ouverney/rh-api/customerror"
	"github.com/beto-ouverney/rh-api/entity"
)

func (r *funcionariosRepository) Register(ctx context.Context, funcionario entity.Funcionario) (*int64, *customerror.CustomError) {

	stmt, err := r.sqlx.PrepareNamedContext(ctx, "INSERT INTO Funcionarios (nome, sobrenome, documento, setor, salario_bruto, data_admissao, saude, dental, transporte, dependente, pensao) "+
		"VALUES (:nome, :sobrenome, :documento, :setor, :salario_bruto, :data_admissao, :saude, :dental, :transporte, :dependente, :pensao) RETURNING id")
	if err != nil {
		return nil, customerror.NewError(customerror.EINTERNAL, "Internal error", "funcionariosrepository.Add", err, nil)
	}
	var id int64
	err = stmt.Get(&id, funcionario)

	if err != nil {
		return nil, customerror.NewError(customerror.EINTERNAL, "Internal error", "funcionariosrepository.Add", err, nil)
	}

	return &id, nil

}
