package entity

//Lancamentos presents the debits and credits of the payslip
type Lancamento struct {
	Tipo      string  `json:"tipo"`
	Valor     float64 `json:"valor"`
	Descricao string  `json:"descricao"`
}
