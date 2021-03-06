// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	io "io"

	mock "github.com/stretchr/testify/mock"
)

// Logger is an autogenerated mock type for the Logger type
type Logger struct {
	mock.Mock
}

// Fatal provides a mock function with given fields: v
func (_m *Logger) Fatal(v ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, v...)
	_m.Called(_ca...)
}

// Fatalf provides a mock function with given fields: format, v
func (_m *Logger) Fatalf(format string, v ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, v...)
	_m.Called(_ca...)
}

// Fatalln provides a mock function with given fields: v
func (_m *Logger) Fatalln(v ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, v...)
	_m.Called(_ca...)
}

// Panic provides a mock function with given fields: v
func (_m *Logger) Panic(v ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, v...)
	_m.Called(_ca...)
}

// Panicf provides a mock function with given fields: format, v
func (_m *Logger) Panicf(format string, v ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, v...)
	_m.Called(_ca...)
}

// Panicln provides a mock function with given fields: v
func (_m *Logger) Panicln(v ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, v...)
	_m.Called(_ca...)
}

// Print provides a mock function with given fields: v
func (_m *Logger) Print(v ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, v...)
	_m.Called(_ca...)
}

// Printf provides a mock function with given fields: format, v
func (_m *Logger) Printf(format string, v ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, v...)
	_m.Called(_ca...)
}

// Println provides a mock function with given fields: v
func (_m *Logger) Println(v ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, v...)
	_m.Called(_ca...)
}

// SetOutput provides a mock function with given fields: w
func (_m *Logger) SetOutput(w io.Writer) {
	_m.Called(w)
}
