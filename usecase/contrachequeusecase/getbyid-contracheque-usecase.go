package contrachequeusecase

import (
	"context"
	"errors"
	"github.com/beto-ouverney/rh-api/customerror"
	"github.com/beto-ouverney/rh-api/entity"
	"github.com/beto-ouverney/rh-api/helpers/cpf"
	"math"
	"sync"
	"time"
)

var faixasINSS = [4]float64{1212.00, 2427.35, 3641.03, 7087.22}
var aliquotasINSS = [4]float64{0.075, 0.09, 0.12, 0.14}

var faixasIR = [4]float64{1903.98, 2826.65, 3751.05, 4664.68}
var aliquotasIR = [4]float64{0, 0.075, 0.15, 0.225}

const dependenteDescontoValor = 189.59

// roundFloat is a function to round float
func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

// calcINSSIR is a function to calculate INSS and IR values based on salary
func calcINSSIR(salario float64, faixas, aliquotas [4]float64) (valorDesconto float64) {
	var faixas2 []float64
	faixas2 = append(faixas2, faixas[0])

	for i := 1; i < len(faixas); i += 1 {
		if salario > faixas[i] {
			faixaTributavel := faixas[i] - (faixas[i-1] + 0.01)
			faixas2 = append(faixas2, faixaTributavel)
		} else {
			ultima := salario - faixas[i-1]
			faixas2 = append(faixas2, ultima)
			break
		}
	}
	for i, f := range faixas2 {
		valorDesconto += f * aliquotas[i]
	}
	return
}

// calcIR calculate IR
func calcIR(salarioTributavel float64, payslip *entity.Contracheque) {

	irValue := calcINSSIR(salarioTributavel, faixasIR, aliquotasIR)

	if salarioTributavel > faixasIR[3] {
		irValue += (salarioTributavel - faixasIR[3]) * 0.275
	}

	ir := entity.Lancamento{
		Descricao: "IR",
		Valor:     roundFloat(irValue, 2),
		Tipo:      "D",
	}
	payslip.Lancamentos = append(payslip.Lancamentos, ir)
}

// calcINSSIR calculate INSS and IR
func calcInssIR(payslip *entity.Contracheque, employee entity.Funcionario, wg *sync.WaitGroup) {
	defer wg.Done()
	inssValue := 0.00
	if employee.SalarioBruto > faixasINSS[0] {
		inssValue = calcINSSIR(employee.SalarioBruto, faixasINSS, aliquotasINSS)
	} else {
		inssValue = employee.SalarioBruto * aliquotasINSS[0]
	}

	inss := entity.Lancamento{
		Descricao: "INSS",
		Valor:     roundFloat(inssValue, 2),
		Tipo:      "D",
	}
	payslip.Lancamentos = append(payslip.Lancamentos, inss)

	descontoDependente := float64(employee.Dependente) * dependenteDescontoValor

	salarioTributavel := employee.SalarioBruto - inssValue - descontoDependente - employee.Pensao

	if salarioTributavel > faixasIR[0] {
		calcIR(salarioTributavel, payslip)
	}
}

// calcSalary calculate FGTS
func calcFGTS(payslip *entity.Contracheque, employee entity.Funcionario, wg *sync.WaitGroup) {
	defer wg.Done()

	fgts := entity.Lancamento{
		Descricao: "FGTS",
		Valor:     employee.SalarioBruto * 0.08,
		Tipo:      "D",
	}

	payslip.Lancamentos = append(payslip.Lancamentos, fgts)
}

// calcBenefits calculate benefits
func calcBenefits(payslip *entity.Contracheque, employee entity.Funcionario, wg *sync.WaitGroup) {
	defer wg.Done()

	if employee.Saude {
		l := entity.Lancamento{
			Descricao: "Plano de Saúde",
			Valor:     10,
			Tipo:      "D",
		}
		payslip.Lancamentos = append(payslip.Lancamentos, l)
	}
	if employee.Dental {
		l := entity.Lancamento{
			Descricao: "Plano Dental",
			Valor:     5,
			Tipo:      "D",
		}
		payslip.Lancamentos = append(payslip.Lancamentos, l)
	}

	if employee.Transporte && employee.SalarioBruto > 1500 {
		l := entity.Lancamento{
			Descricao: "Vale Transporte",
			Valor:     employee.SalarioBruto * 0.06,
			Tipo:      "D",
		}
		payslip.Lancamentos = append(payslip.Lancamentos, l)
	}
}

// GetByIDContrachequeUseCase generate payslip
func (u *contrachequeUseCase) GetByFuncionarioID(ctx context.Context, employeeID string) (*entity.Contracheque, *customerror.CustomError) {
	employee, err := u.r.GetByID(ctx, employeeID)
	if err != nil {
		return nil, err
	}
	if employee == nil {
		return nil, customerror.NewError(customerror.ENOTFOUND, "Funcionário não encontrado", "usecase.GetByID", errors.New("Funcionário não encontrado"), nil)
	}
	payslip := entity.Contracheque{}

	wg := sync.WaitGroup{}
	wg.Add(3)

	go calcInssIR(&payslip, *employee, &wg)
	go calcBenefits(&payslip, *employee, &wg)
	go calcFGTS(&payslip, *employee, &wg)

	wg.Wait()

	payslip.Nome = employee.Nome + " " + employee.Sobrenome
	payslip.Documento = cpf.Mask(employee.Documento)
	payslip.Setor = employee.Setor
	payslip.MesReferencia = time.Now().Format("01/2006")
	payslip.SalarioBruto = employee.SalarioBruto

	totalDescontos := 0.00
	for _, v := range payslip.Lancamentos {
		if v.Tipo == "D" {
			totalDescontos += v.Valor
		}
	}
	payslip.SalarioLiquido = payslip.SalarioBruto - totalDescontos

	return &payslip, nil

}
