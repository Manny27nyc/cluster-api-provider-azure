/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by MockGen. DO NOT EDIT.
// Source: ../client.go

// Package mock_privatedns is a generated GoMock package.
package mock_privatedns

import (
	context "context"
	reflect "reflect"

	privatedns "github.com/Azure/azure-sdk-for-go/services/privatedns/mgmt/2018-09-01/privatedns"
	gomock "github.com/golang/mock/gomock"
)

// Mockclient is a mock of client interface.
type Mockclient struct {
	ctrl     *gomock.Controller
	recorder *MockclientMockRecorder
}

// MockclientMockRecorder is the mock recorder for Mockclient.
type MockclientMockRecorder struct {
	mock *Mockclient
}

// NewMockclient creates a new mock instance.
func NewMockclient(ctrl *gomock.Controller) *Mockclient {
	mock := &Mockclient{ctrl: ctrl}
	mock.recorder = &MockclientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockclient) EXPECT() *MockclientMockRecorder {
	return m.recorder
}

// CreateOrUpdateLink mocks base method.
func (m *Mockclient) CreateOrUpdateLink(arg0 context.Context, arg1, arg2, arg3 string, arg4 privatedns.VirtualNetworkLink) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdateLink", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateOrUpdateLink indicates an expected call of CreateOrUpdateLink.
func (mr *MockclientMockRecorder) CreateOrUpdateLink(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdateLink", reflect.TypeOf((*Mockclient)(nil).CreateOrUpdateLink), arg0, arg1, arg2, arg3, arg4)
}

// CreateOrUpdateRecordSet mocks base method.
func (m *Mockclient) CreateOrUpdateRecordSet(arg0 context.Context, arg1, arg2 string, arg3 privatedns.RecordType, arg4 string, arg5 privatedns.RecordSet) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdateRecordSet", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateOrUpdateRecordSet indicates an expected call of CreateOrUpdateRecordSet.
func (mr *MockclientMockRecorder) CreateOrUpdateRecordSet(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdateRecordSet", reflect.TypeOf((*Mockclient)(nil).CreateOrUpdateRecordSet), arg0, arg1, arg2, arg3, arg4, arg5)
}

// CreateOrUpdateZone mocks base method.
func (m *Mockclient) CreateOrUpdateZone(arg0 context.Context, arg1, arg2 string, arg3 privatedns.PrivateZone) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdateZone", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateOrUpdateZone indicates an expected call of CreateOrUpdateZone.
func (mr *MockclientMockRecorder) CreateOrUpdateZone(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdateZone", reflect.TypeOf((*Mockclient)(nil).CreateOrUpdateZone), arg0, arg1, arg2, arg3)
}

// DeleteLink mocks base method.
func (m *Mockclient) DeleteLink(arg0 context.Context, arg1, arg2, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteLink", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteLink indicates an expected call of DeleteLink.
func (mr *MockclientMockRecorder) DeleteLink(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteLink", reflect.TypeOf((*Mockclient)(nil).DeleteLink), arg0, arg1, arg2, arg3)
}

// DeleteRecordSet mocks base method.
func (m *Mockclient) DeleteRecordSet(arg0 context.Context, arg1, arg2 string, arg3 privatedns.RecordType, arg4 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRecordSet", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRecordSet indicates an expected call of DeleteRecordSet.
func (mr *MockclientMockRecorder) DeleteRecordSet(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRecordSet", reflect.TypeOf((*Mockclient)(nil).DeleteRecordSet), arg0, arg1, arg2, arg3, arg4)
}

// DeleteZone mocks base method.
func (m *Mockclient) DeleteZone(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteZone", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteZone indicates an expected call of DeleteZone.
func (mr *MockclientMockRecorder) DeleteZone(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteZone", reflect.TypeOf((*Mockclient)(nil).DeleteZone), arg0, arg1, arg2)
}

// GetLink mocks base method.
func (m *Mockclient) GetLink(arg0 context.Context, arg1, arg2, arg3 string) (privatedns.VirtualNetworkLink, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLink", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(privatedns.VirtualNetworkLink)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLink indicates an expected call of GetLink.
func (mr *MockclientMockRecorder) GetLink(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLink", reflect.TypeOf((*Mockclient)(nil).GetLink), arg0, arg1, arg2, arg3)
}

// GetZone mocks base method.
func (m *Mockclient) GetZone(arg0 context.Context, arg1, arg2 string) (privatedns.PrivateZone, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetZone", arg0, arg1, arg2)
	ret0, _ := ret[0].(privatedns.PrivateZone)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetZone indicates an expected call of GetZone.
func (mr *MockclientMockRecorder) GetZone(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetZone", reflect.TypeOf((*Mockclient)(nil).GetZone), arg0, arg1, arg2)
}
