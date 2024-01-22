// Code generated by MockGen. DO NOT EDIT.
// Source: k8s.io/client-go/kubernetes/typed/admissionregistration/v1beta1 (interfaces: AdmissionregistrationV1beta1Interface)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1beta1 "k8s.io/client-go/kubernetes/typed/admissionregistration/v1beta1"
	rest "k8s.io/client-go/rest"
)

// MockAdmissionregistrationV1beta1Interface is a mock of AdmissionregistrationV1beta1Interface interface.
type MockAdmissionregistrationV1beta1Interface struct {
	ctrl     *gomock.Controller
	recorder *MockAdmissionregistrationV1beta1InterfaceMockRecorder
}

// MockAdmissionregistrationV1beta1InterfaceMockRecorder is the mock recorder for MockAdmissionregistrationV1beta1Interface.
type MockAdmissionregistrationV1beta1InterfaceMockRecorder struct {
	mock *MockAdmissionregistrationV1beta1Interface
}

// NewMockAdmissionregistrationV1beta1Interface creates a new mock instance.
func NewMockAdmissionregistrationV1beta1Interface(ctrl *gomock.Controller) *MockAdmissionregistrationV1beta1Interface {
	mock := &MockAdmissionregistrationV1beta1Interface{ctrl: ctrl}
	mock.recorder = &MockAdmissionregistrationV1beta1InterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAdmissionregistrationV1beta1Interface) EXPECT() *MockAdmissionregistrationV1beta1InterfaceMockRecorder {
	return m.recorder
}

// MutatingWebhookConfigurations mocks base method.
func (m *MockAdmissionregistrationV1beta1Interface) MutatingWebhookConfigurations() v1beta1.MutatingWebhookConfigurationInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MutatingWebhookConfigurations")
	ret0, _ := ret[0].(v1beta1.MutatingWebhookConfigurationInterface)
	return ret0
}

// MutatingWebhookConfigurations indicates an expected call of MutatingWebhookConfigurations.
func (mr *MockAdmissionregistrationV1beta1InterfaceMockRecorder) MutatingWebhookConfigurations() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MutatingWebhookConfigurations", reflect.TypeOf((*MockAdmissionregistrationV1beta1Interface)(nil).MutatingWebhookConfigurations))
}

// RESTClient mocks base method.
func (m *MockAdmissionregistrationV1beta1Interface) RESTClient() rest.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RESTClient")
	ret0, _ := ret[0].(rest.Interface)
	return ret0
}

// RESTClient indicates an expected call of RESTClient.
func (mr *MockAdmissionregistrationV1beta1InterfaceMockRecorder) RESTClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RESTClient", reflect.TypeOf((*MockAdmissionregistrationV1beta1Interface)(nil).RESTClient))
}

// ValidatingAdmissionPolicies mocks base method.
func (m *MockAdmissionregistrationV1beta1Interface) ValidatingAdmissionPolicies() v1beta1.ValidatingAdmissionPolicyInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidatingAdmissionPolicies")
	ret0, _ := ret[0].(v1beta1.ValidatingAdmissionPolicyInterface)
	return ret0
}

// ValidatingAdmissionPolicies indicates an expected call of ValidatingAdmissionPolicies.
func (mr *MockAdmissionregistrationV1beta1InterfaceMockRecorder) ValidatingAdmissionPolicies() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidatingAdmissionPolicies", reflect.TypeOf((*MockAdmissionregistrationV1beta1Interface)(nil).ValidatingAdmissionPolicies))
}

// ValidatingAdmissionPolicyBindings mocks base method.
func (m *MockAdmissionregistrationV1beta1Interface) ValidatingAdmissionPolicyBindings() v1beta1.ValidatingAdmissionPolicyBindingInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidatingAdmissionPolicyBindings")
	ret0, _ := ret[0].(v1beta1.ValidatingAdmissionPolicyBindingInterface)
	return ret0
}

// ValidatingAdmissionPolicyBindings indicates an expected call of ValidatingAdmissionPolicyBindings.
func (mr *MockAdmissionregistrationV1beta1InterfaceMockRecorder) ValidatingAdmissionPolicyBindings() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidatingAdmissionPolicyBindings", reflect.TypeOf((*MockAdmissionregistrationV1beta1Interface)(nil).ValidatingAdmissionPolicyBindings))
}

// ValidatingWebhookConfigurations mocks base method.
func (m *MockAdmissionregistrationV1beta1Interface) ValidatingWebhookConfigurations() v1beta1.ValidatingWebhookConfigurationInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidatingWebhookConfigurations")
	ret0, _ := ret[0].(v1beta1.ValidatingWebhookConfigurationInterface)
	return ret0
}

// ValidatingWebhookConfigurations indicates an expected call of ValidatingWebhookConfigurations.
func (mr *MockAdmissionregistrationV1beta1InterfaceMockRecorder) ValidatingWebhookConfigurations() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidatingWebhookConfigurations", reflect.TypeOf((*MockAdmissionregistrationV1beta1Interface)(nil).ValidatingWebhookConfigurations))
}
