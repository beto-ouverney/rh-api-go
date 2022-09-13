package test

import (
	"github.com/beto-ouverney/rh-api/entity"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"testing"
)

var allFuncMock = []entity.Funcionario{
	{
		ID:           1,
		Nome:         "Alberto",
		Sobrenome:    "Ouverney Paz",
		Documento:    "58303831798",
		Setor:        "TI",
		SalarioBruto: 7000,
		DataAdmissao: "2020-01-01T00:00:00Z",
		Saude:        false,
		Dental:       false,
		Transporte:   false,
		Dependente:   0,
		Pensao:       0,
	},
	{
		ID:           2,
		Nome:         "Amanda",
		Sobrenome:    "Oliveira",
		Documento:    "61650548702",
		Setor:        "RH",
		SalarioBruto: 9000,
		DataAdmissao: "2020-01-01T00:00:00Z",
		Saude:        false,
		Dental:       false,
		Transporte:   false,
		Dependente:   0,
		Pensao:       0,
	},
}

var contrachequeMock = entity.Contracheque{
	MesReferencia: "01/2021",
	FuncionarioID: 2,
	Nome:          "Alberto Ouverney Paz",
	Documento:     "583.038.317-98",
	Setor:         "TI",
	SalarioBruto:  7000.00,
	Lancamentos: []entity.Lancamento{
		{
			Tipo:      "D",
			Valor:     560,
			Descricao: "FGTS",
		},
		{
			Tipo:      "D",
			Valor:     816.18,
			Descricao: "INSS",
		},
		{
			Tipo:      "D",
			Valor:     831.19,
			Descricao: "IR",
		},
	},
	SalarioLiquido: 4792.63,
}

var schemasInit = [3]string{
	`CREATE TABLE IF NOT EXISTS Funcionarios(
	id SERIAL PRIMARY KEY,
	nome VARCHAR(20),
	sobrenome VARCHAR(60),
	documento VARCHAR(11),
	setor VARCHAR(20),
	salario_bruto DECIMAL(10,2),
	data_admissao DATE,
	saude BOOLEAN,
	dental BOOLEAN,
	transporte BOOLEAN,
	dependente INT,
	pensao DECIMAL(10,2)
);`,

	`INSERT INTO Funcionarios (nome, sobrenome, documento, setor, salario_bruto,data_admissao, saude, dental, transporte, dependente, pensao) VALUES ('Alberto', 'Ouverney Paz', '58303831798', 'TI', 7000.00, '2020-01-01', false, false, false, 0, 0.00);`,

	`INSERT INTO Funcionarios (nome, sobrenome, documento, setor, salario_bruto,data_admissao, saude, dental, transporte, dependente, pensao) VALUES ('Amanda', 'Oliveira', '61650548702', 'RH', 9000.00, '2020-01-01', false, false, false, 0, 0.00);`,
}

const POSTGREES_CONNECTION = "user=root password=password dbname=rh_db_test sslmode=disable"

func dropDBTEST(t *testing.T) {

	conn, err := sqlx.Open("postgres", POSTGREES_CONNECTION)
	if err != nil {
		panic(err.Error())
	}
	_, err = conn.Exec("DROP TABLE IF EXISTS Funcionarios")
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Database test dropped successfully")

	createSchema := `CREATE TABLE IF NOT EXISTS Funcionarios(
	id SERIAL PRIMARY KEY,
	nome VARCHAR(20),
	sobrenome VARCHAR(60),
	documento VARCHAR(11),
	setor VARCHAR(20),
	salario_bruto DECIMAL(10,2),
	data_admissao DATE,
	saude BOOLEAN,
	dental BOOLEAN,
	transporte BOOLEAN,
	dependente INT,
	pensao DECIMAL(10,2)
);`
	_, err = conn.Exec(createSchema)
	if err != nil {
		t.Fatal(err)
	}

}

func initDBTest(t *testing.T) {
	t.Log("Initializing database test")

	conn, err := sqlx.Open("postgres", POSTGREES_CONNECTION)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = conn.Exec("DROP TABLE IF EXISTS Funcionarios")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Database test dropped successfully")

	t.Setenv("POSTGRES_DB", "rh_db_test")

	for _, s := range schemasInit {
		_, err = conn.Exec(s)
		if err != nil {
			t.Fatal(err)
		}
	}
	defer conn.Close()
	t.Log("Database test created successfully")
}
