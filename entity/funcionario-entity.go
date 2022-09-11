package entity

//Funcionario is the entity that represents the employee
type Funcionario struct {
	ID           int64   `json:"ïd" db:"id"`
	Nome         string  `json:"nome" db:"nome"`
	Sobrenome    string  `json:"sobrenome" db:"sobrenome"`
	Documento    string  `json:"documento" db:"documento"`
	Setor        string  `json:"setor" db:"setor"`
	SalarioBruto float64 `json:"salario_bruto" db:"salario_bruto"`
	DataAdmissao string  `json:"data_admissao" db:"data_admissao"`
	Saude        bool    `json:"saude" db:"saude"`
	Dental       bool    `json:"dental" db:"dental"`
	Transporte   bool    `json:"transporte" db:"transporte"`
	Dependente   int     `json:"dependentes" db:"dependente"`
	Pensao       float64 `json:"pensao_alimenticia" db:"pensao"`
}
