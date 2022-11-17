package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-devs-ua/octagon/app/entities"
	"github.com/go-devs-ua/octagon/lgr"
	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestUserHandler_CreateUser(t *testing.T) {

	logger, err := lgr.New(lgr.InfoLevel)
	if err != nil {
		t.Fatal("cannot initialize logger")
	}

	tests := map[string]struct {
		requestJBody         string
		expectedStatusCode   int
		expectedErrorMessage string
		usecaseBuilder       func(ctrl *gomock.Controller) UserUsecase
	}{
		"Success": {
			requestJBody:         `{"first_name":"John","last_name":"Doe","email":"john@email.com","password":"123456Aa"}`,
			expectedStatusCode:   http.StatusCreated,
			expectedErrorMessage: "",
			usecaseBuilder: func(ctrl *gomock.Controller) UserUsecase {
				mock := NewMockUserUsecase(ctrl)
				us := entities.User{
					FirstName: "John",
					LastName:  "Doe",
					Email:     "john@email.com",
					Password:  "123456Aa",
				}
				mock.EXPECT().SignUp(us).Return("UUID", nil).Times(1)

				return mock
			},
		},

		// "Success user without last_name": {
		// 	requestJBody:         `{"email": "john@email.com", "password": "123456Aa", "first_name": "John", "last_name": ""}`,
		// 	expectedStatusCode:   http.StatusCreated,
		// 	expectedErrorMessage: "",
		// 	usecaseBuilder: func(ctrl *gomock.Controller) UserUsecase {
		// 		mock := NewMockUserUsecase(ctrl)
		// 		us := entities.User{
		// 			FirstName: "John",
		// 			Email:     "john@email.com",
		// 			Password:  "123456Aa",
		// 		}
		// 		mock.EXPECT().SignUp(us).Return(nil).Times(1)

		// 		return mock
		// 	},
		// },

		// "Creates new user bad JSON": {
		// 	requestJBody:         `{"email": "john@email.com", "password": "123456Aa", "first_name": "John", "last_ame": "Doe"}`,
		// 	expectedStatusCode:   http.StatusBadRequest,
		// 	expectedErrorMessage: "Bad request",
		// 	usecaseBuilder: func(ctrl *gomock.Controller) UserUsecase {
		// 		return nil
		// 	},
		// },

		// "Creates new user email len 0": {
		// 	requestJBody:         `{"email": "", "password": "123456Aa", "first_name": "John", "last_name": "Doe"}`,
		// 	expectedStatusCode:   http.StatusBadRequest,
		// 	expectedErrorMessage: "Validation error",
		// 	usecaseBuilder: func(ctrl *gomock.Controller) UserUsecase {
		// 		return nil
		// 	},
		// },

		// "Creates new user email too long domain part": {
		// 	requestJBody:         `{"email": "mail@12345678900987654321ab12345678900987654321ab12345678900987654321ab12345678900987654321ab12345678900987654321ab12345678900987654321ab12345678900987654321ab12345678900987654321ab12345678900987654321ab12345678900987654321ab12345678900987654321ab1234567.qwerty", "password": "123456Aa", "firstName": "John", "lastName": "Doe"}`,
		// 	expectedStatusCode:   http.StatusBadRequest,
		// 	expectedErrorMessage: "Validation error",
		// 	usecaseBuilder: func(ctrl *gomock.Controller) UserUsecase {
		// 		return nil
		// 	},
		// },

		// "Creates new user email too long local part": {
		// 	requestJBody:         `{"email": "1234567812345678123456781234567812345678123456781234567812345678a@email.com", "password": "123456Aa", "first_name": "John", "last_name": "Doe"}`,
		// 	expectedStatusCode:   http.StatusBadRequest,
		// 	expectedErrorMessage: "Validation error",
		// 	usecaseBuilder: func(ctrl *gomock.Controller) UserUsecase {
		// 		return nil
		// 	},
		// },

		// "Creates new user email bad symbols": {
		// 	requestJBody:         `{"email": "世界nameBAD@email.com", "password": "123456Aa", "first_name": "John", "last_name": "Doe"}`,
		// 	expectedStatusCode:   http.StatusCreated,
		// 	expectedErrorMessage: "Validation error",
		// 	usecaseBuilder: func(ctrl *gomock.Controller) UserUsecase {
		// 		return nil
		// 	},
		// },

		// "Creates new user email without @": {
		// 	requestJBody:         `{"email": "johnemail.com", "password": "123456Aa", "first_name": "John", "last_name": "Doe"}`,
		// 	expectedStatusCode:   http.StatusBadRequest,
		// 	expectedErrorMessage: "Validation error",
		// 	usecaseBuilder: func(ctrl *gomock.Controller) UserUsecase {
		// 		return nil
		// 	},
		// },

		// "Creates new user password len 2": {
		// 	requestJBody:         `{"email": "john@email.com", "password": "12", "first_name": "John", "last_name": "Doe"}`,
		// 	expectedStatusCode:   http.StatusBadRequest,
		// 	expectedErrorMessage: "Validation error",
		// 	usecaseBuilder: func(ctrl *gomock.Controller) UserUsecase {
		// 		return nil
		// 	},
		// },

		// "Creates new user password bad symbol": {
		// 	requestJBody:         `{"email": "john@email.com", "password": "123456AaÜ", "first_name": "John", "last_name": "Doe"}`,
		// 	expectedStatusCode:   http.StatusBadRequest,
		// 	expectedErrorMessage: "Validation error",
		// 	usecaseBuilder: func(ctrl *gomock.Controller) UserUsecase {
		// 		return nil
		// 	},
		// },

		// "Creates new user first_name len 0": {
		// 	requestJBody:         `{"email": "john@email.com", "password": "123456Aa", "first_name":"", "last_name": "Doe"}`,
		// 	expectedStatusCode:   http.StatusBadRequest,
		// 	expectedErrorMessage: "Validation error",
		// 	usecaseBuilder: func(ctrl *gomock.Controller) UserUsecase {
		// 		return nil
		// 	},
		// },

		// "Creates new user first_name bad symbol": {
		// 	requestJBody:         `{"email": "john@email.com", "password": "123456Aa", "first_name": "John#", "last_name": "Doe"}`,
		// 	expectedStatusCode:   http.StatusBadRequest,
		// 	expectedErrorMessage: "Validation error",
		// 	usecaseBuilder: func(ctrl *gomock.Controller) UserUsecase {
		// 		return nil
		// 	},
		// },

		// "Creates new user last_name bad symbol": {
		// 	requestJBody:         `{"email": "john@email.com", "password": "123456Aa", "first_name": "John", "last_name": "Doe#"}`,
		// 	expectedStatusCode:   http.StatusBadRequest,
		// 	expectedErrorMessage: "Validation error",
		// 	usecaseBuilder: func(ctrl *gomock.Controller) UserUsecase {
		// 		return nil
		// 	},
		// },
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			uh := NewUserHandler(tc.usecaseBuilder(ctrl), logger)

			buf := new(bytes.Buffer)

			fmt.Println("-------------------------------", tc.requestJBody)

			json.NewEncoder(buf).Encode(tc.requestJBody)
			fmt.Println("------------------", buf)
			req := httptest.NewRequest(http.MethodPost, "/users", buf)
			fmt.Println("------------------", req)
			resp := httptest.NewRecorder()

			// Run origin handler.
			uh.CreateUser(resp, req)

			// Check results of testing.
			require.Equal(t, tc.expectedStatusCode, resp.Code)

			if tc.expectedStatusCode != http.StatusCreated {
				require.Contains(t, resp.Body.String(), tc.expectedErrorMessage)
			}
		})
	}

}
