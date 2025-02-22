// Copyright 2023 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/amazon-ecs-cni-plugins/plugins/ipam/ipstore (interfaces: IPAllocator)

// Package mock_ipstore is a generated GoMock package.
package mock_ipstore

import (
	net "net"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIPAllocator is a mock of IPAllocator interface
type MockIPAllocator struct {
	ctrl     *gomock.Controller
	recorder *MockIPAllocatorMockRecorder
}

// MockIPAllocatorMockRecorder is the mock recorder for MockIPAllocator
type MockIPAllocatorMockRecorder struct {
	mock *MockIPAllocator
}

// NewMockIPAllocator creates a new mock instance
func NewMockIPAllocator(ctrl *gomock.Controller) *MockIPAllocator {
	mock := &MockIPAllocator{ctrl: ctrl}
	mock.recorder = &MockIPAllocatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIPAllocator) EXPECT() *MockIPAllocatorMockRecorder {
	return m.recorder
}

// Assign mocks base method
func (m *MockIPAllocator) Assign(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Assign", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Assign indicates an expected call of Assign
func (mr *MockIPAllocatorMockRecorder) Assign(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Assign", reflect.TypeOf((*MockIPAllocator)(nil).Assign), arg0, arg1)
}

// Close mocks base method
func (m *MockIPAllocator) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close
func (mr *MockIPAllocatorMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockIPAllocator)(nil).Close))
}

// Exists mocks base method
func (m *MockIPAllocator) Exists(arg0 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists
func (mr *MockIPAllocatorMockRecorder) Exists(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockIPAllocator)(nil).Exists), arg0)
}

// Get mocks base method
func (m *MockIPAllocator) Get(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockIPAllocatorMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockIPAllocator)(nil).Get), arg0)
}

// GetAvailableIP mocks base method
func (m *MockIPAllocator) GetAvailableIP(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAvailableIP", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAvailableIP indicates an expected call of GetAvailableIP
func (mr *MockIPAllocatorMockRecorder) GetAvailableIP(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAvailableIP", reflect.TypeOf((*MockIPAllocator)(nil).GetAvailableIP), arg0)
}

// Release mocks base method
func (m *MockIPAllocator) Release(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Release", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Release indicates an expected call of Release
func (mr *MockIPAllocatorMockRecorder) Release(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Release", reflect.TypeOf((*MockIPAllocator)(nil).Release), arg0)
}

// ReleaseByID mocks base method
func (m *MockIPAllocator) ReleaseByID(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReleaseByID", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReleaseByID indicates an expected call of ReleaseByID
func (mr *MockIPAllocatorMockRecorder) ReleaseByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReleaseByID", reflect.TypeOf((*MockIPAllocator)(nil).ReleaseByID), arg0)
}

// SetLastKnownIP mocks base method
func (m *MockIPAllocator) SetLastKnownIP(arg0 net.IP) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetLastKnownIP", arg0)
}

// SetLastKnownIP indicates an expected call of SetLastKnownIP
func (mr *MockIPAllocatorMockRecorder) SetLastKnownIP(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLastKnownIP", reflect.TypeOf((*MockIPAllocator)(nil).SetLastKnownIP), arg0)
}

// Update mocks base method
func (m *MockIPAllocator) Update(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockIPAllocatorMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIPAllocator)(nil).Update), arg0, arg1)
}
