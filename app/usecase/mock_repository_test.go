// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package usecase

import (
	"github.com/go-devs-ua/octagon/app/entities"
	"sync"
)

// Ensure, that RepositoryMock does implement Repository.
// If this is not the case, regenerate this file with moq.
var _ Repository = &RepositoryMock{}

// RepositoryMock is a mock implementation of Repository.
//
//	func TestSomethingThatUsesRepository(t *testing.T) {
//
//		// make and configure a mocked Repository
//		mockedRepository := &RepositoryMock{
//			AddUserFunc: func(user entities.User) (string, error) {
//				panic("mock out the AddUser method")
//			},
//			DeleteUserFunc: func(user entities.User) error {
//				panic("mock out the DeleteUser method")
//			},
//			FindUserFunc: func(s string) (*entities.User, error) {
//				panic("mock out the FindUser method")
//			},
//			GetAllUsersFunc: func(queryParams entities.QueryParams) ([]entities.User, error) {
//				panic("mock out the GetAllUsers method")
//			},
//		}
//
//		// use mockedRepository in code that requires Repository
//		// and then make assertions.
//
//	}
type RepositoryMock struct {
	// AddUserFunc mocks the AddUser method.
	AddUserFunc func(user entities.User) (string, error)

	// DeleteUserFunc mocks the DeleteUser method.
	DeleteUserFunc func(user entities.User) error

	// FindUserFunc mocks the FindUser method.
	FindUserFunc func(s string) (*entities.User, error)

	// GetAllUsersFunc mocks the GetAllUsers method.
	GetAllUsersFunc func(queryParams entities.QueryParams) ([]entities.User, error)

	// calls tracks calls to the methods.
	calls struct {
		// AddUser holds details about calls to the AddUser method.
		AddUser []struct {
			// User is the user argument value.
			User entities.User
		}
		// DeleteUser holds details about calls to the DeleteUser method.
		DeleteUser []struct {
			// User is the user argument value.
			User entities.User
		}
		// FindUser holds details about calls to the FindUser method.
		FindUser []struct {
			// S is the s argument value.
			S string
		}
		// GetAllUsers holds details about calls to the GetAllUsers method.
		GetAllUsers []struct {
			// QueryParams is the queryParams argument value.
			QueryParams entities.QueryParams
		}
	}
	lockAddUser     sync.RWMutex
	lockDeleteUser  sync.RWMutex
	lockFindUser    sync.RWMutex
	lockGetAllUsers sync.RWMutex
}

// AddUser calls AddUserFunc.
func (mock *RepositoryMock) AddUser(user entities.User) (string, error) {
	if mock.AddUserFunc == nil {
		panic("RepositoryMock.AddUserFunc: method is nil but Repository.AddUser was just called")
	}
	callInfo := struct {
		User entities.User
	}{
		User: user,
	}
	mock.lockAddUser.Lock()
	mock.calls.AddUser = append(mock.calls.AddUser, callInfo)
	mock.lockAddUser.Unlock()
	return mock.AddUserFunc(user)
}

// AddUserCalls gets all the calls that were made to AddUser.
// Check the length with:
//
//	len(mockedRepository.AddUserCalls())
func (mock *RepositoryMock) AddUserCalls() []struct {
	User entities.User
} {
	var calls []struct {
		User entities.User
	}
	mock.lockAddUser.RLock()
	calls = mock.calls.AddUser
	mock.lockAddUser.RUnlock()
	return calls
}

// DeleteUser calls DeleteUserFunc.
func (mock *RepositoryMock) DeleteUser(user entities.User) error {
	if mock.DeleteUserFunc == nil {
		panic("RepositoryMock.DeleteUserFunc: method is nil but Repository.DeleteUser was just called")
	}
	callInfo := struct {
		User entities.User
	}{
		User: user,
	}
	mock.lockDeleteUser.Lock()
	mock.calls.DeleteUser = append(mock.calls.DeleteUser, callInfo)
	mock.lockDeleteUser.Unlock()
	return mock.DeleteUserFunc(user)
}

// DeleteUserCalls gets all the calls that were made to DeleteUser.
// Check the length with:
//
//	len(mockedRepository.DeleteUserCalls())
func (mock *RepositoryMock) DeleteUserCalls() []struct {
	User entities.User
} {
	var calls []struct {
		User entities.User
	}
	mock.lockDeleteUser.RLock()
	calls = mock.calls.DeleteUser
	mock.lockDeleteUser.RUnlock()
	return calls
}

// FindUser calls FindUserFunc.
func (mock *RepositoryMock) FindUser(s string) (*entities.User, error) {
	if mock.FindUserFunc == nil {
		panic("RepositoryMock.FindUserFunc: method is nil but Repository.FindUser was just called")
	}
	callInfo := struct {
		S string
	}{
		S: s,
	}
	mock.lockFindUser.Lock()
	mock.calls.FindUser = append(mock.calls.FindUser, callInfo)
	mock.lockFindUser.Unlock()
	return mock.FindUserFunc(s)
}

// FindUserCalls gets all the calls that were made to FindUser.
// Check the length with:
//
//	len(mockedRepository.FindUserCalls())
func (mock *RepositoryMock) FindUserCalls() []struct {
	S string
} {
	var calls []struct {
		S string
	}
	mock.lockFindUser.RLock()
	calls = mock.calls.FindUser
	mock.lockFindUser.RUnlock()
	return calls
}

// GetAllUsers calls GetAllUsersFunc.
func (mock *RepositoryMock) GetAllUsers(queryParams entities.QueryParams) ([]entities.User, error) {
	if mock.GetAllUsersFunc == nil {
		panic("RepositoryMock.GetAllUsersFunc: method is nil but Repository.GetAllUsers was just called")
	}
	callInfo := struct {
		QueryParams entities.QueryParams
	}{
		QueryParams: queryParams,
	}
	mock.lockGetAllUsers.Lock()
	mock.calls.GetAllUsers = append(mock.calls.GetAllUsers, callInfo)
	mock.lockGetAllUsers.Unlock()
	return mock.GetAllUsersFunc(queryParams)
}

// GetAllUsersCalls gets all the calls that were made to GetAllUsers.
// Check the length with:
//
//	len(mockedRepository.GetAllUsersCalls())
func (mock *RepositoryMock) GetAllUsersCalls() []struct {
	QueryParams entities.QueryParams
} {
	var calls []struct {
		QueryParams entities.QueryParams
	}
	mock.lockGetAllUsers.RLock()
	calls = mock.calls.GetAllUsers
	mock.lockGetAllUsers.RUnlock()
	return calls
}
