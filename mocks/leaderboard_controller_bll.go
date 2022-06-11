// Code generated by MockGen. DO NOT EDIT.
// Source: bll/leaderboard/controller.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	leaderboard "github.com/byyjoww/leaderboard/dal/leaderboard"
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockLeaderboardController is a mock of LeaderboardController interface.
type MockLeaderboardController struct {
	ctrl     *gomock.Controller
	recorder *MockLeaderboardControllerMockRecorder
}

// MockLeaderboardControllerMockRecorder is the mock recorder for MockLeaderboardController.
type MockLeaderboardControllerMockRecorder struct {
	mock *MockLeaderboardController
}

// NewMockLeaderboardController creates a new mock instance.
func NewMockLeaderboardController(ctrl *gomock.Controller) *MockLeaderboardController {
	mock := &MockLeaderboardController{ctrl: ctrl}
	mock.recorder = &MockLeaderboardControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLeaderboardController) EXPECT() *MockLeaderboardControllerMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockLeaderboardController) Create() (*leaderboard.Leaderboard, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create")
	ret0, _ := ret[0].(*leaderboard.Leaderboard)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockLeaderboardControllerMockRecorder) Create() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockLeaderboardController)(nil).Create))
}

// Get mocks base method.
func (m *MockLeaderboardController) Get(leaderboardId uuid.UUID) (*leaderboard.Leaderboard, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", leaderboardId)
	ret0, _ := ret[0].(*leaderboard.Leaderboard)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockLeaderboardControllerMockRecorder) Get(leaderboardId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockLeaderboardController)(nil).Get), leaderboardId)
}

// List mocks base method.
func (m *MockLeaderboardController) List(limit, offset int) ([]*leaderboard.Leaderboard, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", limit, offset)
	ret0, _ := ret[0].([]*leaderboard.Leaderboard)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockLeaderboardControllerMockRecorder) List(limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockLeaderboardController)(nil).List), limit, offset)
}

// Remove mocks base method.
func (m *MockLeaderboardController) Remove(leaderboardId uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove", leaderboardId)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove.
func (mr *MockLeaderboardControllerMockRecorder) Remove(leaderboardId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockLeaderboardController)(nil).Remove), leaderboardId)
}
