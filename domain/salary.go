package domain

func MonthlyGrossSalary(input PayrollInput) float64 {
	total := input.BaseSalary + input.PersonalComplement
	for _, c := range input.SalaryComplements {
		total += c
	}
	return total
}

func AnnualGrossSalary(input PayrollInput) float64 {
	return MonthlyGrossSalary(input) * 12
}

func MonthlyExtraPay(input PayrollInput) float64 {
	return MonthlyGrossSalary(input)
}

func AnnualExtraPay(input PayrollInput) float64 {
	return MonthlyExtraPay(input) * float64(input.NumberOfExtraPayments)
}

func AnnualGrossSalaryWithExtras(input PayrollInput) float64 {
	return AnnualGrossSalary(input) + AnnualExtraPay(input)
}

func MonthlyProratedExtraPay(input PayrollInput) float64 {
	return AnnualExtraPay(input) / 12
}

func AnnualPersonalComplement(input PayrollInput) float64 {
	return input.GrossSalary - AnnualGrossSalaryWithExtras(input)
}

func MonthlyPersonalComplement(input PayrollInput) float64 {
	return AnnualPersonalComplement(input) / 12
}
