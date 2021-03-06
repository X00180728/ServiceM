// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/openservicemesh/osm/pkg/certificate (interfaces: Manager)

// Package certificate is a generated GoMock package.
package certificate

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
)

// MockManager is a mock of Manager interface.
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *MockManagerMockRecorder
}

// MockManagerMockRecorder is the mock recorder for MockManager.
type MockManagerMockRecorder struct {
	mock *MockManager
}

// NewMockManager creates a new mock instance.
func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &MockManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockManager) EXPECT() *MockManagerMockRecorder {
	return m.recorder
}

// GetCertificate mocks base method.
func (m *MockManager) GetCertificate(arg0 CommonName) (*Certificate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCertificate", arg0)
	ret0, _ := ret[0].(*Certificate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCertificate indicates an expected call of GetCertificate.
func (mr *MockManagerMockRecorder) GetCertificate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCertificate", reflect.TypeOf((*MockManager)(nil).GetCertificate), arg0)
}

// GetRootCertificate mocks base method.
func (m *MockManager) GetRootCertificate() (*Certificate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRootCertificate")
	ret0, _ := ret[0].(*Certificate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRootCertificate indicates an expected call of GetRootCertificate.
func (mr *MockManagerMockRecorder) GetRootCertificate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRootCertificate", reflect.TypeOf((*MockManager)(nil).GetRootCertificate))
}

// IssueCertificate mocks base method.
func (m *MockManager) IssueCertificate(arg0 CommonName, arg1 time.Duration) (*Certificate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IssueCertificate", arg0, arg1)
	ret0, _ := ret[0].(*Certificate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IssueCertificate indicates an expected call of IssueCertificate.
func (mr *MockManagerMockRecorder) IssueCertificate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IssueCertificate", reflect.TypeOf((*MockManager)(nil).IssueCertificate), arg0, arg1)
}

// ListCertificates mocks base method.
func (m *MockManager) ListCertificates() ([]*Certificate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCertificates")
	ret0, _ := ret[0].([]*Certificate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCertificates indicates an expected call of ListCertificates.
func (mr *MockManagerMockRecorder) ListCertificates() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCertificates", reflect.TypeOf((*MockManager)(nil).ListCertificates))
}

// ReleaseCertificate mocks base method.
func (m *MockManager) ReleaseCertificate(arg0 CommonName) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReleaseCertificate", arg0)
}

// ReleaseCertificate indicates an expected call of ReleaseCertificate.
func (mr *MockManagerMockRecorder) ReleaseCertificate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReleaseCertificate", reflect.TypeOf((*MockManager)(nil).ReleaseCertificate), arg0)
}

// RotateCertificate mocks base method.
func (m *MockManager) RotateCertificate(arg0 CommonName) (*Certificate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RotateCertificate", arg0)
	ret0, _ := ret[0].(*Certificate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RotateCertificate indicates an expected call of RotateCertificate.
func (mr *MockManagerMockRecorder) RotateCertificate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RotateCertificate", reflect.TypeOf((*MockManager)(nil).RotateCertificate), arg0)
}
