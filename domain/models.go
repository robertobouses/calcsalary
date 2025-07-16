package domain

// PayrollInput represents the input data required to calculate payroll values.
// All monetary values are expressed in cents (int).
type PayrollInput struct {
	BaseSalary            int   // Monthly base salary in cents
	SalaryComplements     []int // Additional taxable complements in cents
	PersonalComplement    int   // Extra personal bonus or allowance in cents
	GrossSalary           int   // Optional total gross salary (in cents)
	NumberOfExtraPayments int   // Number of extra payments per year (e.g. 2 for summer and Christmas)

	NumberOfExtraHours int // Total number of extra hours worked in the month
	ExtraHourRate      int // Pay per extra hour in cents

	MonthlyHours int // Default to 160 if 0

	NumberOfChildren int // Number of children

	// Disability related fields
	HasDisability       bool // If the taxpayer has a disability
	HasSevereDisability bool // Disability degree ≥ 65%
	NeedsAssistance     bool // Needs help from third party

	// Ascendants
	HasAscendantsOver65   bool // Ascendants over 65 years old
	HasDisabledAscendants bool // Disabled ascendants
}
