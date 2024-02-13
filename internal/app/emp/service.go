package emp

import (
	"errors"
	"time"

	"github.com/SharanyaSD/Payroll-GoLang.git/internal/pkg/dto"
	"github.com/SharanyaSD/Payroll-GoLang.git/internal/repository"
	"github.com/golang-jwt/jwt"
)

type service struct {
	empRepo        repository.EmployeeStorer
	earningsRepo   repository.EarningsStorer
	deductionsRepo repository.DeductionsStorer
}

type Service interface {
	GetAllEmployees() ([]dto.Employee, error)
	GetEmployeeByID(employee_id string) (dto.Employee, error)
	CreateEmployee(employeeDetails dto.CreateEmployeeRequest) (repository.Employee, error)
	DeleteEmployee(id string) (dto.Employee, error)
	GetEarningsByEmpoyeeID(ID string) (repository.Earnings, error)
	GetDeductionsByEmpoyeeID(ID string) (repository.Deductions, error)
	Login(username, password string) (string, error)
	InsertEarnings(earnings repository.Earnings) (repository.Earnings, error)
	InsertDeductions(deductions repository.Deductions) (repository.Deductions, error)
	GetEmployeeByEmail(email string) (dto.Employee, error)
}

var jwtKey = []byte("keymaker")

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	RoleID   string `json:"role_id"`
}

type Claims struct {
	Email string `json:"email"`
	ID    string `json:"id"`
	//Password string `json:"password"`
	RoleID int `json:"role_id"`
	jwt.StandardClaims
}

// type EmployeeService struct {
// }

func NewService(empRepo repository.EmployeeStorer, earningsRepo repository.EarningsStorer, deductionsRepo repository.DeductionsStorer) Service {
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
		Password:           employeeDetails.Password,
	}

	createdEmployee, err := es.empRepo.CreateEmployee(empInfo)
	if err != nil {
		return repository.Employee{}, err
	}

	defaultEarnings := earningsMap[empInfo.Designation]
	defaultEarnings.ID = empInfo.ID
	_, err = es.earningsRepo.InsertEarnings(defaultEarnings)
	if err != nil {
		return repository.Employee{}, err
	}

	defaultDeductions := deductionsMap[empInfo.Designation]
	defaultDeductions.ID = empInfo.ID
	_, err = es.deductionsRepo.InsertDeductions(defaultDeductions)
	if err != nil {
		return repository.Employee{}, err
	}

	return createdEmployee, nil
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
		Password:   employee.Password,
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
			Password:   employee.Password,
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
		Password:           employee.Password,
	}

	//DTO employee
	return dtoEmployee, nil
}

const (
	ManagerBasicEarning = 1000.0
	ManagerHRA          = 0.4 * ManagerBasicEarning
	ManagerDA           = 0.15 * ManagerBasicEarning
	ManagerSA           = 500.0
	ManagerCA           = 200.0
	ManagerBonus        = 500.0
	ManagerTDS          = 6000.0
	ManagerPF           = 0.12 * ManagerBasicEarning
	ManagerMedical      = 3000.0

	DefaultBasicEarning = 800.0
	DefaultHRA          = 0.3 * DefaultBasicEarning
	DefaultDA           = 0.12 * DefaultBasicEarning
	DefaultSA           = 500.0
	DefaultCA           = 200.0
	DefaultBonus        = 150.0
	DefaultTDS          = 5000.0
	DefaultPF           = 0.1 * DefaultBasicEarning
	DefaultMedical      = 2000.0
)

var earningsMap = map[string]repository.Earnings{
	"Manager": {
		Basic:    ManagerBasicEarning,
		HRA:      ManagerHRA,
		DA:       ManagerDA,
		SA:       ManagerSA,
		CA:       ManagerCA,
		Bonus:    ManagerBonus,
		GrossPay: ManagerBasicEarning + ManagerHRA + ManagerDA + ManagerSA + ManagerCA + ManagerBonus,
	},
	"Employee": {
		Basic:    DefaultBasicEarning,
		HRA:      DefaultHRA,
		DA:       DefaultDA,
		SA:       DefaultSA,
		CA:       DefaultCA,
		Bonus:    DefaultBonus,
		GrossPay: DefaultBasicEarning + DefaultHRA + DefaultDA + DefaultSA + DefaultCA + DefaultBonus,
	},
}

var deductionsMap = map[string]repository.Deductions{
	"Manager": {
		TDS:            ManagerTDS,
		PF:             ManagerPF,
		Medical:        ManagerMedical,
		GrossDeduction: ManagerTDS + ManagerPF + ManagerMedical,
	},
	"Employee": {
		TDS:            DefaultTDS,
		PF:             DefaultPF,
		Medical:        DefaultMedical,
		GrossDeduction: DefaultTDS + DefaultPF + DefaultMedical,
	},
}

func (es *service) GetEarningsByEmpoyeeID(ID string) (repository.Earnings, error) {
	employee, err := es.empRepo.GetEmployeeByID(ID)
	if err != nil {
		return repository.Earnings{}, err
	}

	earnings, err := es.earningsRepo.GetEarningsByEmpoyeeID(ID)
	if err != nil {
		return repository.Earnings{}, err
	}

	var designation string
	if employee.Designation != "" {
		designation = employee.Designation
	} else {
		designation = "Default"
	}

	switch designation {
	case "Manager":
		earnings.Basic = ManagerBasicEarning
		earnings.HRA = ManagerHRA
		earnings.DA = ManagerDA
		earnings.SA = ManagerSA
		earnings.CA = ManagerCA
		earnings.Bonus = ManagerBonus

	default:
		earnings.Basic = DefaultBasicEarning
		earnings.HRA = DefaultHRA
		earnings.DA = DefaultDA
		earnings.SA = DefaultSA
		earnings.CA = DefaultCA
		earnings.Bonus = DefaultBonus
	}
	grossPay := earnings.Basic + earnings.HRA + earnings.DA + earnings.SA + earnings.CA + earnings.Bonus
	// earnings.HRA = hra
	// earnings.DA = da
	earnings.GrossPay = grossPay
	return earnings, nil
}

func (es *service) GetDeductionsByEmpoyeeID(ID string) (repository.Deductions, error) {

	employee, err := es.empRepo.GetEmployeeByID(ID)
	if err != nil {
		return repository.Deductions{}, err
	}

	deductions, err := es.deductionsRepo.GetDeductionsByEmpoyeeID(ID)
	if err != nil {
		return repository.Deductions{}, err
	}

	var designation string
	if employee.Designation != "" {
		designation = employee.Designation
	} else {
		designation = "Default"
	}

	switch designation {
	case "Manager":
		deductions.TDS = ManagerTDS
		deductions.PF = ManagerPF
		deductions.Medical = ManagerMedical
	default:
		deductions.TDS = DefaultTDS
		deductions.PF = DefaultPF
		deductions.Medical = DefaultMedical
	}

	//greter deduction if -  is manager
	// if employee.Salary > 10000 {
	// 	additionalTax := 0.1 * employee.Salary //10K tax

	// }
	grossDeduction := deductions.TDS + deductions.PF + deductions.Medical

	deductions.GrossDeduction = grossDeduction

	return deductions, nil
}

func (es *service) InsertEarnings(earnings repository.Earnings) (repository.Earnings, error) {
	insertedEarnings, err := es.earningsRepo.InsertEarnings(earnings)
	if err != nil {
		return repository.Earnings{}, err
	}
	return insertedEarnings, nil
}

func (es *service) InsertDeductions(deductions repository.Deductions) (repository.Deductions, error) {
	insertedDeductions, err := es.deductionsRepo.InsertDeductions(deductions)
	if err != nil {
		return repository.Deductions{}, err
	}
	return insertedDeductions, nil
}

func (es *service) GetEmployeeByEmail(email string) (dto.Employee, error) {
	employee, err := es.empRepo.GetEmployeeByEmail(email)
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
		Password:           employee.Password,
	}

	//DTO employee
	return dtoEmployee, nil
}

func (es *service) Login(email, password string) (string, error) {

	emp, err := es.empRepo.GetEmployeeByEmail(email)
	if err != nil {
		return "", err
	}

	// expectedPassword, ok := users[username]
	if emp.Password != password {
		return "", errors.New("invalid email or password")
	}

	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &Claims{
		Email:  email,
		ID:     emp.ID,
		RoleID: emp.RoleId,
		//Password: emp.Password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
