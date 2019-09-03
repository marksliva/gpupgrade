// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/greenplum-db/gpupgrade/idl (interfaces: CliToHubClient,CliToHubServer,CliToHub_UpgradeConvertMasterServer)

// Package mock_idl is a generated GoMock package.
package mock_idl

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	idl "github.com/greenplum-db/gpupgrade/idl"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
	reflect "reflect"
)

// MockCliToHubClient is a mock of CliToHubClient interface
type MockCliToHubClient struct {
	ctrl     *gomock.Controller
	recorder *MockCliToHubClientMockRecorder
}

// MockCliToHubClientMockRecorder is the mock recorder for MockCliToHubClient
type MockCliToHubClientMockRecorder struct {
	mock *MockCliToHubClient
}

// NewMockCliToHubClient creates a new mock instance
func NewMockCliToHubClient(ctrl *gomock.Controller) *MockCliToHubClient {
	mock := &MockCliToHubClient{ctrl: ctrl}
	mock.recorder = &MockCliToHubClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCliToHubClient) EXPECT() *MockCliToHubClientMockRecorder {
	return m.recorder
}

// CheckConfig mocks base method
func (m *MockCliToHubClient) CheckConfig(arg0 context.Context, arg1 *idl.CheckConfigRequest, arg2 ...grpc.CallOption) (*idl.CheckConfigReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CheckConfig", varargs...)
	ret0, _ := ret[0].(*idl.CheckConfigReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckConfig indicates an expected call of CheckConfig
func (mr *MockCliToHubClientMockRecorder) CheckConfig(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckConfig", reflect.TypeOf((*MockCliToHubClient)(nil).CheckConfig), varargs...)
}

// CheckDiskSpace mocks base method
func (m *MockCliToHubClient) CheckDiskSpace(arg0 context.Context, arg1 *idl.CheckDiskSpaceRequest, arg2 ...grpc.CallOption) (*idl.CheckDiskSpaceReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CheckDiskSpace", varargs...)
	ret0, _ := ret[0].(*idl.CheckDiskSpaceReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckDiskSpace indicates an expected call of CheckDiskSpace
func (mr *MockCliToHubClientMockRecorder) CheckDiskSpace(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckDiskSpace", reflect.TypeOf((*MockCliToHubClient)(nil).CheckDiskSpace), varargs...)
}

// CheckObjectCount mocks base method
func (m *MockCliToHubClient) CheckObjectCount(arg0 context.Context, arg1 *idl.CheckObjectCountRequest, arg2 ...grpc.CallOption) (*idl.CheckObjectCountReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CheckObjectCount", varargs...)
	ret0, _ := ret[0].(*idl.CheckObjectCountReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckObjectCount indicates an expected call of CheckObjectCount
func (mr *MockCliToHubClientMockRecorder) CheckObjectCount(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckObjectCount", reflect.TypeOf((*MockCliToHubClient)(nil).CheckObjectCount), varargs...)
}

// CheckSeginstall mocks base method
func (m *MockCliToHubClient) CheckSeginstall(arg0 context.Context, arg1 *idl.CheckSeginstallRequest, arg2 ...grpc.CallOption) (*idl.CheckSeginstallReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CheckSeginstall", varargs...)
	ret0, _ := ret[0].(*idl.CheckSeginstallReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckSeginstall indicates an expected call of CheckSeginstall
func (mr *MockCliToHubClientMockRecorder) CheckSeginstall(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckSeginstall", reflect.TypeOf((*MockCliToHubClient)(nil).CheckSeginstall), varargs...)
}

// CheckVersion mocks base method
func (m *MockCliToHubClient) CheckVersion(arg0 context.Context, arg1 *idl.CheckVersionRequest, arg2 ...grpc.CallOption) (*idl.CheckVersionReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CheckVersion", varargs...)
	ret0, _ := ret[0].(*idl.CheckVersionReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckVersion indicates an expected call of CheckVersion
func (mr *MockCliToHubClientMockRecorder) CheckVersion(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckVersion", reflect.TypeOf((*MockCliToHubClient)(nil).CheckVersion), varargs...)
}

// GetConfig mocks base method
func (m *MockCliToHubClient) GetConfig(arg0 context.Context, arg1 *idl.GetConfigRequest, arg2 ...grpc.CallOption) (*idl.GetConfigReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetConfig", varargs...)
	ret0, _ := ret[0].(*idl.GetConfigReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConfig indicates an expected call of GetConfig
func (mr *MockCliToHubClientMockRecorder) GetConfig(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfig", reflect.TypeOf((*MockCliToHubClient)(nil).GetConfig), varargs...)
}

// Ping mocks base method
func (m *MockCliToHubClient) Ping(arg0 context.Context, arg1 *idl.PingRequest, arg2 ...grpc.CallOption) (*idl.PingReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Ping", varargs...)
	ret0, _ := ret[0].(*idl.PingReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Ping indicates an expected call of Ping
func (mr *MockCliToHubClientMockRecorder) Ping(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockCliToHubClient)(nil).Ping), varargs...)
}

// PrepareInitCluster mocks base method
func (m *MockCliToHubClient) PrepareInitCluster(arg0 context.Context, arg1 *idl.PrepareInitClusterRequest, arg2 ...grpc.CallOption) (*idl.PrepareInitClusterReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PrepareInitCluster", varargs...)
	ret0, _ := ret[0].(*idl.PrepareInitClusterReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrepareInitCluster indicates an expected call of PrepareInitCluster
func (mr *MockCliToHubClientMockRecorder) PrepareInitCluster(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareInitCluster", reflect.TypeOf((*MockCliToHubClient)(nil).PrepareInitCluster), varargs...)
}

// PrepareShutdownClusters mocks base method
func (m *MockCliToHubClient) PrepareShutdownClusters(arg0 context.Context, arg1 *idl.PrepareShutdownClustersRequest, arg2 ...grpc.CallOption) (*idl.PrepareShutdownClustersReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PrepareShutdownClusters", varargs...)
	ret0, _ := ret[0].(*idl.PrepareShutdownClustersReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrepareShutdownClusters indicates an expected call of PrepareShutdownClusters
func (mr *MockCliToHubClientMockRecorder) PrepareShutdownClusters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareShutdownClusters", reflect.TypeOf((*MockCliToHubClient)(nil).PrepareShutdownClusters), varargs...)
}

// PrepareStartAgents mocks base method
func (m *MockCliToHubClient) PrepareStartAgents(arg0 context.Context, arg1 *idl.PrepareStartAgentsRequest, arg2 ...grpc.CallOption) (*idl.PrepareStartAgentsReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PrepareStartAgents", varargs...)
	ret0, _ := ret[0].(*idl.PrepareStartAgentsReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrepareStartAgents indicates an expected call of PrepareStartAgents
func (mr *MockCliToHubClientMockRecorder) PrepareStartAgents(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareStartAgents", reflect.TypeOf((*MockCliToHubClient)(nil).PrepareStartAgents), varargs...)
}

// SetConfig mocks base method
func (m *MockCliToHubClient) SetConfig(arg0 context.Context, arg1 *idl.SetConfigRequest, arg2 ...grpc.CallOption) (*idl.SetConfigReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SetConfig", varargs...)
	ret0, _ := ret[0].(*idl.SetConfigReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetConfig indicates an expected call of SetConfig
func (mr *MockCliToHubClientMockRecorder) SetConfig(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetConfig", reflect.TypeOf((*MockCliToHubClient)(nil).SetConfig), varargs...)
}

// StatusConversion mocks base method
func (m *MockCliToHubClient) StatusConversion(arg0 context.Context, arg1 *idl.StatusConversionRequest, arg2 ...grpc.CallOption) (*idl.StatusConversionReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StatusConversion", varargs...)
	ret0, _ := ret[0].(*idl.StatusConversionReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StatusConversion indicates an expected call of StatusConversion
func (mr *MockCliToHubClientMockRecorder) StatusConversion(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StatusConversion", reflect.TypeOf((*MockCliToHubClient)(nil).StatusConversion), varargs...)
}

// StatusUpgrade mocks base method
func (m *MockCliToHubClient) StatusUpgrade(arg0 context.Context, arg1 *idl.StatusUpgradeRequest, arg2 ...grpc.CallOption) (*idl.StatusUpgradeReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StatusUpgrade", varargs...)
	ret0, _ := ret[0].(*idl.StatusUpgradeReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StatusUpgrade indicates an expected call of StatusUpgrade
func (mr *MockCliToHubClientMockRecorder) StatusUpgrade(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StatusUpgrade", reflect.TypeOf((*MockCliToHubClient)(nil).StatusUpgrade), varargs...)
}

// UpgradeConvertMaster mocks base method
func (m *MockCliToHubClient) UpgradeConvertMaster(arg0 context.Context, arg1 *idl.UpgradeConvertMasterRequest, arg2 ...grpc.CallOption) (idl.CliToHub_UpgradeConvertMasterClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpgradeConvertMaster", varargs...)
	ret0, _ := ret[0].(idl.CliToHub_UpgradeConvertMasterClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpgradeConvertMaster indicates an expected call of UpgradeConvertMaster
func (mr *MockCliToHubClientMockRecorder) UpgradeConvertMaster(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpgradeConvertMaster", reflect.TypeOf((*MockCliToHubClient)(nil).UpgradeConvertMaster), varargs...)
}

// UpgradeConvertPrimaries mocks base method
func (m *MockCliToHubClient) UpgradeConvertPrimaries(arg0 context.Context, arg1 *idl.UpgradeConvertPrimariesRequest, arg2 ...grpc.CallOption) (*idl.UpgradeConvertPrimariesReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpgradeConvertPrimaries", varargs...)
	ret0, _ := ret[0].(*idl.UpgradeConvertPrimariesReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpgradeConvertPrimaries indicates an expected call of UpgradeConvertPrimaries
func (mr *MockCliToHubClientMockRecorder) UpgradeConvertPrimaries(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpgradeConvertPrimaries", reflect.TypeOf((*MockCliToHubClient)(nil).UpgradeConvertPrimaries), varargs...)
}

// UpgradeCopyMasterDataDir mocks base method
func (m *MockCliToHubClient) UpgradeCopyMasterDataDir(arg0 context.Context, arg1 *idl.UpgradeCopyMasterDataDirRequest, arg2 ...grpc.CallOption) (*idl.UpgradeCopyMasterDataDirReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpgradeCopyMasterDataDir", varargs...)
	ret0, _ := ret[0].(*idl.UpgradeCopyMasterDataDirReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpgradeCopyMasterDataDir indicates an expected call of UpgradeCopyMasterDataDir
func (mr *MockCliToHubClientMockRecorder) UpgradeCopyMasterDataDir(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpgradeCopyMasterDataDir", reflect.TypeOf((*MockCliToHubClient)(nil).UpgradeCopyMasterDataDir), varargs...)
}

// UpgradeReconfigurePorts mocks base method
func (m *MockCliToHubClient) UpgradeReconfigurePorts(arg0 context.Context, arg1 *idl.UpgradeReconfigurePortsRequest, arg2 ...grpc.CallOption) (*idl.UpgradeReconfigurePortsReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpgradeReconfigurePorts", varargs...)
	ret0, _ := ret[0].(*idl.UpgradeReconfigurePortsReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpgradeReconfigurePorts indicates an expected call of UpgradeReconfigurePorts
func (mr *MockCliToHubClientMockRecorder) UpgradeReconfigurePorts(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpgradeReconfigurePorts", reflect.TypeOf((*MockCliToHubClient)(nil).UpgradeReconfigurePorts), varargs...)
}

// UpgradeValidateStartCluster mocks base method
func (m *MockCliToHubClient) UpgradeValidateStartCluster(arg0 context.Context, arg1 *idl.UpgradeValidateStartClusterRequest, arg2 ...grpc.CallOption) (*idl.UpgradeValidateStartClusterReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpgradeValidateStartCluster", varargs...)
	ret0, _ := ret[0].(*idl.UpgradeValidateStartClusterReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpgradeValidateStartCluster indicates an expected call of UpgradeValidateStartCluster
func (mr *MockCliToHubClientMockRecorder) UpgradeValidateStartCluster(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpgradeValidateStartCluster", reflect.TypeOf((*MockCliToHubClient)(nil).UpgradeValidateStartCluster), varargs...)
}

// MockCliToHubServer is a mock of CliToHubServer interface
type MockCliToHubServer struct {
	ctrl     *gomock.Controller
	recorder *MockCliToHubServerMockRecorder
}

// MockCliToHubServerMockRecorder is the mock recorder for MockCliToHubServer
type MockCliToHubServerMockRecorder struct {
	mock *MockCliToHubServer
}

// NewMockCliToHubServer creates a new mock instance
func NewMockCliToHubServer(ctrl *gomock.Controller) *MockCliToHubServer {
	mock := &MockCliToHubServer{ctrl: ctrl}
	mock.recorder = &MockCliToHubServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCliToHubServer) EXPECT() *MockCliToHubServerMockRecorder {
	return m.recorder
}

// CheckConfig mocks base method
func (m *MockCliToHubServer) CheckConfig(arg0 context.Context, arg1 *idl.CheckConfigRequest) (*idl.CheckConfigReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckConfig", arg0, arg1)
	ret0, _ := ret[0].(*idl.CheckConfigReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckConfig indicates an expected call of CheckConfig
func (mr *MockCliToHubServerMockRecorder) CheckConfig(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckConfig", reflect.TypeOf((*MockCliToHubServer)(nil).CheckConfig), arg0, arg1)
}

// CheckDiskSpace mocks base method
func (m *MockCliToHubServer) CheckDiskSpace(arg0 context.Context, arg1 *idl.CheckDiskSpaceRequest) (*idl.CheckDiskSpaceReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckDiskSpace", arg0, arg1)
	ret0, _ := ret[0].(*idl.CheckDiskSpaceReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckDiskSpace indicates an expected call of CheckDiskSpace
func (mr *MockCliToHubServerMockRecorder) CheckDiskSpace(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckDiskSpace", reflect.TypeOf((*MockCliToHubServer)(nil).CheckDiskSpace), arg0, arg1)
}

// CheckObjectCount mocks base method
func (m *MockCliToHubServer) CheckObjectCount(arg0 context.Context, arg1 *idl.CheckObjectCountRequest) (*idl.CheckObjectCountReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckObjectCount", arg0, arg1)
	ret0, _ := ret[0].(*idl.CheckObjectCountReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckObjectCount indicates an expected call of CheckObjectCount
func (mr *MockCliToHubServerMockRecorder) CheckObjectCount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckObjectCount", reflect.TypeOf((*MockCliToHubServer)(nil).CheckObjectCount), arg0, arg1)
}

// CheckSeginstall mocks base method
func (m *MockCliToHubServer) CheckSeginstall(arg0 context.Context, arg1 *idl.CheckSeginstallRequest) (*idl.CheckSeginstallReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckSeginstall", arg0, arg1)
	ret0, _ := ret[0].(*idl.CheckSeginstallReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckSeginstall indicates an expected call of CheckSeginstall
func (mr *MockCliToHubServerMockRecorder) CheckSeginstall(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckSeginstall", reflect.TypeOf((*MockCliToHubServer)(nil).CheckSeginstall), arg0, arg1)
}

// CheckVersion mocks base method
func (m *MockCliToHubServer) CheckVersion(arg0 context.Context, arg1 *idl.CheckVersionRequest) (*idl.CheckVersionReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckVersion", arg0, arg1)
	ret0, _ := ret[0].(*idl.CheckVersionReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckVersion indicates an expected call of CheckVersion
func (mr *MockCliToHubServerMockRecorder) CheckVersion(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckVersion", reflect.TypeOf((*MockCliToHubServer)(nil).CheckVersion), arg0, arg1)
}

// GetConfig mocks base method
func (m *MockCliToHubServer) GetConfig(arg0 context.Context, arg1 *idl.GetConfigRequest) (*idl.GetConfigReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConfig", arg0, arg1)
	ret0, _ := ret[0].(*idl.GetConfigReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConfig indicates an expected call of GetConfig
func (mr *MockCliToHubServerMockRecorder) GetConfig(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfig", reflect.TypeOf((*MockCliToHubServer)(nil).GetConfig), arg0, arg1)
}

// Ping mocks base method
func (m *MockCliToHubServer) Ping(arg0 context.Context, arg1 *idl.PingRequest) (*idl.PingReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ping", arg0, arg1)
	ret0, _ := ret[0].(*idl.PingReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Ping indicates an expected call of Ping
func (mr *MockCliToHubServerMockRecorder) Ping(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockCliToHubServer)(nil).Ping), arg0, arg1)
}

// PrepareInitCluster mocks base method
func (m *MockCliToHubServer) PrepareInitCluster(arg0 context.Context, arg1 *idl.PrepareInitClusterRequest) (*idl.PrepareInitClusterReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrepareInitCluster", arg0, arg1)
	ret0, _ := ret[0].(*idl.PrepareInitClusterReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrepareInitCluster indicates an expected call of PrepareInitCluster
func (mr *MockCliToHubServerMockRecorder) PrepareInitCluster(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareInitCluster", reflect.TypeOf((*MockCliToHubServer)(nil).PrepareInitCluster), arg0, arg1)
}

// PrepareShutdownClusters mocks base method
func (m *MockCliToHubServer) PrepareShutdownClusters(arg0 context.Context, arg1 *idl.PrepareShutdownClustersRequest) (*idl.PrepareShutdownClustersReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrepareShutdownClusters", arg0, arg1)
	ret0, _ := ret[0].(*idl.PrepareShutdownClustersReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrepareShutdownClusters indicates an expected call of PrepareShutdownClusters
func (mr *MockCliToHubServerMockRecorder) PrepareShutdownClusters(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareShutdownClusters", reflect.TypeOf((*MockCliToHubServer)(nil).PrepareShutdownClusters), arg0, arg1)
}

// PrepareStartAgents mocks base method
func (m *MockCliToHubServer) PrepareStartAgents(arg0 context.Context, arg1 *idl.PrepareStartAgentsRequest) (*idl.PrepareStartAgentsReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrepareStartAgents", arg0, arg1)
	ret0, _ := ret[0].(*idl.PrepareStartAgentsReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrepareStartAgents indicates an expected call of PrepareStartAgents
func (mr *MockCliToHubServerMockRecorder) PrepareStartAgents(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareStartAgents", reflect.TypeOf((*MockCliToHubServer)(nil).PrepareStartAgents), arg0, arg1)
}

// SetConfig mocks base method
func (m *MockCliToHubServer) SetConfig(arg0 context.Context, arg1 *idl.SetConfigRequest) (*idl.SetConfigReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetConfig", arg0, arg1)
	ret0, _ := ret[0].(*idl.SetConfigReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetConfig indicates an expected call of SetConfig
func (mr *MockCliToHubServerMockRecorder) SetConfig(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetConfig", reflect.TypeOf((*MockCliToHubServer)(nil).SetConfig), arg0, arg1)
}

// StatusConversion mocks base method
func (m *MockCliToHubServer) StatusConversion(arg0 context.Context, arg1 *idl.StatusConversionRequest) (*idl.StatusConversionReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StatusConversion", arg0, arg1)
	ret0, _ := ret[0].(*idl.StatusConversionReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StatusConversion indicates an expected call of StatusConversion
func (mr *MockCliToHubServerMockRecorder) StatusConversion(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StatusConversion", reflect.TypeOf((*MockCliToHubServer)(nil).StatusConversion), arg0, arg1)
}

// StatusUpgrade mocks base method
func (m *MockCliToHubServer) StatusUpgrade(arg0 context.Context, arg1 *idl.StatusUpgradeRequest) (*idl.StatusUpgradeReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StatusUpgrade", arg0, arg1)
	ret0, _ := ret[0].(*idl.StatusUpgradeReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StatusUpgrade indicates an expected call of StatusUpgrade
func (mr *MockCliToHubServerMockRecorder) StatusUpgrade(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StatusUpgrade", reflect.TypeOf((*MockCliToHubServer)(nil).StatusUpgrade), arg0, arg1)
}

// UpgradeConvertMaster mocks base method
func (m *MockCliToHubServer) UpgradeConvertMaster(arg0 *idl.UpgradeConvertMasterRequest, arg1 idl.CliToHub_UpgradeConvertMasterServer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpgradeConvertMaster", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpgradeConvertMaster indicates an expected call of UpgradeConvertMaster
func (mr *MockCliToHubServerMockRecorder) UpgradeConvertMaster(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpgradeConvertMaster", reflect.TypeOf((*MockCliToHubServer)(nil).UpgradeConvertMaster), arg0, arg1)
}

// UpgradeConvertPrimaries mocks base method
func (m *MockCliToHubServer) UpgradeConvertPrimaries(arg0 context.Context, arg1 *idl.UpgradeConvertPrimariesRequest) (*idl.UpgradeConvertPrimariesReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpgradeConvertPrimaries", arg0, arg1)
	ret0, _ := ret[0].(*idl.UpgradeConvertPrimariesReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpgradeConvertPrimaries indicates an expected call of UpgradeConvertPrimaries
func (mr *MockCliToHubServerMockRecorder) UpgradeConvertPrimaries(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpgradeConvertPrimaries", reflect.TypeOf((*MockCliToHubServer)(nil).UpgradeConvertPrimaries), arg0, arg1)
}

// UpgradeCopyMasterDataDir mocks base method
func (m *MockCliToHubServer) UpgradeCopyMasterDataDir(arg0 context.Context, arg1 *idl.UpgradeCopyMasterDataDirRequest) (*idl.UpgradeCopyMasterDataDirReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpgradeCopyMasterDataDir", arg0, arg1)
	ret0, _ := ret[0].(*idl.UpgradeCopyMasterDataDirReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpgradeCopyMasterDataDir indicates an expected call of UpgradeCopyMasterDataDir
func (mr *MockCliToHubServerMockRecorder) UpgradeCopyMasterDataDir(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpgradeCopyMasterDataDir", reflect.TypeOf((*MockCliToHubServer)(nil).UpgradeCopyMasterDataDir), arg0, arg1)
}

// UpgradeReconfigurePorts mocks base method
func (m *MockCliToHubServer) UpgradeReconfigurePorts(arg0 context.Context, arg1 *idl.UpgradeReconfigurePortsRequest) (*idl.UpgradeReconfigurePortsReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpgradeReconfigurePorts", arg0, arg1)
	ret0, _ := ret[0].(*idl.UpgradeReconfigurePortsReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpgradeReconfigurePorts indicates an expected call of UpgradeReconfigurePorts
func (mr *MockCliToHubServerMockRecorder) UpgradeReconfigurePorts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpgradeReconfigurePorts", reflect.TypeOf((*MockCliToHubServer)(nil).UpgradeReconfigurePorts), arg0, arg1)
}

// UpgradeValidateStartCluster mocks base method
func (m *MockCliToHubServer) UpgradeValidateStartCluster(arg0 context.Context, arg1 *idl.UpgradeValidateStartClusterRequest) (*idl.UpgradeValidateStartClusterReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpgradeValidateStartCluster", arg0, arg1)
	ret0, _ := ret[0].(*idl.UpgradeValidateStartClusterReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpgradeValidateStartCluster indicates an expected call of UpgradeValidateStartCluster
func (mr *MockCliToHubServerMockRecorder) UpgradeValidateStartCluster(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpgradeValidateStartCluster", reflect.TypeOf((*MockCliToHubServer)(nil).UpgradeValidateStartCluster), arg0, arg1)
}

// MockCliToHub_UpgradeConvertMasterServer is a mock of CliToHub_UpgradeConvertMasterServer interface
type MockCliToHub_UpgradeConvertMasterServer struct {
	ctrl     *gomock.Controller
	recorder *MockCliToHub_UpgradeConvertMasterServerMockRecorder
}

// MockCliToHub_UpgradeConvertMasterServerMockRecorder is the mock recorder for MockCliToHub_UpgradeConvertMasterServer
type MockCliToHub_UpgradeConvertMasterServerMockRecorder struct {
	mock *MockCliToHub_UpgradeConvertMasterServer
}

// NewMockCliToHub_UpgradeConvertMasterServer creates a new mock instance
func NewMockCliToHub_UpgradeConvertMasterServer(ctrl *gomock.Controller) *MockCliToHub_UpgradeConvertMasterServer {
	mock := &MockCliToHub_UpgradeConvertMasterServer{ctrl: ctrl}
	mock.recorder = &MockCliToHub_UpgradeConvertMasterServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCliToHub_UpgradeConvertMasterServer) EXPECT() *MockCliToHub_UpgradeConvertMasterServerMockRecorder {
	return m.recorder
}

// Context mocks base method
func (m *MockCliToHub_UpgradeConvertMasterServer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockCliToHub_UpgradeConvertMasterServerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockCliToHub_UpgradeConvertMasterServer)(nil).Context))
}

// RecvMsg mocks base method
func (m *MockCliToHub_UpgradeConvertMasterServer) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg
func (mr *MockCliToHub_UpgradeConvertMasterServerMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockCliToHub_UpgradeConvertMasterServer)(nil).RecvMsg), arg0)
}

// Send mocks base method
func (m *MockCliToHub_UpgradeConvertMasterServer) Send(arg0 *idl.Chunk) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send
func (mr *MockCliToHub_UpgradeConvertMasterServerMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockCliToHub_UpgradeConvertMasterServer)(nil).Send), arg0)
}

// SendHeader mocks base method
func (m *MockCliToHub_UpgradeConvertMasterServer) SendHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader
func (mr *MockCliToHub_UpgradeConvertMasterServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockCliToHub_UpgradeConvertMasterServer)(nil).SendHeader), arg0)
}

// SendMsg mocks base method
func (m *MockCliToHub_UpgradeConvertMasterServer) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg
func (mr *MockCliToHub_UpgradeConvertMasterServerMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockCliToHub_UpgradeConvertMasterServer)(nil).SendMsg), arg0)
}

// SetHeader mocks base method
func (m *MockCliToHub_UpgradeConvertMasterServer) SetHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader
func (mr *MockCliToHub_UpgradeConvertMasterServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockCliToHub_UpgradeConvertMasterServer)(nil).SetHeader), arg0)
}

// SetTrailer mocks base method
func (m *MockCliToHub_UpgradeConvertMasterServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer
func (mr *MockCliToHub_UpgradeConvertMasterServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockCliToHub_UpgradeConvertMasterServer)(nil).SetTrailer), arg0)
}
