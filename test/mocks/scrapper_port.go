// Code generated by MockGen. DO NOT EDIT.
// Source: path/to/your/project/scrapper_port.go

// Package mock_ports is a generated GoMock package.
package mock_ports

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	"github.com/unawaretub86/club-hub/src/core/domain"
)

// MockScrapperPort is a mock of ScrapperPort interface.
type MockScrapperPort struct {
	ctrl     *gomock.Controller
	recorder *MockScrapperPortMockRecorder
}

// MockScrapperPortMockRecorder is the mock recorder for MockScrapperPort.
type MockScrapperPortMockRecorder struct {
	mock *MockScrapperPort
}

// NewMockScrapperPort creates a new mock instance.
func NewMockScrapperPort(ctrl *gomock.Controller) *MockScrapperPort {
	mock := &MockScrapperPort{ctrl: ctrl}
	mock.recorder = &MockScrapperPortMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockScrapperPort) EXPECT() *MockScrapperPortMockRecorder {
	return m.recorder
}

// ScrapCompanyData mocks base method.
func (m *MockScrapperPort) ScrapCompanyData(franchises []domain.Franchise) ([]domain.FranchiseScrapData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScrapCompanyData", franchises)
	ret0, _ := ret[0].([]domain.FranchiseScrapData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ScrapCompanyData indicates an expected call of ScrapCompanyData.
func (mr *MockScrapperPortMockRecorder) ScrapCompanyData(franchises interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScrapCompanyData", reflect.TypeOf((*MockScrapperPort)(nil).ScrapCompanyData), franchises)
}