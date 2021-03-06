// Code generated by mockery v1.0.0. DO NOT EDIT.

package gocb

import (
	gocbcore "github.com/couchbase/gocbcore/v9"
	mock "github.com/stretchr/testify/mock"
)

// mockDiagnosticsProvider is an autogenerated mock type for the diagnosticsProvider type
type mockDiagnosticsProvider struct {
	mock.Mock
}

// Diagnostics provides a mock function with given fields: opts
func (_m *mockDiagnosticsProvider) Diagnostics(opts gocbcore.DiagnosticsOptions) (*gocbcore.DiagnosticInfo, error) {
	ret := _m.Called(opts)

	var r0 *gocbcore.DiagnosticInfo
	if rf, ok := ret.Get(0).(func(gocbcore.DiagnosticsOptions) *gocbcore.DiagnosticInfo); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gocbcore.DiagnosticInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(gocbcore.DiagnosticsOptions) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Ping provides a mock function with given fields: opts
func (_m *mockDiagnosticsProvider) Ping(opts gocbcore.PingOptions) (*gocbcore.PingResult, error) {
	ret := _m.Called(opts)

	var r0 *gocbcore.PingResult
	if rf, ok := ret.Get(0).(func(gocbcore.PingOptions) *gocbcore.PingResult); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gocbcore.PingResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(gocbcore.PingOptions) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
