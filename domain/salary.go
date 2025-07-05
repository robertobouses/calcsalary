package domain

// MonthlyGrossSalary returns the total gross monthly salary
func MonthlyGrossSalary(input PayrollInput) float64 {
	total := input.BaseSalary + input.PersonalComplement
	for _, c := range input.SalaryComplements {
		total += c
	}
	return total
}

// AnnualGrossSalary returns the total gross annual salary (12 months)
func AnnualGrossSalary(input PayrollInput) float64 {
	return MonthlyGrossSalary(input) * 12
}

// MonthlyExtraPay returns the monthly amount of extra pay
func MonthlyExtraPay(input PayrollInput) float64 {
	return MonthlyGrossSalary(input)
}

// AnnualExtraPay returns the total amount of extra pay per year
func AnnualExtraPay(input PayrollInput) float64 {
	return MonthlyExtraPay(input) * float64(input.NumberOfExtraPayments)
}

// MonthlyProratedExtraPay returns the monthly prorated extra pay
func MonthlyProratedExtraPay(input PayrollInput) float64 {
	return AnnualExtraPay(input) / 12
}

// MonthlyGrossSalaryWithExtras returns the gross monthly salary including extras
func MonthlyGrossSalaryWithExtras(input PayrollInput) float64 {
	return MonthlyGrossSalary(input) + MonthlyExtraPay(input)
}

// AnnualGrossSalaryWithExtras returns the gross annual salary including extras
func AnnualGrossSalaryWithExtras(input PayrollInput) float64 {
	return AnnualGrossSalary(input) + AnnualExtraPay(input)
}

// AnnualPersonalComplement returns the annual personal complement amount
func AnnualPersonalComplement(input PayrollInput) float64 {
	return input.GrossSalary - AnnualGrossSalaryWithExtras(input)
}

// MonthlyPersonalComplement returns the monthly personal complement amount
func MonthlyPersonalComplement(input PayrollInput) float64 {
	return AnnualPersonalComplement(input) / 12
}
