// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/issho-ni/issho/internal/pkg/mongo (interfaces: Client)

// Package mock_mongo is a generated GoMock package.
package mock_mongo

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	mongo "github.com/issho-ni/issho/internal/pkg/mongo"
	mongo0 "go.mongodb.org/mongo-driver/mongo"
	options "go.mongodb.org/mongo-driver/mongo/options"
	reflect "reflect"
)

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// Collection mocks base method
func (m *MockClient) Collection(arg0 string, arg1 ...*options.CollectionOptions) *mongo0.Collection {
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Collection", varargs...)
	ret0, _ := ret[0].(*mongo0.Collection)
	return ret0
}

// Collection indicates an expected call of Collection
func (mr *MockClientMockRecorder) Collection(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Collection", reflect.TypeOf((*MockClient)(nil).Collection), varargs...)
}

// Connect mocks base method
func (m *MockClient) Connect() context.CancelFunc {
	ret := m.ctrl.Call(m, "Connect")
	ret0, _ := ret[0].(context.CancelFunc)
	return ret0
}

// Connect indicates an expected call of Connect
func (mr *MockClientMockRecorder) Connect() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connect", reflect.TypeOf((*MockClient)(nil).Connect))
}

// Database mocks base method
func (m *MockClient) Database() *mongo0.Database {
	ret := m.ctrl.Call(m, "Database")
	ret0, _ := ret[0].(*mongo0.Database)
	return ret0
}

// Database indicates an expected call of Database
func (mr *MockClientMockRecorder) Database() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Database", reflect.TypeOf((*MockClient)(nil).Database))
}

// DefineIndexes mocks base method
func (m *MockClient) DefineIndexes(arg0 ...mongo.IndexSet) {
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "DefineIndexes", varargs...)
}

// DefineIndexes indicates an expected call of DefineIndexes
func (mr *MockClientMockRecorder) DefineIndexes(arg0 ...interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DefineIndexes", reflect.TypeOf((*MockClient)(nil).DefineIndexes), arg0...)
}