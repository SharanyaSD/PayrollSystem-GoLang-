package repository

// "context"

// "github.com/SharanyaSD/PayrollSystem.git/repository"

type EmployeeStorer interface {
	GetAllEmployees() ([]Employee, error)
	GetEmployeeByID(ID string) (Employee, error)
	CreateEmployee(emp Employee) (Employee, error)
	//UpdateEmployee(ID string) (Employee, error)
	DeleteEmployee(ID string) (Employee, error)
}

type PayrollStorer interface {
	CreatePayroll(payroll Payroll) (Payroll, error) // affecting original data everytime made changes *payroll
	GetPayroll() ([]Payroll, error)
	GetEarningsByEmpoyeeID(ID string) (Earnings, error)
	GetDeductionsByEmpoyeeID(ID string) (Deductions, error)
}

// type EarningsStorer interface {
// 	GetEarningsByEmpoyeeID(ID string) (Earnings, error)
// }

// type DeductionStorer interface {
// 	GetDeductionsByEmpoyeeID(ID string) (Deductions, error)
// }

type Earnings struct {
	ID       string  `db:"id"`
	Basic    float64 `db:"basic"`
	HRA      float64 `db:"hra"`
	DA       float64 `db:"da"`
	SA       float64 `db:"sa"`
	CA       float64 `db:"ca"`
	Bonus    float64 `db:"bonus"`
	GrossPay float64 `db:"gross_pay"`
}

type Deductions struct {
	ID             string  `db:"id"`
	TDS            float64 `db:"tds"`
	PF             float64 `db:"pf"`
	Medical        float64 `db:"medical"`
	GrossDeduction float64 `db:"gross_deduction"`
}
