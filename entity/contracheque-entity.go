package entity

// Contracheque presents the entity of the payslip
type Contracheque struct {
	MesReferencia  string       `json:"mes_referencia"`
	FuncionarioID  int64        `json:"funcionario_id"`
	Nome           string       `json:"nome"`
	Documento      string       `json:"documento"`
	Setor          string       `json:"setor"`
	SalarioBruto   float64      `json:"salario_bruto"`
	Lancamentos    []Lancamento `json:"lancamentos"`
	SalarioLiquido float64      `json:"salario_liquido"`
}
