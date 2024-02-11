package app

import (

	//"github.com/SharanyaSD/PayrollSystem.git/repository"
	"github.com/SharanyaSD/PayrollSystem.git/internal/app/emp"
	"github.com/SharanyaSD/PayrollSystem.git/internal/app/payroll"
	"github.com/SharanyaSD/PayrollSystem.git/internal/pkg/email"
	repository "github.com/SharanyaSD/PayrollSystem.git/internal/repository/postgres"
	"github.com/jmoiron/sqlx"
)

type Dependencies struct {
	EmployeeService emp.Service
	PayrollService  payroll.PayrollService
	EmailService    email.Service
}

func NewServices(db *sqlx.DB, apiKey string) Dependencies {
	empRepo := repository.NewEmployeeRepo(db)
	earningsRepo := repository.NewPayrollRepo(db)
	deductionsRepo := repository.NewPayrollRepo(db)
	payrollRepo := repository.NewPayrollRepo(db)

	employeeService := emp.NewService(empRepo, earningsRepo, deductionsRepo)
	payrollService := payroll.NewPayrollService(payrollRepo, employeeService, email.NewEmailService(apiKey))

	return Dependencies{
		EmployeeService: employeeService,
		PayrollService:  payrollService,
		EmailService:    email.NewEmailService(apiKey),
	}

}
