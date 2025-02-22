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
// Source: github.com/aws/amazon-ecs-cni-plugins/plugins/eni/engine (interfaces: Engine)

// Package mock_engine is a generated GoMock package.
package mock_engine

import (
	reflect "reflect"

	skel "github.com/containernetworking/cni/pkg/skel"
	gomock "github.com/golang/mock/gomock"
)

// MockEngine is a mock of Engine interface
type MockEngine struct {
	ctrl     *gomock.Controller
	recorder *MockEngineMockRecorder
}

// MockEngineMockRecorder is the mock recorder for MockEngine
type MockEngineMockRecorder struct {
	mock *MockEngine
}

// NewMockEngine creates a new mock instance
func NewMockEngine(ctrl *gomock.Controller) *MockEngine {
	mock := &MockEngine{ctrl: ctrl}
	mock.recorder = &MockEngineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEngine) EXPECT() *MockEngineMockRecorder {
	return m.recorder
}

// DoesMACAddressMapToIPV4Address mocks base method
func (m *MockEngine) DoesMACAddressMapToIPV4Address(arg0, arg1 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoesMACAddressMapToIPV4Address", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DoesMACAddressMapToIPV4Address indicates an expected call of DoesMACAddressMapToIPV4Address
func (mr *MockEngineMockRecorder) DoesMACAddressMapToIPV4Address(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoesMACAddressMapToIPV4Address", reflect.TypeOf((*MockEngine)(nil).DoesMACAddressMapToIPV4Address), arg0, arg1)
}

// DoesMACAddressMapToIPV6Address mocks base method
func (m *MockEngine) DoesMACAddressMapToIPV6Address(arg0, arg1 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoesMACAddressMapToIPV6Address", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DoesMACAddressMapToIPV6Address indicates an expected call of DoesMACAddressMapToIPV6Address
func (mr *MockEngineMockRecorder) DoesMACAddressMapToIPV6Address(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoesMACAddressMapToIPV6Address", reflect.TypeOf((*MockEngine)(nil).DoesMACAddressMapToIPV6Address), arg0, arg1)
}

// GetAllMACAddresses mocks base method
func (m *MockEngine) GetAllMACAddresses() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllMACAddresses")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllMACAddresses indicates an expected call of GetAllMACAddresses
func (mr *MockEngineMockRecorder) GetAllMACAddresses() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllMACAddresses", reflect.TypeOf((*MockEngine)(nil).GetAllMACAddresses))
}

// GetIPV4GatewayNetmask mocks base method
func (m *MockEngine) GetIPV4GatewayNetmask(arg0 string) (string, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIPV4GatewayNetmask", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetIPV4GatewayNetmask indicates an expected call of GetIPV4GatewayNetmask
func (mr *MockEngineMockRecorder) GetIPV4GatewayNetmask(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIPV4GatewayNetmask", reflect.TypeOf((*MockEngine)(nil).GetIPV4GatewayNetmask), arg0)
}

// GetIPV6Gateway mocks base method
func (m *MockEngine) GetIPV6Gateway(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIPV6Gateway", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIPV6Gateway indicates an expected call of GetIPV6Gateway
func (mr *MockEngineMockRecorder) GetIPV6Gateway(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIPV6Gateway", reflect.TypeOf((*MockEngine)(nil).GetIPV6Gateway), arg0)
}

// GetIPV6PrefixLength mocks base method
func (m *MockEngine) GetIPV6PrefixLength(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIPV6PrefixLength", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIPV6PrefixLength indicates an expected call of GetIPV6PrefixLength
func (mr *MockEngineMockRecorder) GetIPV6PrefixLength(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIPV6PrefixLength", reflect.TypeOf((*MockEngine)(nil).GetIPV6PrefixLength), arg0)
}

// GetInterfaceDeviceName mocks base method
func (m *MockEngine) GetInterfaceDeviceName(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInterfaceDeviceName", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInterfaceDeviceName indicates an expected call of GetInterfaceDeviceName
func (mr *MockEngineMockRecorder) GetInterfaceDeviceName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInterfaceDeviceName", reflect.TypeOf((*MockEngine)(nil).GetInterfaceDeviceName), arg0)
}

// GetMACAddressOfENI mocks base method
func (m *MockEngine) GetMACAddressOfENI(arg0 []string, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMACAddressOfENI", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMACAddressOfENI indicates an expected call of GetMACAddressOfENI
func (mr *MockEngineMockRecorder) GetMACAddressOfENI(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMACAddressOfENI", reflect.TypeOf((*MockEngine)(nil).GetMACAddressOfENI), arg0, arg1)
}

// SetupContainerNamespace mocks base method
func (m *MockEngine) SetupContainerNamespace(arg0 *skel.CmdArgs, arg1, arg2 string, arg3, arg4 []string, arg5, arg6 bool, arg7 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetupContainerNamespace", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetupContainerNamespace indicates an expected call of SetupContainerNamespace
func (mr *MockEngineMockRecorder) SetupContainerNamespace(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetupContainerNamespace", reflect.TypeOf((*MockEngine)(nil).SetupContainerNamespace), arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

// TeardownContainerNamespace mocks base method
func (m *MockEngine) TeardownContainerNamespace(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TeardownContainerNamespace", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// TeardownContainerNamespace indicates an expected call of TeardownContainerNamespace
func (mr *MockEngineMockRecorder) TeardownContainerNamespace(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TeardownContainerNamespace", reflect.TypeOf((*MockEngine)(nil).TeardownContainerNamespace), arg0, arg1)
}
