// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	url "net/url"
)

// LinksInterface is an autogenerated mock type for the LinksInterface type
type LinksInterface struct {
	mock.Mock
}

// GetNextPage provides a mock function with given fields:
func (_m *LinksInterface) GetNextPage() {
	_m.Called()
}

// GetPrevPage provides a mock function with given fields:
func (_m *LinksInterface) GetPrevPage() {
	_m.Called()
}

// MakeFirstPage provides a mock function with given fields:
func (_m *LinksInterface) MakeFirstPage() {
	_m.Called()
}

// MakeLastPage provides a mock function with given fields:
func (_m *LinksInterface) MakeLastPage() {
	_m.Called()
}

// getKey provides a mock function with given fields:
func (_m *LinksInterface) getKey() url.Values {
	ret := _m.Called()

	var r0 url.Values
	if rf, ok := ret.Get(0).(func() url.Values); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(url.Values)
		}
	}

	return r0
}
