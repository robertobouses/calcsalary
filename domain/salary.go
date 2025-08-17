package domain

const StandardMonthlyHours = 160.0 // Default hours per month

// MonthlyGrossSalary returns the total gross monthly salary in cents
// Uses input.PersonalComplement if > 0; otherwise calculates it.
func MonthlyGrossSalary(input PayrollInput) int {
	if input.BaseSalary <= 0 {
		return 0
	}

	personalComplement := input.PersonalComplement
	if personalComplement == 0 && input.GrossSalary > 0 {
		annualBase := input.BaseSalary * 12
		for _, c := range input.SalaryComplements {
			annualBase += c * 12
		}
		diff := input.GrossSalary - annualBase
		if diff > 0 {
			personalComplement = diff / 12
		}
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

// AnnualExtraPayments returns the total amount of extra pay per year in cents
func AnnualExtraPayments(input PayrollInput) int {
	if input.NumberOfExtraPayments <= 0 {
		return 0
	}
	monthlyExtra := 0
	if input.PersonalComplement == 0 && input.GrossSalary > 0 {
		annualWithExtras := input.BaseSalary * input.NumberOfExtraPayments
		for _, c := range input.SalaryComplements {
			annualWithExtras += c * 12
		}
		diff := input.GrossSalary - annualWithExtras
		if diff > 0 {
			monthlyExtra = diff / 12
		}
	}
	total := monthlyExtra
	for _, c := range input.SalaryComplements {
		total += c
	}
	return total * int(input.NumberOfExtraPayments)
}

// MonthlyProratedExtraPay returns the monthly prorated extra pay in cents
func MonthlyProratedExtraPay(input PayrollInput) int {
	if input.NumberOfExtraPayments <= 0 {
		return 0
	}
	return AnnualExtraPayments(input) / 12
}

// MonthlyGrossSalaryWithExtra returns the gross monthly salary including extras in cents
func MonthlyGrossSalaryWithExtra(input PayrollInput) int {
	return MonthlyGrossSalary(input) + MonthlyProratedExtraPay(input)
}

// AnnualGrossSalaryWithExtras returns the gross annual salary including extras in cents
func AnnualGrossSalaryWithExtras(input PayrollInput) int {
	return AnnualGrossSalary(input) + AnnualExtraPayments(input)
}

// AnnualPersonalComplement returns the annual personal complement amount in cents
func AnnualPersonalComplement(input PayrollInput) int {
	if input.PersonalComplement > 0 {
		return input.PersonalComplement * 12
	}
	if input.GrossSalary <= 0 || input.BaseSalary <= 0 {
		return 0
	}

	knownAnnual := input.BaseSalary*12 + len(input.SalaryComplements)*12
	for _, c := range input.SalaryComplements {
		knownAnnual += c * 12
	}
	diff := input.GrossSalary - knownAnnual
	if diff < 0 {
		return 0
	}
	return diff
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
	rate := input.ExtraHourRate
	if rate == 0 {
		rate = ExtraHourRate(input)
	}
	return int(input.NumberOfExtraHours) * rate
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
	return BCCC(input) + ExtraHoursPay(input)
}
