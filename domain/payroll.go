package domain

// PayrollOutput contains the detailed components of a payroll calculation.
// All monetary values are expressed in cents (int).
type PayrollOutput struct {
	// Gross salary components (in cents)
	BaseSalary             int   // Monthly base salary
	SalaryComplements      []int // Other salary complements
	PersonalComplement     int   // Personal complement or allowance
	ExtraHoursPay          int   // Total pay for extra hours worked
	MonthlyGrossWithExtras int   // Total gross monthly salary including extras and extra hours

	// Social Security contribution bases (in cents)
	BaseBCCC int // Contribution base for common contingencies
	BaseBCCP int // Contribution base for professional contingencies

	// Income tax (IRPF)
	IrpfAmount        int // Monthly IRPF withholding amount in cents
	IrpfEffectiveRate int // Effective IRPF tax rate (basis points, per 10,000)

	// Social Security contributions breakdown (worker and employer)
	SSContributions SSCotisations

	// Net salary after deductions (in cents)
	NetSalary int
}

// GeneratePayrollOutput calculates and returns all payroll parts based on the input data.
func GeneratePayrollOutput(input PayrollInput) PayrollOutput {
	ss := CalculateSSCotisations(input)
	gross := MonthlyGrossSalary(input)       // in cents
	extras := MonthlyProratedExtraPay(input) // in cents
	extraHours := ExtraHoursPay(input)       // in cents

	bccc := BCCC(input) // in cents
	bccp := BCCP(input) // in cents

	irpfAnnual, irpfRate := AnnualIrpf(input) // irpfAnnual in cents, rate in basis points
	irpfMonthly := irpfAnnual / 12            // monthly IRPF in cents

	net := gross + extras + extraHours - irpfMonthly - ss.TotalWorker

	return PayrollOutput{
		BaseSalary:             input.BaseSalary,
		SalaryComplements:      input.SalaryComplements,
		PersonalComplement:     MonthlyPersonalComplement(input),
		ExtraHoursPay:          extraHours,
		MonthlyGrossWithExtras: gross + extras + extraHours,

		BaseBCCC: bccc,
		BaseBCCP: bccp,

		IrpfAmount:        irpfMonthly,
		IrpfEffectiveRate: irpfRate,

		SSContributions: ss,
		NetSalary:       net,
	}
}
