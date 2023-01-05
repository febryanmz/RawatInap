// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	patient "github.com/KamarRS-App/KamarRS-App/features/patient"
	mock "github.com/stretchr/testify/mock"
)

// PatientRepo is an autogenerated mock type for the RepositoryInterface type
type PatientRepo struct {
	mock.Mock
}

// Create provides a mock function with given fields: input
func (_m *PatientRepo) Create(input patient.CorePatient) error {
	ret := _m.Called(input)

	var r0 error
	if rf, ok := ret.Get(0).(func(patient.CorePatient) error); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteById provides a mock function with given fields: id
func (_m *PatientRepo) DeleteById(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllPatient provides a mock function with given fields:
func (_m *PatientRepo) GetAllPatient() ([]patient.CorePatient, error) {
	ret := _m.Called()

	var r0 []patient.CorePatient
	if rf, ok := ret.Get(0).(func() []patient.CorePatient); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]patient.CorePatient)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByPatientId provides a mock function with given fields: id
func (_m *PatientRepo) GetByPatientId(id int) (patient.CorePatient, error) {
	ret := _m.Called(id)

	var r0 patient.CorePatient
	if rf, ok := ret.Get(0).(func(int) patient.CorePatient); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(patient.CorePatient)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByUserId provides a mock function with given fields: limit, offset, id
func (_m *PatientRepo) GetByUserId(limit int, offset int, id int) ([]patient.CorePatient, int, error) {
	ret := _m.Called(limit, offset, id)

	var r0 []patient.CorePatient
	if rf, ok := ret.Get(0).(func(int, int, int) []patient.CorePatient); ok {
		r0 = rf(limit, offset, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]patient.CorePatient)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(int, int, int) int); ok {
		r1 = rf(limit, offset, id)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(int, int, int) error); ok {
		r2 = rf(limit, offset, id)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Update provides a mock function with given fields: id, userId, input
func (_m *PatientRepo) Update(id int, userId int, input patient.CorePatient) error {
	ret := _m.Called(id, userId, input)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, int, patient.CorePatient) error); ok {
		r0 = rf(id, userId, input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewPatientRepo interface {
	mock.TestingT
	Cleanup(func())
}

// NewPatientRepo creates a new instance of PatientRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPatientRepo(t mockConstructorTestingTNewPatientRepo) *PatientRepo {
	mock := &PatientRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
