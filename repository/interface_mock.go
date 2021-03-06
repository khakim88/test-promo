// Code generated by MockGen. DO NOT EDIT.
// Source: repository/interface.go

// Package repository is a generated GoMock package.
package repository

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/khakim88/test-promo/model"
)

// MockDBReaderWriter is a mock of DBReaderWriter interface.
type MockDBReaderWriter struct {
	ctrl     *gomock.Controller
	recorder *MockDBReaderWriterMockRecorder
}

// MockDBReaderWriterMockRecorder is the mock recorder for MockDBReaderWriter.
type MockDBReaderWriterMockRecorder struct {
	mock *MockDBReaderWriter
}

// NewMockDBReaderWriter creates a new mock instance.
func NewMockDBReaderWriter(ctrl *gomock.Controller) *MockDBReaderWriter {
	mock := &MockDBReaderWriter{ctrl: ctrl}
	mock.recorder = &MockDBReaderWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDBReaderWriter) EXPECT() *MockDBReaderWriterMockRecorder {
	return m.recorder
}

// GetProductBySKU mocks base method.
func (m *MockDBReaderWriter) GetProductBySKU(sku string) (*model.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductBySKU", sku)
	ret0, _ := ret[0].(*model.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductBySKU indicates an expected call of GetProductBySKU.
func (mr *MockDBReaderWriterMockRecorder) GetProductBySKU(sku interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductBySKU", reflect.TypeOf((*MockDBReaderWriter)(nil).GetProductBySKU), sku)
}
