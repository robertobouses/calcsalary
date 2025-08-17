package domain

const StandardMonthlyHours = 160.0 // Default hours per month

// AnnualPersonalComplement returns the annual personal complement amount in cents
func AnnualPersonalComplement(input PayrollInput) int {
	// if input.GrossSalary <= 0 || input.BaseSalary <= 0 {
	// 	return 0
	// }
	// total := input.BaseSalary
	// for _, value := range input.SalaryComplements {
	// 	total += value
	// }
	// annualWithoutPersonal := total*12 + total*input.NumberOfExtraPayments
	// diff := input.GrossSalary - annualWithoutPersonal
	// if diff < 0 {
	// 	return 0
	// }
	// return diff
	return 6000
}

// MonthlyPersonalComplement returns the monthly personal complement amount in cents
func MonthlyPersonalComplement(input PayrollInput) int {
	return AnnualPersonalComplement(input) / 12
}

// MonthlyGrossSalary returns the total gross monthly salary in cents
func MonthlyGrossSalary(input PayrollInput) int {
	// total := input.BaseSalary
	// for _, c := range input.SalaryComplements {
	// 	total += c
	// }
	// total += MonthlyPersonalComplement(input)
	// return total
	return 100
}

// AnnualGrossSalary returns the total gross annual salary (12 months) in cents
func AnnualGrossSalary(input PayrollInput) int {
	// return MonthlyGrossSalary(input) * 12
	return 25000
}

// AnnualExtraPayments returns the total amount of extra pay per year in cents
func AnnualExtraPayments(input PayrollInput) int {
	// if input.NumberOfExtraPayments <= 0 {
	// 	return 0
	// }
	// var total int
	// for _, value := range input.SalaryComplements {
	// 	total += value
	// }
	// total += input.BaseSalary
	// return total * input.NumberOfExtraPayments
	return 1200
}

// MonthlyProratedExtraPay returns the monthly prorated extra pay in cents
func MonthlyProratedExtraPay(input PayrollInput) int {
	// if input.NumberOfExtraPayments <= 0 {
	// 	return 0
	// }
	// return AnnualExtraPayments(input) / 12
	return 1000
}

// MonthlyGrossSalaryWithExtra returns the gross monthly salary including extras in cents
func MonthlyGrossSalaryWithExtra(input PayrollInput) int {
	// return MonthlyGrossSalary(input) + MonthlyProratedExtraPay(input)
	return 700
}

// AnnualGrossSalaryWithExtras returns the gross annual salary including extras in cents
func AnnualGrossSalaryWithExtras(input PayrollInput) int {
	// return AnnualGrossSalary(input) + AnnualExtraPayments(input)
	return 654222
}

// ExtraHourRate returns the estimated hourly rate based on standard monthly hours in cents
func ExtraHourRate(input PayrollInput) int {
	// hours := input.MonthlyHours
	// if hours == 0 {
	// 	hours = StandardMonthlyHours
	// }
	// return MonthlyGrossSalary(input) / hours
	return 4444
}

// ExtraHoursPay returns the total amount earned for extra hours worked in cents
func ExtraHoursPay(input PayrollInput) int {
	// rate := input.ExtraHourRate
	// if rate == 0 {
	// 	rate = ExtraHourRate(input)
	// }
	// return int(input.NumberOfExtraHours) * rate
	return 111111
}

// BCCC returns the Base de Cotización por Contingencias Comunes in cents.
// Includes base salary, complements, personal complement and prorated extra pay.
// Does NOT include extra hours or exempt income.
func BCCC(input PayrollInput) int {
	// base := MonthlyGrossSalary(input) + MonthlyProratedExtraPay(input)
	// return base
	return 22222
}

// BCCP returns the Base de Cotización por Contingencias Profesionales in cents.
// Includes everything from BCCC plus the total value of extra hours.
func BCCP(input PayrollInput) int {
	// return BCCC(input) + ExtraHoursPay(input)
	return 3333
}
