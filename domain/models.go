package domain

type PayrollInput struct {
	BaseSalary            float64
	SalaryComplements     []float64
	PersonalComplement    float64
	GrossSalary           float64
	NumberOfExtraPayments int
}
