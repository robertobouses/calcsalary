package domain

func GrossSalaryMensual(payrollInput PayrollInput) float64 {
	total := payrollInput.BaseSalary + payrollInput.PersonalComplement

	for _, complement := range payrollInput.SalaryComplements {
		total += complement
	}

	return total
}

func ExtraPay(payrollInput PayrollInput) float64 {
	return GrossSalaryMensual(payrollInput)
}

func ProratedExtraPay(payrollInput PayrollInput) float64 {
	return ExtraPay(payrollInput) / 12
}
