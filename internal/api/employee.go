package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/SharanyaSD/PayrollSystem.git/internal/app/emp"
	"github.com/SharanyaSD/PayrollSystem.git/internal/pkg/dto"
)

func CreateEmployeeHandler(empSvc emp.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		var req dto.CreateEmployeeRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		fmt.Printf("%+v", req)

		employeeInfo, err := empSvc.CreateEmployee(req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		responseJSON, err := json.Marshal(employeeInfo)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(responseJSON)
		return
	}
}

// func UpdateEmployee(empSvc emp.Service) func(w http.ResponseWriter, r *http.Request) {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		ctx := r.Context()
// 		var req dto.UpdateEmployeeRequest
// 		err := json.NewDecoder(r.Body).Decode(&req)
// 		if err != nil {
// 			w.WriteHeader(http.StatusBadRequest)
// 			return
// 		}

// 		err := empSvc.CreateEmployee(ctx, req)
// 		if err != nil {
// 			w.WriteHeader(http.StatusBadRequest)
// 			return
// 		}
// 		w.WriteHeader(http.StatusOK)
// 	}

// }

func DeleteEmployeeHandler(empSvc emp.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		empID := r.URL.Query().Get("id")
		if empID == "" {
			http.Error(w, "Employee ID is required", http.StatusBadRequest)
			return
		}

		deletedEmp, err := empSvc.DeleteEmployee(empID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		jsonData, err := json.Marshal(deletedEmp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Employee deleted successfully"))
		w.Write(jsonData)
	}
}

func GetAllEmployeesHandler(empSvc emp.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		response, err := empSvc.GetAllEmployees()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		responseJSON, err := json.Marshal(response)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(responseJSON)
		return

	}
}

func GetEmployeeByIDHandler(empSvc emp.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")

		employee, err := empSvc.GetEmployeeByID(id)
		if err != nil {
			// w.WriteHeader(http.StatusInternalServerError)
			// w.Write([]byte(err.Error()))
			// return
			http.Error(w, "Failed to get employee: "+err.Error(), http.StatusInternalServerError)
			return
		}

		responseJSON, err := json.Marshal(employee)
		if err != nil {
			// w.WriteHeader(http.StatusInternalServerError)
			// w.Write([]byte(err.Error()))
			// return
			http.Error(w, "Failed to serialize employee to JSON: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(responseJSON)

	}
}