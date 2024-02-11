package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SharanyaSD/PayrollSystem.git/internal/app"
	"github.com/gorilla/mux"
)

func NewRouter(deps app.Dependencies) *mux.Router {
	r := mux.NewRouter()
	//r.HandleFunc("/", ).Methods(http.MethodGet)
	r.HandleFunc("/getAllEmployees", GetAllEmployeesHandler(deps.EmployeeService)).Methods(http.MethodGet)
	r.HandleFunc("/getEmployeeByID", GetEmployeeByIDHandler(deps.EmployeeService)).Methods(http.MethodGet)
	r.HandleFunc("/createEmployee", CreateEmployeeHandler(deps.EmployeeService)).Methods(http.MethodPost)
	//r.HandleFunc("/updateEmployee", UpdateEmployee).Methods(http.MethodPut)
	r.HandleFunc("/deleteEmployee", DeleteEmployeeHandler(deps.EmployeeService)).Methods(http.MethodDelete)
	r.HandleFunc("/createPayroll", CreatePayrollHandler(deps.PayrollService)).Methods(http.MethodPost)
	r.HandleFunc("/getPayroll", GetPayrollHandler(deps.PayrollService)).Methods(http.MethodGet)
	//	r.HandleFunc("/getPayrollByID", GetPayrollByIDHandler(deps.PayrollService)).Methods(http.MethodGet)

	fmt.Printf("Starting server at 8080 ")
	log.Fatal(http.ListenAndServe(":8080", r))
	return r
}
