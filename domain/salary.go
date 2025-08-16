package domain

const StandardMonthlyHours = 160.0 // Default hours per month

// MonthlyGrossSalary returns the total gross monthly salary in cents
// Uses input.PersonalComplement if > 0; otherwise calculates it.
func MonthlyGrossSalary(input PayrollInput) int {
	if input.BaseSalary <= 0 {
		return 0
	}
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

// AnnualGrossSalary returns the total gross annual salary (12 months) in cents
func AnnualGrossSalary(input PayrollInput) int {
	return MonthlyGrossSalary(input) * 12
}

// MonthlyExtraPay returns the monthly amount of extra pay in cents
func MonthlyExtraPay(input PayrollInput) int {
	return MonthlyGrossSalary(input)
}

// AnnualExtraPay returns the total amount of extra pay per year in cents
func AnnualExtraPay(input PayrollInput) int {
	return MonthlyExtraPay(input) * int(input.NumberOfExtraPayments)
}

// MonthlyProratedExtraPay returns the monthly prorated extra pay in cents
func MonthlyProratedExtraPay(input PayrollInput) int {
	return AnnualExtraPay(input) / 12
}

// MonthlyGrossSalaryWithExtras returns the gross monthly salary including extras in cents
func MonthlyGrossSalaryWithExtras(input PayrollInput) int {
	return MonthlyGrossSalary(input) + MonthlyExtraPay(input)
}

// AnnualGrossSalaryWithExtras returns the gross annual salary including extras in cents
func AnnualGrossSalaryWithExtras(input PayrollInput) int {
	return AnnualGrossSalary(input) + AnnualExtraPay(input)
}

// AnnualPersonalComplement returns the annual personal complement amount in cents
func AnnualPersonalComplement(input PayrollInput) int {
	if input.BaseSalary <= 0 || input.GrossSalary <= 0 {
		return 0
	}
	return input.GrossSalary - AnnualGrossSalaryWithExtras(input)
}

// MonthlyPersonalComplement returns the monthly personal complement amount in cents
func MonthlyPersonalComplement(input PayrollInput) int {
	if input.BaseSalary <= 0 || input.GrossSalary <= 0 {
		return 0
	}
	return AnnualPersonalComplement(input) / 12
}

// ExtraHourRate returns the estimated hourly rate based on standard monthly hours in cents
func ExtraHourRate(input PayrollInput) int {
	hours := input.MonthlyHours
	if hours == 0 {
		hours = StandardMonthlyHours
	}
	return MonthlyGrossSalary(input) / hours
}

// ExtraHoursPay returns the total amount earned for extra hours worked in cents
func ExtraHoursPay(input PayrollInput) int {
	if input.ExtraHourRate != 0 {
		return int(input.NumberOfExtraHours) * input.ExtraHourRate
	} else {
		extraHourRate := ExtraHourRate(input)
		return int(input.NumberOfExtraHours) * extraHourRate
	}
}

// BCCC returns the Base de Cotización por Contingencias Comunes in cents.
// Includes base salary, complements, personal complement and prorated extra pay.
// Does NOT include extra hours or exempt income.
func BCCC(input PayrollInput) int {
	base := MonthlyGrossSalary(input) + MonthlyProratedExtraPay(input)
	return base
}

// BCCP returns the Base de Cotización por Contingencias Profesionales in cents.
// Includes everything from BCCC plus the total value of extra hours.
func BCCP(input PayrollInput) int {
	base := MonthlyGrossSalary(input) + MonthlyProratedExtraPay(input)
	base += ExtraHoursPay(input)
	return base
}
