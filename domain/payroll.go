package domain

// PayrollOutput contains the detailed components of a payroll calculation.
type PayrollOutput struct {
	// Gross salary components
	BaseSalary             float64   // Monthly base salary
	PersonalComplement     float64   // Personal complement or allowance
	SalaryComplements      []float64 // Other salary complements
	ExtraHoursPay          float64   // Total pay for extra hours worked
	MonthlyGrossWithExtras float64   // Total gross monthly salary including extras and extra hours

	// Social Security contribution bases
	BaseBCCC float64 // Contribution base for common contingencies
	BaseBCCP float64 // Contribution base for professional contingencies

	// Income tax (IRPF)
	IrpfAmount        float64 // Monthly IRPF withholding amount
	IrpfEffectiveRate float64 // Effective IRPF tax rate

	// Social Security contributions breakdown (worker and employer)
	SSContributions SSCotisations

	// Net salary after deductions
	NetSalary float64 // Net monthly salary payable to the employee
}

// GeneratePayrollOutput calculates and returns all payroll parts based on the input data.
func GeneratePayrollOutput(input PayrollInput) PayrollOutput {
	ss := CalculateSSCotisations(input)
	gross := MonthlyGrossSalary(input)
	extras := MonthlyProratedExtraPay(input)
	extraHours := ExtraHoursPay(input)

	bccc := BCCC(input)
	bccp := BCCP(input)

	irpfAnnual, irpfRate := AnnualIrpf(input)
	irpfMonthly := irpfAnnual / 12

	net := gross + extras + extraHours - irpfMonthly - ss.TotalWorker

	return PayrollOutput{
		BaseSalary:             input.BaseSalary,
		PersonalComplement:     input.PersonalComplement,
		SalaryComplements:      input.SalaryComplements,
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
