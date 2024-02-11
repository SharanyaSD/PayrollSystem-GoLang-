package repository

import (
	"fmt"

	"github.com/SharanyaSD/PayrollSystem.git/internal/repository"
	"github.com/jmoiron/sqlx"
)

type PayrollStore struct {
	Db *sqlx.DB
}

func NewPayrollRepo(db *sqlx.DB) repository.PayrollStorer {
	return &PayrollStore{
		Db: db,
	}
}

func (pr *PayrollStore) CreatePayroll(payroll repository.Payroll) (repository.Payroll, error) {

	_, err := pr.Db.Exec("INSERT INTO payroll (ID, salary, net_pay_salary, pay_date) VALUES ($1, $2, $3, $4)",
		payroll.ID, payroll.Salary, payroll.NetPaySalary, payroll.PayDate)
	fmt.Println("ID passed ", payroll.ID)
	if err != nil {
		return repository.Payroll{}, err
	}
	return payroll, nil
}

func (pr *PayrollStore) GetEarningsByEmpoyeeID(ID string) (repository.Earnings, error) {
	var earning repository.Earnings
	query := "SELECT * from earnings where id=$1"
	fmt.Println("SQL Query:", query)
	row := pr.Db.QueryRow(query, ID)
	err := row.Scan(
		&earning.ID, &earning.Basic, &earning.HRA, &earning.DA, &earning.SA, &earning.CA,
		&earning.Bonus, &earning.GrossPay,
	)
	if err != nil {
		return earning, err
	}
	return earning, nil
}

func (pr *PayrollStore) GetDeductionsByEmpoyeeID(ID string) (repository.Deductions, error) {
	var deduction repository.Deductions
	query := "SELECT * from deduction where id=$1"
	fmt.Println("SQL Query:", query)
	row := pr.Db.QueryRow(query, ID)
	err := row.Scan(
		&deduction.ID, &deduction.TDS, &deduction.PF, &deduction.Medical, &deduction.GrossDeduction,
	)
	if err != nil {
		return deduction, err
	}
	return deduction, nil
}

func (pr *PayrollStore) GetPayroll() ([]repository.Payroll, error) {

	var payrolls []repository.Payroll
	err := pr.Db.Select(&payrolls, "SELECT * FROM payroll")
	if err != nil {
		return nil, err
	}
	return payrolls, nil
}

// func (pr *PayrollStore) GetPayrollByID(ID string) (repository.Payroll, error) {
// 	var payroll repository.Payroll
// 	query := "SELECT * FROM Payroll WHERE id=$1"
// 	fmt.Println("SQL Query:", query)
// 	row := pr.Db.QueryRow(query, ID)
// 	err := row.Scan(
// 		&payroll.ID, &payroll.Salary, &payroll.NetPaySalary, &payroll.PayDate,
// 	)
// 	if err != nil {
// 		return payroll, err
// 	}
// 	return payroll, nil

// }
