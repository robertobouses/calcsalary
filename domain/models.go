package domain

// PayrollInput represents the input data required to calculate payroll values.
type PayrollInput struct {
	BaseSalary            float64   // Monthly base salary.
	SalaryComplements     []float64 // Additional taxable complements.
	PersonalComplement    float64   // Extra personal bonus or allowance.
	GrossSalary           float64   // Optional total gross salary (can be used for adjustments).
	NumberOfExtraPayments int       // Number of extra payments per year (e.g. 2 for summer and Christmas).
	NumberOfChildren      int       // Number of children

	// Disability related fields
	HasDisability       bool // If the taxpayer has a disability
	HasSevereDisability bool // Disability degree ≥ 65%
	NeedsAssistance     bool // Needs help from third party

	// Ascendants
	HasAscendantsOver65   bool // Ascendants over 65 years old
	HasDisabledAscendants bool // Disabled ascendants
}
