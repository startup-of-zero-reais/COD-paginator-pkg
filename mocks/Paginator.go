// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	reflect "reflect"

	paginator "github.com/startup-of-zero-reais/COD-paginator-pkg"
	mock "github.com/stretchr/testify/mock"
)

// Paginator is an autogenerated mock type for the Paginator type
type Paginator struct {
	mock.Mock
}

// Paginate provides a mock function with given fields: items, result
func (_m *Paginator) Paginate(items interface{}, result interface{}) (paginator.DataResultInterface, error) {
	ret := _m.Called(items, result)

	var r0 paginator.DataResultInterface
	if rf, ok := ret.Get(0).(func(interface{}, interface{}) paginator.DataResultInterface); ok {
		r0 = rf(items, result)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(paginator.DataResultInterface)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}, interface{}) error); ok {
		r1 = rf(items, result)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WithMeta provides a mock function with given fields: metadata
func (_m *Paginator) WithMeta(metadata *paginator.Metadata) paginator.Paginator {
	ret := _m.Called(metadata)

	var r0 paginator.Paginator
	if rf, ok := ret.Get(0).(func(*paginator.Metadata) paginator.Paginator); ok {
		r0 = rf(metadata)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(paginator.Paginator)
		}
	}

	return r0
}

// extractTags provides a mock function with given fields: tag, field
func (_m *Paginator) extractTags(tag string, field reflect.Value) error {
	ret := _m.Called(tag, field)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, reflect.Value) error); ok {
		r0 = rf(tag, field)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// paginateCollection provides a mock function with given fields:
func (_m *Paginator) paginateCollection() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// paginateSingle provides a mock function with given fields: items, result
func (_m *Paginator) paginateSingle(items interface{}, result interface{}) error {
	ret := _m.Called(items, result)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, interface{}) error); ok {
		r0 = rf(items, result)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// scanMode provides a mock function with given fields:
func (_m *Paginator) scanMode() (paginator.Mode, error) {
	ret := _m.Called()

	var r0 paginator.Mode
	if rf, ok := ret.Get(0).(func() paginator.Mode); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(paginator.Mode)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}