package domain

// Worker contribution rates
const (
	WorkerCCRate       = 470 // 4.70% General contingencies (common illness, retirement, etc.)
	WorkerUnemployment = 155 // 1.55% Unemployment insurance
	WorkerTraining     = 10  // 0.10% Vocational training
)

// Employer contribution rates
const (
	EmployerCCRate       = 2360 // General contingencies
	EmployerUnemployment = 550  // Unemployment insurance
	EmployerTraining     = 60   // Vocational training
	EmployerFogasa       = 20   // FOGASA (Wage Guarantee Fund)
	EmployerAccidentRate = 150  // 1.50% Work-related accidents and occupational diseases (CNAE dependent)
)

// SSCotisations represents all Social Security contributions per payroll,
// split by concept and by contributor (worker vs employer).
type SSCotisations struct {
	// Worker contributions
	WorkerCC           int // General contingencies (worker)
	WorkerUnemployment int // Unemployment (worker)
	WorkerTraining     int // Training (worker)

	// Employer contributions
	EmployerCC           int // General contingencies (employer)
	EmployerUnemployment int // Unemployment (employer)
	EmployerTraining     int // Training (employer)
	EmployerFogasa       int // FOGASA (employer)
	EmployerAccidents    int // Occupational accidents/diseases (employer)

	// Totals
	TotalWorker   int // Total amount paid by the worker
	TotalEmployer int // Total amount paid by the employer
	Total         int // Global total of Social Security contributions
}

// CalculateSSCotisations returns the full breakdown of Social Security contributions
// based on BCCC (common contingencies base) and BCCP (professional contingencies base).
func CalculateSSCotisations(input PayrollInput) SSCotisations {
	bccc := BCCC(input) // returns base in cents
	bccp := BCCP(input) // returns base in cents

	// Worker contributions
	workerCC := bccc * WorkerCCRate / 10000
	workerUnemployment := bccc * WorkerUnemployment / 10000
	workerTraining := bccc * WorkerTraining / 10000

	// Employer contributions
	employerCC := bccc * EmployerCCRate / 10000
	employerUnemployment := bccc * EmployerUnemployment / 10000
	employerTraining := bccc * EmployerTraining / 10000
	employerFogasa := bccc * EmployerFogasa / 10000
	employerAccidents := bccp * EmployerAccidentRate / 10000

	totalWorker := workerCC + workerUnemployment + workerTraining
	totalEmployer := employerCC + employerUnemployment + employerTraining + employerFogasa + employerAccidents
	total := totalWorker + totalEmployer

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
		TotalWorker:   totalWorker,
		TotalEmployer: totalEmployer,
		Total:         total,
	}
}
