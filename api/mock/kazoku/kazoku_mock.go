// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/issho-ni/issho/api/kazoku (interfaces: KazokuClient)

// Package mock_kazoku is a generated GoMock package.
package mock_kazoku

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	kazoku "github.com/issho-ni/issho/api/kazoku"
	grpc "google.golang.org/grpc"
)

// MockKazokuClient is a mock of KazokuClient interface
type MockKazokuClient struct {
	ctrl     *gomock.Controller
	recorder *MockKazokuClientMockRecorder
}

// MockKazokuClientMockRecorder is the mock recorder for MockKazokuClient
type MockKazokuClientMockRecorder struct {
	mock *MockKazokuClient
}

// NewMockKazokuClient creates a new mock instance
func NewMockKazokuClient(ctrl *gomock.Controller) *MockKazokuClient {
	mock := &MockKazokuClient{ctrl: ctrl}
	mock.recorder = &MockKazokuClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockKazokuClient) EXPECT() *MockKazokuClientMockRecorder {
	return m.recorder
}

// CreateAccount mocks base method
func (m *MockKazokuClient) CreateAccount(arg0 context.Context, arg1 *kazoku.Account, arg2 ...grpc.CallOption) (*kazoku.Account, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateAccount", varargs...)
	ret0, _ := ret[0].(*kazoku.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAccount indicates an expected call of CreateAccount
func (mr *MockKazokuClientMockRecorder) CreateAccount(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccount", reflect.TypeOf((*MockKazokuClient)(nil).CreateAccount), varargs...)
}

// CreateUserAccount mocks base method
func (m *MockKazokuClient) CreateUserAccount(arg0 context.Context, arg1 *kazoku.UserAccount, arg2 ...grpc.CallOption) (*kazoku.UserAccount, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateUserAccount", varargs...)
	ret0, _ := ret[0].(*kazoku.UserAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUserAccount indicates an expected call of CreateUserAccount
func (mr *MockKazokuClientMockRecorder) CreateUserAccount(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUserAccount", reflect.TypeOf((*MockKazokuClient)(nil).CreateUserAccount), varargs...)
}

// GetAccount mocks base method
func (m *MockKazokuClient) GetAccount(arg0 context.Context, arg1 *kazoku.Account, arg2 ...grpc.CallOption) (*kazoku.Account, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAccount", varargs...)
	ret0, _ := ret[0].(*kazoku.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccount indicates an expected call of GetAccount
func (mr *MockKazokuClientMockRecorder) GetAccount(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockKazokuClient)(nil).GetAccount), varargs...)
}

// GetUserAccounts mocks base method
func (m *MockKazokuClient) GetUserAccounts(arg0 context.Context, arg1 *kazoku.UserAccount, arg2 ...grpc.CallOption) (*kazoku.UserAccounts, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetUserAccounts", varargs...)
	ret0, _ := ret[0].(*kazoku.UserAccounts)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserAccounts indicates an expected call of GetUserAccounts
func (mr *MockKazokuClientMockRecorder) GetUserAccounts(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserAccounts", reflect.TypeOf((*MockKazokuClient)(nil).GetUserAccounts), varargs...)
}
