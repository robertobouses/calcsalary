package domain

// Worker contribution rates
const (
	WorkerCCRate       = 0.047  // General contingencies (common illness, retirement, etc.)
	WorkerUnemployment = 0.0155 // Unemployment insurance
	WorkerTraining     = 0.001  // Vocational training
)

// Employer contribution rates
const (
	EmployerCCRate       = 0.236 // General contingencies
	EmployerUnemployment = 0.055 // Unemployment insurance
	EmployerTraining     = 0.006 // Vocational training
	EmployerFogasa       = 0.002 // FOGASA (Wage Guarantee Fund)
	EmployerAccidentRate = 0.015 // Work-related accidents and occupational diseases (CNAE dependent)
)

// SSCotisations represents all Social Security contributions per payroll,
// split by concept and by contributor (worker vs employer).
type SSCotisations struct {
	// Worker contributions
	WorkerCC           float64 // General contingencies (worker)
	WorkerUnemployment float64 // Unemployment (worker)
	WorkerTraining     float64 // Training (worker)

	// Employer contributions
	EmployerCC           float64 // General contingencies (employer)
	EmployerUnemployment float64 // Unemployment (employer)
	EmployerTraining     float64 // Training (employer)
	EmployerFogasa       float64 // FOGASA (employer)
	EmployerAccidents    float64 // Occupational accidents/diseases (employer)

	// Totals
	TotalWorker   float64 // Total amount paid by the worker
	TotalEmployer float64 // Total amount paid by the employer
	Total         float64 // Global total of Social Security contributions
}

// CalculateSSCotisations returns the full breakdown of Social Security contributions
// based on BCCC (common contingencies base) and BCCP (professional contingencies base).
func CalculateSSCotisations(input PayrollInput) SSCotisations {
	bccc := BCCC(input)
	bccp := BCCP(input)

	// Worker contributions
	workerCC := bccc * WorkerCCRate
	workerUnemployment := bccc * WorkerUnemployment
	workerTraining := bccc * WorkerTraining

	// Employer contributions
	employerCC := bccc * EmployerCCRate
	employerUnemployment := bccc * EmployerUnemployment
	employerTraining := bccc * EmployerTraining
	employerFogasa := bccc * EmployerFogasa
	employerAccidents := bccp * EmployerAccidentRate

	return SSCotisations{
		// Worker
		WorkerCC:           workerCC,
		WorkerUnemployment: workerUnemployment,
		WorkerTraining:     workerTraining,

		// Employer
		EmployerCC:           employerCC,
		EmployerUnemployment: employerUnemployment,
		EmployerTraining:     employerTraining,
		EmployerFogasa:       employerFogasa,
		EmployerAccidents:    employerAccidents,

		// Totals
		TotalWorker:   workerCC + workerUnemployment + workerTraining,
		TotalEmployer: employerCC + employerUnemployment + employerTraining + employerFogasa + employerAccidents,
		Total:         workerCC + workerUnemployment + workerTraining + employerCC + employerUnemployment + employerTraining + employerFogasa + employerAccidents,
	}
}
