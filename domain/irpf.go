package domain

const (
	ReductionMax            = 5565.0
	ReductionMin            = 2000.0
	ReductionThresholdStart = 14047.5
	ReductionThresholdEnd   = 17707.5
	ReductionSlope          = 0.34
)

const (
	// Personal minimum
	PersonalMin = 5550.0

	// Minimums per dependent (children)
	FirstChildMin  = 2400.0
	SecondChildMin = 2700.0
	ThirdChildMin  = 4000.0
	FourthChildMin = 4500.0

	// Disability minimums for the taxpayer
	DisabilityMin_Light  = 3000.0 // Disability degree < 65%
	DisabilityMin_Severe = 9000.0 // Disability degree ≥ 65%
	DisabilityAssistance = 3000.0 // If assistance from a third party is required

	// Minimums for ascendants over 65 years old or disabled
	AscendantMin_65       = 1150.0 // Over 65 years old
	AscendantMin_Disabled = 2550.0 // Disabled
)

func ReductionForWorkPerformance(input PayrollInput) float64 {
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

func PersonalAndFamilyMinimum(input PayrollInput) float64 {
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
