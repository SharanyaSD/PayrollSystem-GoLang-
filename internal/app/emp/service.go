package emp

import (
	"github.com/SharanyaSD/PayrollSystem.git/internal/pkg/dto"
	"github.com/SharanyaSD/PayrollSystem.git/internal/repository"
)

type service struct {
	empRepo        repository.EmployeeStorer
	earningsRepo   repository.PayrollStorer
	deductionsRepo repository.PayrollStorer
}

type Service interface {
	GetAllEmployees() ([]dto.Employee, error)
	GetEmployeeByID(employee_id string) (dto.Employee, error)
	CreateEmployee(employeeDetails dto.CreateEmployeeRequest) (repository.Employee, error)
	// UpdateEmployee(ctx context.Context, id string) (dto.Employee, error)
	DeleteEmployee(id string) (dto.Employee, error)
	GetEarningsByEmpoyeeID(ID string) (repository.Earnings, error)
	GetDeductionsByEmpoyeeID(ID string) (repository.Deductions, error)
}

// type EmployeeService struct {
// }

func NewService(empRepo repository.EmployeeStorer, earningsRepo repository.PayrollStorer, deductionsRepo repository.PayrollStorer) Service {
	return &service{
		empRepo:        empRepo,
		earningsRepo:   earningsRepo,
		deductionsRepo: deductionsRepo,
	}
}

// func (es *service) CreateEmployee(employeeDetails dto.CreateEmployeeRequest) (employee repository.Employee, error)
func (es *service) CreateEmployee(employeeDetails dto.CreateEmployeeRequest) (repository.Employee, error) {
	empInfo := repository.Employee{
		ID:                 employeeDetails.ID,
		FirstName:          employeeDetails.FirstName,
		MiddleName:         employeeDetails.MiddleName,
		LastName:           employeeDetails.LastName,
		Email:              employeeDetails.Email,
		DateOfBirth:        employeeDetails.DateOfBirth,
		DateOfJoining:      employeeDetails.DateOfJoining,
		Designation:        employeeDetails.Designation,
		YearsOfExperience:  employeeDetails.YearsOfExperience,
		ProofId:            employeeDetails.ProofId,
		ResidentialAddress: employeeDetails.ResidentialAddress,
		HiredLocation:      employeeDetails.HiredLocation,
		RoleId:             employeeDetails.RoleId,
		WorkStatus:         employeeDetails.WorkStatus,
		Salary:             employeeDetails.Salary,
	}
	return es.empRepo.CreateEmployee(empInfo)
}

// func UpdateEmployee(ctx context.Context, id string) {

// }

func (es *service) DeleteEmployee(id string) (dto.Employee, error) {
	employee, err := es.empRepo.DeleteEmployee(id)
	if err != nil {
		return dto.Employee{}, err
	}

	// Converting to DTO
	dtoEmployee := dto.Employee{
		ID:                 employee.ID,
		FirstName:          employee.FirstName,
		MiddleName:         employee.MiddleName,
		LastName:           employee.LastName,
		Email:              employee.Email,
		DateOfBirth:        employee.DateOfBirth,
		DateOfJoining:      employee.DateOfJoining,
		Designation:        employee.Designation,
		YearsOfExperience:  employee.YearsOfExperience,
		ProofId:            employee.ProofId,
		ResidentialAddress: employee.ResidentialAddress,
		HiredLocation:      employee.HiredLocation,
		//	RoleId:             employee.RoleId,
		WorkStatus: employee.WorkStatus,
		Salary:     employee.Salary,
	}

	//DTO employee
	return dtoEmployee, nil
}

func (es *service) GetAllEmployees() ([]dto.Employee, error) {
	employees, err := es.empRepo.GetAllEmployees()
	if err != nil {
		return []dto.Employee{}, err
	}
	var dtoEmployees []dto.Employee
	for _, employee := range employees {
		dtoEmployee := dto.Employee{
			ID:                 employee.ID,
			FirstName:          employee.FirstName,
			MiddleName:         employee.MiddleName,
			LastName:           employee.LastName,
			Email:              employee.Email,
			DateOfBirth:        employee.DateOfBirth,
			DateOfJoining:      employee.DateOfJoining,
			Designation:        employee.Designation,
			YearsOfExperience:  employee.YearsOfExperience,
			ProofId:            employee.ProofId,
			ResidentialAddress: employee.ResidentialAddress,
			HiredLocation:      employee.HiredLocation,
			//		RoleId:             employee.RoleId,
			WorkStatus: employee.FirstName,
			Salary:     employee.Salary,
		}
		dtoEmployees = append(dtoEmployees, dtoEmployee)
	}
	return dtoEmployees, nil
}

func (es *service) GetEmployeeByID(id string) (dto.Employee, error) {

	employee, err := es.empRepo.GetEmployeeByID(id)
	if err != nil {
		return dto.Employee{}, err
	}

	// Converting to DTO
	dtoEmployee := dto.Employee{
		ID:                 employee.ID,
		FirstName:          employee.FirstName,
		MiddleName:         employee.MiddleName,
		LastName:           employee.LastName,
		Email:              employee.Email,
		DateOfBirth:        employee.DateOfBirth,
		DateOfJoining:      employee.DateOfJoining,
		Designation:        employee.Designation,
		YearsOfExperience:  employee.YearsOfExperience,
		ProofId:            employee.ProofId,
		ResidentialAddress: employee.ResidentialAddress,
		HiredLocation:      employee.HiredLocation,
		RoleId:             employee.RoleId,
		WorkStatus:         employee.WorkStatus,
		Salary:             employee.Salary,
	}

	//DTO employee
	return dtoEmployee, nil
}

func (es *service) GetEarningsByEmpoyeeID(ID string) (repository.Earnings, error) {
	earnings, err := es.earningsRepo.GetEarningsByEmpoyeeID(ID)
	if err != nil {
		return repository.Earnings{}, err
	}

	// hra := 0.4 * earnings.Basic
	// da := 0.15 * earnings.Basic

	grossPay := earnings.Basic + earnings.HRA + earnings.DA + earnings.SA + earnings.CA + earnings.Bonus
	// earnings.HRA = hra
	// earnings.DA = da
	earnings.GrossPay = grossPay

	return earnings, nil
}

func (es *service) GetDeductionsByEmpoyeeID(ID string) (repository.Deductions, error) {
	deductions, err := es.deductionsRepo.GetDeductionsByEmpoyeeID(ID)
	if err != nil {
		return repository.Deductions{}, err
	}
	GrossDeduction := deductions.TDS + deductions.PF + deductions.Medical

	deductions.GrossDeduction = GrossDeduction

	return deductions, nil
}
