// Code generated by MockGen. DO NOT EDIT.
// Source: guardduty.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	guardduty "github.com/aws/aws-sdk-go-v2/service/guardduty"
	gomock "github.com/golang/mock/gomock"
)

// MockGuarddutyClient is a mock of GuarddutyClient interface.
type MockGuarddutyClient struct {
	ctrl     *gomock.Controller
	recorder *MockGuarddutyClientMockRecorder
}

// MockGuarddutyClientMockRecorder is the mock recorder for MockGuarddutyClient.
type MockGuarddutyClientMockRecorder struct {
	mock *MockGuarddutyClient
}

// NewMockGuarddutyClient creates a new mock instance.
func NewMockGuarddutyClient(ctrl *gomock.Controller) *MockGuarddutyClient {
	mock := &MockGuarddutyClient{ctrl: ctrl}
	mock.recorder = &MockGuarddutyClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGuarddutyClient) EXPECT() *MockGuarddutyClientMockRecorder {
	return m.recorder
}

// DescribeMalwareScans mocks base method.
func (m *MockGuarddutyClient) DescribeMalwareScans(arg0 context.Context, arg1 *guardduty.DescribeMalwareScansInput, arg2 ...func(*guardduty.Options)) (*guardduty.DescribeMalwareScansOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeMalwareScans", varargs...)
	ret0, _ := ret[0].(*guardduty.DescribeMalwareScansOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeMalwareScans indicates an expected call of DescribeMalwareScans.
func (mr *MockGuarddutyClientMockRecorder) DescribeMalwareScans(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeMalwareScans", reflect.TypeOf((*MockGuarddutyClient)(nil).DescribeMalwareScans), varargs...)
}

// DescribeOrganizationConfiguration mocks base method.
func (m *MockGuarddutyClient) DescribeOrganizationConfiguration(arg0 context.Context, arg1 *guardduty.DescribeOrganizationConfigurationInput, arg2 ...func(*guardduty.Options)) (*guardduty.DescribeOrganizationConfigurationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeOrganizationConfiguration", varargs...)
	ret0, _ := ret[0].(*guardduty.DescribeOrganizationConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeOrganizationConfiguration indicates an expected call of DescribeOrganizationConfiguration.
func (mr *MockGuarddutyClientMockRecorder) DescribeOrganizationConfiguration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeOrganizationConfiguration", reflect.TypeOf((*MockGuarddutyClient)(nil).DescribeOrganizationConfiguration), varargs...)
}

// DescribePublishingDestination mocks base method.
func (m *MockGuarddutyClient) DescribePublishingDestination(arg0 context.Context, arg1 *guardduty.DescribePublishingDestinationInput, arg2 ...func(*guardduty.Options)) (*guardduty.DescribePublishingDestinationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribePublishingDestination", varargs...)
	ret0, _ := ret[0].(*guardduty.DescribePublishingDestinationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribePublishingDestination indicates an expected call of DescribePublishingDestination.
func (mr *MockGuarddutyClientMockRecorder) DescribePublishingDestination(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribePublishingDestination", reflect.TypeOf((*MockGuarddutyClient)(nil).DescribePublishingDestination), varargs...)
}

// GetAdministratorAccount mocks base method.
func (m *MockGuarddutyClient) GetAdministratorAccount(arg0 context.Context, arg1 *guardduty.GetAdministratorAccountInput, arg2 ...func(*guardduty.Options)) (*guardduty.GetAdministratorAccountOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAdministratorAccount", varargs...)
	ret0, _ := ret[0].(*guardduty.GetAdministratorAccountOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAdministratorAccount indicates an expected call of GetAdministratorAccount.
func (mr *MockGuarddutyClientMockRecorder) GetAdministratorAccount(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAdministratorAccount", reflect.TypeOf((*MockGuarddutyClient)(nil).GetAdministratorAccount), varargs...)
}

// GetDetector mocks base method.
func (m *MockGuarddutyClient) GetDetector(arg0 context.Context, arg1 *guardduty.GetDetectorInput, arg2 ...func(*guardduty.Options)) (*guardduty.GetDetectorOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetDetector", varargs...)
	ret0, _ := ret[0].(*guardduty.GetDetectorOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDetector indicates an expected call of GetDetector.
func (mr *MockGuarddutyClientMockRecorder) GetDetector(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDetector", reflect.TypeOf((*MockGuarddutyClient)(nil).GetDetector), varargs...)
}

// GetFilter mocks base method.
func (m *MockGuarddutyClient) GetFilter(arg0 context.Context, arg1 *guardduty.GetFilterInput, arg2 ...func(*guardduty.Options)) (*guardduty.GetFilterOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetFilter", varargs...)
	ret0, _ := ret[0].(*guardduty.GetFilterOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFilter indicates an expected call of GetFilter.
func (mr *MockGuarddutyClientMockRecorder) GetFilter(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFilter", reflect.TypeOf((*MockGuarddutyClient)(nil).GetFilter), varargs...)
}

// GetFindings mocks base method.
func (m *MockGuarddutyClient) GetFindings(arg0 context.Context, arg1 *guardduty.GetFindingsInput, arg2 ...func(*guardduty.Options)) (*guardduty.GetFindingsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetFindings", varargs...)
	ret0, _ := ret[0].(*guardduty.GetFindingsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFindings indicates an expected call of GetFindings.
func (mr *MockGuarddutyClientMockRecorder) GetFindings(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFindings", reflect.TypeOf((*MockGuarddutyClient)(nil).GetFindings), varargs...)
}

// GetFindingsStatistics mocks base method.
func (m *MockGuarddutyClient) GetFindingsStatistics(arg0 context.Context, arg1 *guardduty.GetFindingsStatisticsInput, arg2 ...func(*guardduty.Options)) (*guardduty.GetFindingsStatisticsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetFindingsStatistics", varargs...)
	ret0, _ := ret[0].(*guardduty.GetFindingsStatisticsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFindingsStatistics indicates an expected call of GetFindingsStatistics.
func (mr *MockGuarddutyClientMockRecorder) GetFindingsStatistics(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFindingsStatistics", reflect.TypeOf((*MockGuarddutyClient)(nil).GetFindingsStatistics), varargs...)
}

// GetIPSet mocks base method.
func (m *MockGuarddutyClient) GetIPSet(arg0 context.Context, arg1 *guardduty.GetIPSetInput, arg2 ...func(*guardduty.Options)) (*guardduty.GetIPSetOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetIPSet", varargs...)
	ret0, _ := ret[0].(*guardduty.GetIPSetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIPSet indicates an expected call of GetIPSet.
func (mr *MockGuarddutyClientMockRecorder) GetIPSet(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIPSet", reflect.TypeOf((*MockGuarddutyClient)(nil).GetIPSet), varargs...)
}

// GetInvitationsCount mocks base method.
func (m *MockGuarddutyClient) GetInvitationsCount(arg0 context.Context, arg1 *guardduty.GetInvitationsCountInput, arg2 ...func(*guardduty.Options)) (*guardduty.GetInvitationsCountOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetInvitationsCount", varargs...)
	ret0, _ := ret[0].(*guardduty.GetInvitationsCountOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInvitationsCount indicates an expected call of GetInvitationsCount.
func (mr *MockGuarddutyClientMockRecorder) GetInvitationsCount(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInvitationsCount", reflect.TypeOf((*MockGuarddutyClient)(nil).GetInvitationsCount), varargs...)
}

// GetMalwareScanSettings mocks base method.
func (m *MockGuarddutyClient) GetMalwareScanSettings(arg0 context.Context, arg1 *guardduty.GetMalwareScanSettingsInput, arg2 ...func(*guardduty.Options)) (*guardduty.GetMalwareScanSettingsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMalwareScanSettings", varargs...)
	ret0, _ := ret[0].(*guardduty.GetMalwareScanSettingsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMalwareScanSettings indicates an expected call of GetMalwareScanSettings.
func (mr *MockGuarddutyClientMockRecorder) GetMalwareScanSettings(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMalwareScanSettings", reflect.TypeOf((*MockGuarddutyClient)(nil).GetMalwareScanSettings), varargs...)
}

// GetMasterAccount mocks base method.
func (m *MockGuarddutyClient) GetMasterAccount(arg0 context.Context, arg1 *guardduty.GetMasterAccountInput, arg2 ...func(*guardduty.Options)) (*guardduty.GetMasterAccountOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMasterAccount", varargs...)
	ret0, _ := ret[0].(*guardduty.GetMasterAccountOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMasterAccount indicates an expected call of GetMasterAccount.
func (mr *MockGuarddutyClientMockRecorder) GetMasterAccount(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMasterAccount", reflect.TypeOf((*MockGuarddutyClient)(nil).GetMasterAccount), varargs...)
}

// GetMemberDetectors mocks base method.
func (m *MockGuarddutyClient) GetMemberDetectors(arg0 context.Context, arg1 *guardduty.GetMemberDetectorsInput, arg2 ...func(*guardduty.Options)) (*guardduty.GetMemberDetectorsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMemberDetectors", varargs...)
	ret0, _ := ret[0].(*guardduty.GetMemberDetectorsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMemberDetectors indicates an expected call of GetMemberDetectors.
func (mr *MockGuarddutyClientMockRecorder) GetMemberDetectors(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMemberDetectors", reflect.TypeOf((*MockGuarddutyClient)(nil).GetMemberDetectors), varargs...)
}

// GetMembers mocks base method.
func (m *MockGuarddutyClient) GetMembers(arg0 context.Context, arg1 *guardduty.GetMembersInput, arg2 ...func(*guardduty.Options)) (*guardduty.GetMembersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMembers", varargs...)
	ret0, _ := ret[0].(*guardduty.GetMembersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMembers indicates an expected call of GetMembers.
func (mr *MockGuarddutyClientMockRecorder) GetMembers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMembers", reflect.TypeOf((*MockGuarddutyClient)(nil).GetMembers), varargs...)
}

// GetRemainingFreeTrialDays mocks base method.
func (m *MockGuarddutyClient) GetRemainingFreeTrialDays(arg0 context.Context, arg1 *guardduty.GetRemainingFreeTrialDaysInput, arg2 ...func(*guardduty.Options)) (*guardduty.GetRemainingFreeTrialDaysOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRemainingFreeTrialDays", varargs...)
	ret0, _ := ret[0].(*guardduty.GetRemainingFreeTrialDaysOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRemainingFreeTrialDays indicates an expected call of GetRemainingFreeTrialDays.
func (mr *MockGuarddutyClientMockRecorder) GetRemainingFreeTrialDays(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRemainingFreeTrialDays", reflect.TypeOf((*MockGuarddutyClient)(nil).GetRemainingFreeTrialDays), varargs...)
}

// GetThreatIntelSet mocks base method.
func (m *MockGuarddutyClient) GetThreatIntelSet(arg0 context.Context, arg1 *guardduty.GetThreatIntelSetInput, arg2 ...func(*guardduty.Options)) (*guardduty.GetThreatIntelSetOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetThreatIntelSet", varargs...)
	ret0, _ := ret[0].(*guardduty.GetThreatIntelSetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetThreatIntelSet indicates an expected call of GetThreatIntelSet.
func (mr *MockGuarddutyClientMockRecorder) GetThreatIntelSet(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetThreatIntelSet", reflect.TypeOf((*MockGuarddutyClient)(nil).GetThreatIntelSet), varargs...)
}

// GetUsageStatistics mocks base method.
func (m *MockGuarddutyClient) GetUsageStatistics(arg0 context.Context, arg1 *guardduty.GetUsageStatisticsInput, arg2 ...func(*guardduty.Options)) (*guardduty.GetUsageStatisticsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetUsageStatistics", varargs...)
	ret0, _ := ret[0].(*guardduty.GetUsageStatisticsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsageStatistics indicates an expected call of GetUsageStatistics.
func (mr *MockGuarddutyClientMockRecorder) GetUsageStatistics(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsageStatistics", reflect.TypeOf((*MockGuarddutyClient)(nil).GetUsageStatistics), varargs...)
}

// ListDetectors mocks base method.
func (m *MockGuarddutyClient) ListDetectors(arg0 context.Context, arg1 *guardduty.ListDetectorsInput, arg2 ...func(*guardduty.Options)) (*guardduty.ListDetectorsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListDetectors", varargs...)
	ret0, _ := ret[0].(*guardduty.ListDetectorsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDetectors indicates an expected call of ListDetectors.
func (mr *MockGuarddutyClientMockRecorder) ListDetectors(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDetectors", reflect.TypeOf((*MockGuarddutyClient)(nil).ListDetectors), varargs...)
}

// ListFilters mocks base method.
func (m *MockGuarddutyClient) ListFilters(arg0 context.Context, arg1 *guardduty.ListFiltersInput, arg2 ...func(*guardduty.Options)) (*guardduty.ListFiltersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListFilters", varargs...)
	ret0, _ := ret[0].(*guardduty.ListFiltersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFilters indicates an expected call of ListFilters.
func (mr *MockGuarddutyClientMockRecorder) ListFilters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFilters", reflect.TypeOf((*MockGuarddutyClient)(nil).ListFilters), varargs...)
}

// ListFindings mocks base method.
func (m *MockGuarddutyClient) ListFindings(arg0 context.Context, arg1 *guardduty.ListFindingsInput, arg2 ...func(*guardduty.Options)) (*guardduty.ListFindingsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListFindings", varargs...)
	ret0, _ := ret[0].(*guardduty.ListFindingsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFindings indicates an expected call of ListFindings.
func (mr *MockGuarddutyClientMockRecorder) ListFindings(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFindings", reflect.TypeOf((*MockGuarddutyClient)(nil).ListFindings), varargs...)
}

// ListIPSets mocks base method.
func (m *MockGuarddutyClient) ListIPSets(arg0 context.Context, arg1 *guardduty.ListIPSetsInput, arg2 ...func(*guardduty.Options)) (*guardduty.ListIPSetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListIPSets", varargs...)
	ret0, _ := ret[0].(*guardduty.ListIPSetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListIPSets indicates an expected call of ListIPSets.
func (mr *MockGuarddutyClientMockRecorder) ListIPSets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListIPSets", reflect.TypeOf((*MockGuarddutyClient)(nil).ListIPSets), varargs...)
}

// ListInvitations mocks base method.
func (m *MockGuarddutyClient) ListInvitations(arg0 context.Context, arg1 *guardduty.ListInvitationsInput, arg2 ...func(*guardduty.Options)) (*guardduty.ListInvitationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListInvitations", varargs...)
	ret0, _ := ret[0].(*guardduty.ListInvitationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListInvitations indicates an expected call of ListInvitations.
func (mr *MockGuarddutyClientMockRecorder) ListInvitations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListInvitations", reflect.TypeOf((*MockGuarddutyClient)(nil).ListInvitations), varargs...)
}

// ListMembers mocks base method.
func (m *MockGuarddutyClient) ListMembers(arg0 context.Context, arg1 *guardduty.ListMembersInput, arg2 ...func(*guardduty.Options)) (*guardduty.ListMembersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListMembers", varargs...)
	ret0, _ := ret[0].(*guardduty.ListMembersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListMembers indicates an expected call of ListMembers.
func (mr *MockGuarddutyClientMockRecorder) ListMembers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMembers", reflect.TypeOf((*MockGuarddutyClient)(nil).ListMembers), varargs...)
}

// ListOrganizationAdminAccounts mocks base method.
func (m *MockGuarddutyClient) ListOrganizationAdminAccounts(arg0 context.Context, arg1 *guardduty.ListOrganizationAdminAccountsInput, arg2 ...func(*guardduty.Options)) (*guardduty.ListOrganizationAdminAccountsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListOrganizationAdminAccounts", varargs...)
	ret0, _ := ret[0].(*guardduty.ListOrganizationAdminAccountsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListOrganizationAdminAccounts indicates an expected call of ListOrganizationAdminAccounts.
func (mr *MockGuarddutyClientMockRecorder) ListOrganizationAdminAccounts(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOrganizationAdminAccounts", reflect.TypeOf((*MockGuarddutyClient)(nil).ListOrganizationAdminAccounts), varargs...)
}

// ListPublishingDestinations mocks base method.
func (m *MockGuarddutyClient) ListPublishingDestinations(arg0 context.Context, arg1 *guardduty.ListPublishingDestinationsInput, arg2 ...func(*guardduty.Options)) (*guardduty.ListPublishingDestinationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListPublishingDestinations", varargs...)
	ret0, _ := ret[0].(*guardduty.ListPublishingDestinationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPublishingDestinations indicates an expected call of ListPublishingDestinations.
func (mr *MockGuarddutyClientMockRecorder) ListPublishingDestinations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPublishingDestinations", reflect.TypeOf((*MockGuarddutyClient)(nil).ListPublishingDestinations), varargs...)
}

// ListTagsForResource mocks base method.
func (m *MockGuarddutyClient) ListTagsForResource(arg0 context.Context, arg1 *guardduty.ListTagsForResourceInput, arg2 ...func(*guardduty.Options)) (*guardduty.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTagsForResource", varargs...)
	ret0, _ := ret[0].(*guardduty.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsForResource indicates an expected call of ListTagsForResource.
func (mr *MockGuarddutyClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResource", reflect.TypeOf((*MockGuarddutyClient)(nil).ListTagsForResource), varargs...)
}

// ListThreatIntelSets mocks base method.
func (m *MockGuarddutyClient) ListThreatIntelSets(arg0 context.Context, arg1 *guardduty.ListThreatIntelSetsInput, arg2 ...func(*guardduty.Options)) (*guardduty.ListThreatIntelSetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListThreatIntelSets", varargs...)
	ret0, _ := ret[0].(*guardduty.ListThreatIntelSetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListThreatIntelSets indicates an expected call of ListThreatIntelSets.
func (mr *MockGuarddutyClientMockRecorder) ListThreatIntelSets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListThreatIntelSets", reflect.TypeOf((*MockGuarddutyClient)(nil).ListThreatIntelSets), varargs...)
}
