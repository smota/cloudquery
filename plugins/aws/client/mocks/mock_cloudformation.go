// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cq-provider-aws/client (interfaces: CloudFormationClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	cloudformation "github.com/aws/aws-sdk-go-v2/service/cloudformation"
	gomock "github.com/golang/mock/gomock"
)

// MockCloudFormationClient is a mock of CloudFormationClient interface.
type MockCloudFormationClient struct {
	ctrl     *gomock.Controller
	recorder *MockCloudFormationClientMockRecorder
}

// MockCloudFormationClientMockRecorder is the mock recorder for MockCloudFormationClient.
type MockCloudFormationClientMockRecorder struct {
	mock *MockCloudFormationClient
}

// NewMockCloudFormationClient creates a new mock instance.
func NewMockCloudFormationClient(ctrl *gomock.Controller) *MockCloudFormationClient {
	mock := &MockCloudFormationClient{ctrl: ctrl}
	mock.recorder = &MockCloudFormationClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCloudFormationClient) EXPECT() *MockCloudFormationClientMockRecorder {
	return m.recorder
}

// DescribeStacks mocks base method.
func (m *MockCloudFormationClient) DescribeStacks(arg0 context.Context, arg1 *cloudformation.DescribeStacksInput, arg2 ...func(*cloudformation.Options)) (*cloudformation.DescribeStacksOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeStacks", varargs...)
	ret0, _ := ret[0].(*cloudformation.DescribeStacksOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeStacks indicates an expected call of DescribeStacks.
func (mr *MockCloudFormationClientMockRecorder) DescribeStacks(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeStacks", reflect.TypeOf((*MockCloudFormationClient)(nil).DescribeStacks), varargs...)
}

// ListStackResources mocks base method.
func (m *MockCloudFormationClient) ListStackResources(arg0 context.Context, arg1 *cloudformation.ListStackResourcesInput, arg2 ...func(*cloudformation.Options)) (*cloudformation.ListStackResourcesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListStackResources", varargs...)
	ret0, _ := ret[0].(*cloudformation.ListStackResourcesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListStackResources indicates an expected call of ListStackResources.
func (mr *MockCloudFormationClientMockRecorder) ListStackResources(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListStackResources", reflect.TypeOf((*MockCloudFormationClient)(nil).ListStackResources), varargs...)
}
