// Code generated by MockGen. DO NOT EDIT.
// Source: dal/leaderboard/dal.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	leaderboard "github.com/byyjoww/leaderboard/dal/leaderboard"
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockLeaderboardDAL is a mock of LeaderboardDAL interface.
type MockLeaderboardDAL struct {
	ctrl     *gomock.Controller
	recorder *MockLeaderboardDALMockRecorder
}

// MockLeaderboardDALMockRecorder is the mock recorder for MockLeaderboardDAL.
type MockLeaderboardDALMockRecorder struct {
	mock *MockLeaderboardDAL
}

// NewMockLeaderboardDAL creates a new mock instance.
func NewMockLeaderboardDAL(ctrl *gomock.Controller) *MockLeaderboardDAL {
	mock := &MockLeaderboardDAL{ctrl: ctrl}
	mock.recorder = &MockLeaderboardDALMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLeaderboardDAL) EXPECT() *MockLeaderboardDALMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockLeaderboardDAL) Create(leaderboard *leaderboard.Leaderboard) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", leaderboard)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockLeaderboardDALMockRecorder) Create(leaderboard interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockLeaderboardDAL)(nil).Create), leaderboard)
}

// Delete mocks base method.
func (m *MockLeaderboardDAL) Delete(leaderboard *leaderboard.Leaderboard) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", leaderboard)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockLeaderboardDALMockRecorder) Delete(leaderboard interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockLeaderboardDAL)(nil).Delete), leaderboard)
}

// GetByPK mocks base method.
func (m *MockLeaderboardDAL) GetByPK(id uuid.UUID) (*leaderboard.Leaderboard, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByPK", id)
	ret0, _ := ret[0].(*leaderboard.Leaderboard)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByPK indicates an expected call of GetByPK.
func (mr *MockLeaderboardDALMockRecorder) GetByPK(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByPK", reflect.TypeOf((*MockLeaderboardDAL)(nil).GetByPK), id)
}

// List mocks base method.
func (m *MockLeaderboardDAL) List() ([]*leaderboard.Leaderboard, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]*leaderboard.Leaderboard)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockLeaderboardDALMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockLeaderboardDAL)(nil).List))
}