package domain

const StandardMonthlyHours = 160.0

// MonthlyGrossSalary returns the total gross monthly salary.
// Uses input.PersonalComplement if > 0; otherwise calculates it.
func MonthlyGrossSalary(input PayrollInput) float64 {
	personalComplement := input.PersonalComplement
	if personalComplement == 0 {
		personalComplement = MonthlyPersonalComplement(input)
	}
	total := input.BaseSalary + personalComplement
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

// ExtraHourRate returns the estimated hourly rate based on standard monthly hours.
func ExtraHourRate(input PayrollInput) float64 {
	hours := input.MonthlyHours
	if hours == 0 {
		hours = StandardMonthlyHours
	}
	return MonthlyGrossSalary(input) / hours
}

// ExtraHoursPay returns the total amount earned for extra hours worked
func ExtraHoursPay(input PayrollInput) float64 {
	if input.ExtraHourRate != 0 {
		return float64(input.NumberOfExtraHours) * input.ExtraHourRate
	} else {
		extraHourRate := ExtraHourRate(input)
		return float64(input.NumberOfExtraHours) * extraHourRate
	}
}

// BCCC returns the Base de Cotización por Contingencias Comunes.
// Includes base salary, complements, personal complement and prorated extra pay.
// Does NOT include extra hours or exempt income.
func BCCC(input PayrollInput) float64 {
	base := MonthlyGrossSalary(input) + MonthlyProratedExtraPay(input)
	return base
}

// BCCP returns the Base de Cotización por Contingencias Profesionales.
// Includes everything from BCCC plus the total value of extra hours.
func BCCP(input PayrollInput) float64 {
	base := MonthlyGrossSalary(input) + MonthlyProratedExtraPay(input)
	base += ExtraHoursPay(input)
	return base
}
