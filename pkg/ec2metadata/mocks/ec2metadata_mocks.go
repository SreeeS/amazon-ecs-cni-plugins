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
// Source: github.com/aws/amazon-ecs-cni-plugins/pkg/ec2metadata (interfaces: EC2Metadata)

// Package mock_ec2metadata is a generated GoMock package.
package mock_ec2metadata

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockEC2Metadata is a mock of EC2Metadata interface
type MockEC2Metadata struct {
	ctrl     *gomock.Controller
	recorder *MockEC2MetadataMockRecorder
}

// MockEC2MetadataMockRecorder is the mock recorder for MockEC2Metadata
type MockEC2MetadataMockRecorder struct {
	mock *MockEC2Metadata
}

// NewMockEC2Metadata creates a new mock instance
func NewMockEC2Metadata(ctrl *gomock.Controller) *MockEC2Metadata {
	mock := &MockEC2Metadata{ctrl: ctrl}
	mock.recorder = &MockEC2MetadataMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEC2Metadata) EXPECT() *MockEC2MetadataMockRecorder {
	return m.recorder
}

// GetMetadata mocks base method
func (m *MockEC2Metadata) GetMetadata(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMetadata", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMetadata indicates an expected call of GetMetadata
func (mr *MockEC2MetadataMockRecorder) GetMetadata(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMetadata", reflect.TypeOf((*MockEC2Metadata)(nil).GetMetadata), arg0)
}
