// Code generated by MockGen. DO NOT EDIT.
// Source: ../repository.go

// Package mock_film is a generated GoMock package.
package mock_film

import (
	models "github.com/go-park-mail-ru/2020_1_k-on/application/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockRepository is a mock of Repository interface
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockRepository) Create(film *models.Film) (models.Film, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", film)
	ret0, _ := ret[0].(models.Film)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockRepositoryMockRecorder) Create(film interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRepository)(nil).Create), film)
}

// GetById mocks base method
func (m *MockRepository) GetById(id uint) (*models.Film, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", id)
	ret0, _ := ret[0].(*models.Film)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetById indicates an expected call of GetById
func (mr *MockRepositoryMockRecorder) GetById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockRepository)(nil).GetById), id)
}

// GetByName mocks base method
func (m *MockRepository) GetByName(name string) (*models.Film, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByName", name)
	ret0, _ := ret[0].(*models.Film)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetByName indicates an expected call of GetByName
func (mr *MockRepositoryMockRecorder) GetByName(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByName", reflect.TypeOf((*MockRepository)(nil).GetByName), name)
}

// GetFilmsArr mocks base method
func (m *MockRepository) GetFilmsArr(begin, end uint) (*models.Films, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFilmsArr", begin, end)
	ret0, _ := ret[0].(*models.Films)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetFilmsArr indicates an expected call of GetFilmsArr
func (mr *MockRepositoryMockRecorder) GetFilmsArr(begin, end interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFilmsArr", reflect.TypeOf((*MockRepository)(nil).GetFilmsArr), begin, end)
}
