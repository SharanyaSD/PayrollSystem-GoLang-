// Code generated by mockery v2.41.0. DO NOT EDIT.

package mocks

import (
	dto "github.com/SharanyaSD/Payroll-GoLang.git/internal/pkg/dto"
	mock "github.com/stretchr/testify/mock"

	repository "github.com/SharanyaSD/Payroll-GoLang.git/internal/repository"
)

// PayrollService is an autogenerated mock type for the PayrollService type
type PayrollService struct {
	mock.Mock
}

// CreatePayroll provides a mock function with given fields: payrollDetails
func (_m *PayrollService) CreatePayroll(payrollDetails dto.CreatePayrollRequest) (repository.Payroll, error) {
	ret := _m.Called(payrollDetails)

	if len(ret) == 0 {
		panic("no return value specified for CreatePayroll")
	}

	var r0 repository.Payroll
	var r1 error
	if rf, ok := ret.Get(0).(func(dto.CreatePayrollRequest) (repository.Payroll, error)); ok {
		return rf(payrollDetails)
	}
	if rf, ok := ret.Get(0).(func(dto.CreatePayrollRequest) repository.Payroll); ok {
		r0 = rf(payrollDetails)
	} else {
		r0 = ret.Get(0).(repository.Payroll)
	}

	if rf, ok := ret.Get(1).(func(dto.CreatePayrollRequest) error); ok {
		r1 = rf(payrollDetails)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPayroll provides a mock function with given fields:
func (_m *PayrollService) GetPayroll() ([]dto.Payroll, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetPayroll")
	}

	var r0 []dto.Payroll
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]dto.Payroll, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []dto.Payroll); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dto.Payroll)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewPayrollService creates a new instance of PayrollService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPayrollService(t interface {
	mock.TestingT
	Cleanup(func())
}) *PayrollService {
	mock := &PayrollService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
