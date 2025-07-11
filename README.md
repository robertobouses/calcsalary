# calcsalary


How to Import and Use the Payroll Library in Your Go Project

1. Initialize your Go module (if you haven't already)
In your project folder, run:
go mod init github.com/yourusername/your-project-name
Replace the path with your own project repository.

2. Import the Payroll Library module
Add the import statement in your Go code where you want to use it:
import "github.com/robertobouses/payroll-library/domain"
Replace "github.com/robertobouses/payroll-library" with the actual path of your library repository.

3. Use the library functions
Example usage:
```go
package main

import (
    "fmt"
    "github.com/robertobouses/payroll-library/domain"
)

func main() {
    input := domain.PayrollInput{
        BaseSalary:            1500.0,
        SalaryComplements:     []float64{200.0, 100.0},
        PersonalComplement:    50.0,
        NumberOfExtraPayments: 2,
        NumberOfChildren:      2,
        NumberOfExtraHours:    10,
        HasDisability:         false,
        HasSevereDisability:   false,
        NeedsAssistance:       false,
        HasAscendantsOver65:   true,
        HasDisabledAscendants: false,
    }

    ss := domain.CalculateSSCotisations(input)
    fmt.Printf("Worker SS contribution: %.2f\n", ss.TotalWorker)
    fmt.Printf("Employer SS contribution: %.2f\n", ss.TotalEmployer)

    tax, rate := domain.Irpf(input)
    fmt.Printf("Annual IRPF: %.2f, Effective Rate: %.2f%%\n", tax, rate*100)

    payroll := domain.GeneratePayrollBreakdown(input)
    fmt.Printf("Net Salary: %.2f\n", payroll.NetSalary)
}
```

4. Download dependencies
Run this command to download your library and its dependencies:
go mod tidy

5. (Optional) Use your library locally during development
If you want to use your local copy of the library before publishing it to a repository, run:
go mod edit -replace github.com/robertobouses/payroll-library=../path/to/your/local/library
go mod tidy
That’s all! Now you can start using the payroll library functions in your Go projects.

