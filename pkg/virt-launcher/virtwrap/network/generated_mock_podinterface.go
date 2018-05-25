// Automatically generated by MockGen. DO NOT EDIT!
// Source: podinterface.go

package network

import (
	gomock "github.com/golang/mock/gomock"

	api "kubevirt.io/kubevirt/pkg/virt-launcher/virtwrap/api"
)

// Mock of BindMechanism interface
type MockBindMechanism struct {
	ctrl     *gomock.Controller
	recorder *_MockBindMechanismRecorder
}

// Recorder for MockBindMechanism (not exported)
type _MockBindMechanismRecorder struct {
	mock *MockBindMechanism
}

func NewMockBindMechanism(ctrl *gomock.Controller) *MockBindMechanism {
	mock := &MockBindMechanism{ctrl: ctrl}
	mock.recorder = &_MockBindMechanismRecorder{mock}
	return mock
}

func (_m *MockBindMechanism) EXPECT() *_MockBindMechanismRecorder {
	return _m.recorder
}

func (_m *MockBindMechanism) discoverPodNetworkInterface() error {
	ret := _m.ctrl.Call(_m, "discoverPodNetworkInterface")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockBindMechanismRecorder) discoverPodNetworkInterface() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "discoverPodNetworkInterface")
}

func (_m *MockBindMechanism) preparePodNetworkInterfaces() error {
	ret := _m.ctrl.Call(_m, "preparePodNetworkInterfaces")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockBindMechanismRecorder) preparePodNetworkInterfaces() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "preparePodNetworkInterfaces")
}

func (_m *MockBindMechanism) startDHCPServer() {
	_m.ctrl.Call(_m, "startDHCPServer")
}

func (_mr *_MockBindMechanismRecorder) startDHCPServer() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "startDHCPServer")
}

func (_m *MockBindMechanism) decorateInterfaceConfig(ifconf *api.Interface) {
	_m.ctrl.Call(_m, "decorateInterfaceConfig", ifconf)
}

func (_mr *_MockBindMechanismRecorder) decorateInterfaceConfig(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "decorateInterfaceConfig", arg0)
}
