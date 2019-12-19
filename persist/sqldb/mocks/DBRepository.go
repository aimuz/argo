// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	db "upper.io/db.v3"

	v1alpha1 "github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"
)

// DBRepository is an autogenerated mock type for the DBRepository type
type DBRepository struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *DBRepository) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: condition
func (_m *DBRepository) Delete(condition db.Cond) error {
	ret := _m.Called(condition)

	var r0 error
	if rf, ok := ret.Get(0).(func(db.Cond) error); ok {
		r0 = rf(condition)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: uid
func (_m *DBRepository) Get(uid string) (*v1alpha1.Workflow, error) {
	ret := _m.Called(uid)

	var r0 *v1alpha1.Workflow
	if rf, ok := ret.Get(0).(func(string) *v1alpha1.Workflow); ok {
		r0 = rf(uid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.Workflow)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(uid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsNodeStatusOffload provides a mock function with given fields:
func (_m *DBRepository) IsNodeStatusOffload() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// List provides a mock function with given fields: orderBy
func (_m *DBRepository) List(orderBy string) (*v1alpha1.WorkflowList, error) {
	ret := _m.Called(orderBy)

	var r0 *v1alpha1.WorkflowList
	if rf, ok := ret.Get(0).(func(string) *v1alpha1.WorkflowList); ok {
		r0 = rf(orderBy)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.WorkflowList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(orderBy)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Query provides a mock function with given fields: condition, orderBy
func (_m *DBRepository) Query(condition db.Cond, orderBy ...interface{}) ([]v1alpha1.Workflow, error) {
	var _ca []interface{}
	_ca = append(_ca, condition)
	_ca = append(_ca, orderBy...)
	ret := _m.Called(_ca...)

	var r0 []v1alpha1.Workflow
	if rf, ok := ret.Get(0).(func(db.Cond, ...interface{}) []v1alpha1.Workflow); ok {
		r0 = rf(condition, orderBy...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]v1alpha1.Workflow)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(db.Cond, ...interface{}) error); ok {
		r1 = rf(condition, orderBy...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QueryWithPagination provides a mock function with given fields: condition, pageSize, lastID, orderBy
func (_m *DBRepository) QueryWithPagination(condition db.Cond, pageSize uint, lastID string, orderBy ...interface{}) (*v1alpha1.WorkflowList, error) {
	var _ca []interface{}
	_ca = append(_ca, condition, pageSize, lastID)
	_ca = append(_ca, orderBy...)
	ret := _m.Called(_ca...)

	var r0 *v1alpha1.WorkflowList
	if rf, ok := ret.Get(0).(func(db.Cond, uint, string, ...interface{}) *v1alpha1.WorkflowList); ok {
		r0 = rf(condition, pageSize, lastID, orderBy...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.WorkflowList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(db.Cond, uint, string, ...interface{}) error); ok {
		r1 = rf(condition, pageSize, lastID, orderBy...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: wf
func (_m *DBRepository) Save(wf *v1alpha1.Workflow) error {
	ret := _m.Called(wf)

	var r0 error
	if rf, ok := ret.Get(0).(func(*v1alpha1.Workflow) error); ok {
		r0 = rf(wf)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}