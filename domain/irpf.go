package domain

const (
	ReductionMax            = 556500  // 5565.0 Maximum reduction for work income
	ReductionMin            = 200000  // 2000.0 Minimum fixed reduction
	ReductionThresholdStart = 1404750 // 14047.5 Start of progressive reduction
	ReductionThresholdEnd   = 1770750 // 17707.5 End of progressive reduction
	ReductionSlope          = 34      // 0.34 Slope for progressive reduction calculation
)

const (
	// Personal minimum 5550.0
	PersonalMin = 555000

	// Minimums per dependent (children)
	FirstChildMin  = 240000 // 2400.0
	SecondChildMin = 270000 // 2700.0
	ThirdChildMin  = 400000 // 4000.0
	FourthChildMin = 450000 // 4500.0

	// Disability minimums for the taxpayer
	DisabilityMin_Light  = 300000 // 3000.0 Disability degree < 65%
	DisabilityMin_Severe = 900000 // 9000.0 Disability degree ≥ 65%
	DisabilityAssistance = 300000 // 3000.0 If assistance from a third party is required

	// Minimums for ascendants over 65 years old or disabled
	AscendantMin_65       = 115000 // 1150.0 // Over 65 years old
	AscendantMin_Disabled = 255000 // 2550.0 // Disabled
)

// IRPF tax brackets and rates (2024–2025) – full (state + regional)
var TaxBrackets = []struct {
	Limit int // Upper limit of the bracket
	Rate  int // Full tax rate (state + CCAA)
}{
	{12450.0, 1900},  // 19% up to 12,450 €
	{20200.0, 2400},  // 24% up to 20,200 €
	{35200.0, 3000},  // 30% up to 35,200 €
	{60000.0, 3700},  // 37% up to 60,000 €
	{300000.0, 4500}, // 45% up to 300,000 €
	{1e12, 4700},     // 47% from 300,000 € upwards
}

// WorkIncomeReduction returns the reduction applied to gross income from work.
func WorkIncomeReduction(input PayrollInput) int {
	annualGrossSalarywithExtras := AnnualGrossSalaryWithExtras(input)
	switch {
	case annualGrossSalarywithExtras <= ReductionThresholdStart:
		return ReductionMax
	case annualGrossSalarywithExtras >= ReductionThresholdEnd:
		return ReductionMin
	default:
		return ReductionMax - ((annualGrossSalarywithExtras - ReductionThresholdStart) * ReductionSlope)
	}
}

// PersonalAndFamilyMinimum returns the total personal and family tax-free allowance.
func PersonalAndFamilyMinimum(input PayrollInput) int {
	min := PersonalMin

	// Dependents (children)
	for i := 0; i < input.NumberOfChildren; i++ {
		switch i {
		case 0:
			min += FirstChildMin
		case 1:
			min += SecondChildMin
		case 2:
			min += ThirdChildMin
		default:
			min += FourthChildMin
		}
	}

	// Disability of the taxpayer
	if input.HasDisability {
		if input.HasSevereDisability {
			min += DisabilityMin_Severe
		} else {
			min += DisabilityMin_Light
		}
		if input.NeedsAssistance {
			min += DisabilityAssistance
		}
	}

	// Ascendants
	if input.HasAscendantsOver65 {
		min += AscendantMin_65
	}
	if input.HasDisabledAscendants {
		min += AscendantMin_Disabled
	}

	return min
}

// TaxableBase returns the net taxable income after reductions and allowances.
func TaxableBase(input PayrollInput) int {
	annualGross := AnnualGrossSalaryWithExtras(input)
	reduction := WorkIncomeReduction(input)
	minimum := PersonalAndFamilyMinimum(input)

	base := annualGross - reduction - minimum
	if base < 0 {
		return 0
	}
	return base
}

// AnnualIrpf returns the total annual IRPF and effective tax rate based on progressive brackets.
func AnnualIrpf(input PayrollInput) (int, int) {
	base := TaxableBase(input)
	if base == 0 {
		return 0, 0
	}

	var tax int
	var previousLimit int = 0

	for _, bracket := range TaxBrackets {
		if base <= bracket.Limit {
			tax += (base - previousLimit) * bracket.Rate / 10000
			break
		} else {
			tax += (bracket.Limit - previousLimit) * bracket.Rate / 10000
			previousLimit = bracket.Limit
		}
	}

	effectiveRate := tax * 10000 / base
	return tax, effectiveRate
}
