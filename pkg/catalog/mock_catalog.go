// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/openservicemesh/osm/pkg/catalog (interfaces: MeshCataloger)

// Package catalog is a generated GoMock package.
package catalog

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	certificate "github.com/openservicemesh/osm/pkg/certificate"
	endpoint "github.com/openservicemesh/osm/pkg/endpoint"
	envoy "github.com/openservicemesh/osm/pkg/envoy"
	service "github.com/openservicemesh/osm/pkg/service"
	smi "github.com/openservicemesh/osm/pkg/smi"
	trafficpolicy "github.com/openservicemesh/osm/pkg/trafficpolicy"
	v1alpha2 "github.com/servicemeshinterface/smi-sdk-go/pkg/apis/access/v1alpha2"
	v1alpha3 "github.com/servicemeshinterface/smi-sdk-go/pkg/apis/specs/v1alpha3"
	v1alpha20 "github.com/servicemeshinterface/smi-sdk-go/pkg/apis/split/v1alpha2"
)

// MockMeshCataloger is a mock of MeshCataloger interface
type MockMeshCataloger struct {
	ctrl     *gomock.Controller
	recorder *MockMeshCatalogerMockRecorder
}

// MockMeshCatalogerMockRecorder is the mock recorder for MockMeshCataloger
type MockMeshCatalogerMockRecorder struct {
	mock *MockMeshCataloger
}

// NewMockMeshCataloger creates a new mock instance
func NewMockMeshCataloger(ctrl *gomock.Controller) *MockMeshCataloger {
	mock := &MockMeshCataloger{ctrl: ctrl}
	mock.recorder = &MockMeshCatalogerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMeshCataloger) EXPECT() *MockMeshCatalogerMockRecorder {
	return m.recorder
}

// ExpectProxy mocks base method
func (m *MockMeshCataloger) ExpectProxy(arg0 certificate.CommonName) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ExpectProxy", arg0)
}

// ExpectProxy indicates an expected call of ExpectProxy
func (mr *MockMeshCatalogerMockRecorder) ExpectProxy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExpectProxy", reflect.TypeOf((*MockMeshCataloger)(nil).ExpectProxy), arg0)
}

// GetIngressRoutesPerHost mocks base method
func (m *MockMeshCataloger) GetIngressRoutesPerHost(arg0 service.MeshService) (map[string][]trafficpolicy.HTTPRoute, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIngressRoutesPerHost", arg0)
	ret0, _ := ret[0].(map[string][]trafficpolicy.HTTPRoute)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIngressRoutesPerHost indicates an expected call of GetIngressRoutesPerHost
func (mr *MockMeshCatalogerMockRecorder) GetIngressRoutesPerHost(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIngressRoutesPerHost", reflect.TypeOf((*MockMeshCataloger)(nil).GetIngressRoutesPerHost), arg0)
}

// GetResolvableHostnamesForUpstreamService mocks base method
func (m *MockMeshCataloger) GetResolvableHostnamesForUpstreamService(arg0, arg1 service.MeshService) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResolvableHostnamesForUpstreamService", arg0, arg1)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetResolvableHostnamesForUpstreamService indicates an expected call of GetResolvableHostnamesForUpstreamService
func (mr *MockMeshCatalogerMockRecorder) GetResolvableHostnamesForUpstreamService(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResolvableHostnamesForUpstreamService", reflect.TypeOf((*MockMeshCataloger)(nil).GetResolvableHostnamesForUpstreamService), arg0, arg1)
}

// GetResolvableServiceEndpoints mocks base method
func (m *MockMeshCataloger) GetResolvableServiceEndpoints(arg0 service.MeshService) ([]endpoint.Endpoint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResolvableServiceEndpoints", arg0)
	ret0, _ := ret[0].([]endpoint.Endpoint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetResolvableServiceEndpoints indicates an expected call of GetResolvableServiceEndpoints
func (mr *MockMeshCatalogerMockRecorder) GetResolvableServiceEndpoints(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResolvableServiceEndpoints", reflect.TypeOf((*MockMeshCataloger)(nil).GetResolvableServiceEndpoints), arg0)
}

// GetSMISpec mocks base method
func (m *MockMeshCataloger) GetSMISpec() smi.MeshSpec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSMISpec")
	ret0, _ := ret[0].(smi.MeshSpec)
	return ret0
}

// GetSMISpec indicates an expected call of GetSMISpec
func (mr *MockMeshCatalogerMockRecorder) GetSMISpec() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSMISpec", reflect.TypeOf((*MockMeshCataloger)(nil).GetSMISpec))
}

// GetServicesForServiceAccount mocks base method
func (m *MockMeshCataloger) GetServicesForServiceAccount(arg0 service.K8sServiceAccount) ([]service.MeshService, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServicesForServiceAccount", arg0)
	ret0, _ := ret[0].([]service.MeshService)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetServicesForServiceAccount indicates an expected call of GetServicesForServiceAccount
func (mr *MockMeshCatalogerMockRecorder) GetServicesForServiceAccount(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServicesForServiceAccount", reflect.TypeOf((*MockMeshCataloger)(nil).GetServicesForServiceAccount), arg0)
}

// GetServicesFromEnvoyCertificate mocks base method
func (m *MockMeshCataloger) GetServicesFromEnvoyCertificate(arg0 certificate.CommonName) ([]service.MeshService, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServicesFromEnvoyCertificate", arg0)
	ret0, _ := ret[0].([]service.MeshService)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetServicesFromEnvoyCertificate indicates an expected call of GetServicesFromEnvoyCertificate
func (mr *MockMeshCatalogerMockRecorder) GetServicesFromEnvoyCertificate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServicesFromEnvoyCertificate", reflect.TypeOf((*MockMeshCataloger)(nil).GetServicesFromEnvoyCertificate), arg0)
}

// GetWeightedClusterForService mocks base method
func (m *MockMeshCataloger) GetWeightedClusterForService(arg0 service.MeshService) (service.WeightedCluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWeightedClusterForService", arg0)
	ret0, _ := ret[0].(service.WeightedCluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWeightedClusterForService indicates an expected call of GetWeightedClusterForService
func (mr *MockMeshCatalogerMockRecorder) GetWeightedClusterForService(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWeightedClusterForService", reflect.TypeOf((*MockMeshCataloger)(nil).GetWeightedClusterForService), arg0)
}

// ListAllowedInboundServiceAccounts mocks base method
func (m *MockMeshCataloger) ListAllowedInboundServiceAccounts(arg0 service.K8sServiceAccount) ([]service.K8sServiceAccount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllowedInboundServiceAccounts", arg0)
	ret0, _ := ret[0].([]service.K8sServiceAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllowedInboundServiceAccounts indicates an expected call of ListAllowedInboundServiceAccounts
func (mr *MockMeshCatalogerMockRecorder) ListAllowedInboundServiceAccounts(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllowedInboundServiceAccounts", reflect.TypeOf((*MockMeshCataloger)(nil).ListAllowedInboundServiceAccounts), arg0)
}

// ListAllowedInboundServices mocks base method
func (m *MockMeshCataloger) ListAllowedInboundServices(arg0 service.MeshService) ([]service.MeshService, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllowedInboundServices", arg0)
	ret0, _ := ret[0].([]service.MeshService)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllowedInboundServices indicates an expected call of ListAllowedInboundServices
func (mr *MockMeshCatalogerMockRecorder) ListAllowedInboundServices(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllowedInboundServices", reflect.TypeOf((*MockMeshCataloger)(nil).ListAllowedInboundServices), arg0)
}

// ListAllowedOutboundServiceAccounts mocks base method
func (m *MockMeshCataloger) ListAllowedOutboundServiceAccounts(arg0 service.K8sServiceAccount) ([]service.K8sServiceAccount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllowedOutboundServiceAccounts", arg0)
	ret0, _ := ret[0].([]service.K8sServiceAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllowedOutboundServiceAccounts indicates an expected call of ListAllowedOutboundServiceAccounts
func (mr *MockMeshCatalogerMockRecorder) ListAllowedOutboundServiceAccounts(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllowedOutboundServiceAccounts", reflect.TypeOf((*MockMeshCataloger)(nil).ListAllowedOutboundServiceAccounts), arg0)
}

// ListAllowedOutboundServices mocks base method
func (m *MockMeshCataloger) ListAllowedOutboundServices(arg0 service.MeshService) ([]service.MeshService, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllowedOutboundServices", arg0)
	ret0, _ := ret[0].([]service.MeshService)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllowedOutboundServices indicates an expected call of ListAllowedOutboundServices
func (mr *MockMeshCatalogerMockRecorder) ListAllowedOutboundServices(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllowedOutboundServices", reflect.TypeOf((*MockMeshCataloger)(nil).ListAllowedOutboundServices), arg0)
}

// ListEndpointsForService mocks base method
func (m *MockMeshCataloger) ListEndpointsForService(arg0 service.MeshService) ([]endpoint.Endpoint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListEndpointsForService", arg0)
	ret0, _ := ret[0].([]endpoint.Endpoint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListEndpointsForService indicates an expected call of ListEndpointsForService
func (mr *MockMeshCatalogerMockRecorder) ListEndpointsForService(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEndpointsForService", reflect.TypeOf((*MockMeshCataloger)(nil).ListEndpointsForService), arg0)
}

// ListMonitoredNamespaces mocks base method
func (m *MockMeshCataloger) ListMonitoredNamespaces() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListMonitoredNamespaces")
	ret0, _ := ret[0].([]string)
	return ret0
}

// ListMonitoredNamespaces indicates an expected call of ListMonitoredNamespaces
func (mr *MockMeshCatalogerMockRecorder) ListMonitoredNamespaces() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMonitoredNamespaces", reflect.TypeOf((*MockMeshCataloger)(nil).ListMonitoredNamespaces))
}

// ListSMIPolicies mocks base method
func (m *MockMeshCataloger) ListSMIPolicies() ([]*v1alpha20.TrafficSplit, []service.WeightedService, []service.K8sServiceAccount, []*v1alpha3.HTTPRouteGroup, []*v1alpha2.TrafficTarget) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSMIPolicies")
	ret0, _ := ret[0].([]*v1alpha20.TrafficSplit)
	ret1, _ := ret[1].([]service.WeightedService)
	ret2, _ := ret[2].([]service.K8sServiceAccount)
	ret3, _ := ret[3].([]*v1alpha3.HTTPRouteGroup)
	ret4, _ := ret[4].([]*v1alpha2.TrafficTarget)
	return ret0, ret1, ret2, ret3, ret4
}

// ListSMIPolicies indicates an expected call of ListSMIPolicies
func (mr *MockMeshCatalogerMockRecorder) ListSMIPolicies() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSMIPolicies", reflect.TypeOf((*MockMeshCataloger)(nil).ListSMIPolicies))
}

// ListTrafficPolicies mocks base method
func (m *MockMeshCataloger) ListTrafficPolicies(arg0 service.MeshService) ([]trafficpolicy.TrafficTarget, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTrafficPolicies", arg0)
	ret0, _ := ret[0].([]trafficpolicy.TrafficTarget)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTrafficPolicies indicates an expected call of ListTrafficPolicies
func (mr *MockMeshCatalogerMockRecorder) ListTrafficPolicies(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTrafficPolicies", reflect.TypeOf((*MockMeshCataloger)(nil).ListTrafficPolicies), arg0)
}

// RegisterProxy mocks base method
func (m *MockMeshCataloger) RegisterProxy(arg0 *envoy.Proxy) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterProxy", arg0)
}

// RegisterProxy indicates an expected call of RegisterProxy
func (mr *MockMeshCatalogerMockRecorder) RegisterProxy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterProxy", reflect.TypeOf((*MockMeshCataloger)(nil).RegisterProxy), arg0)
}

// UnregisterProxy mocks base method
func (m *MockMeshCataloger) UnregisterProxy(arg0 *envoy.Proxy) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UnregisterProxy", arg0)
}

// UnregisterProxy indicates an expected call of UnregisterProxy
func (mr *MockMeshCatalogerMockRecorder) UnregisterProxy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnregisterProxy", reflect.TypeOf((*MockMeshCataloger)(nil).UnregisterProxy), arg0)
}
