package rest

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"path"
	"testing"

	"github.com/go-devs-ua/octagon/app/globals"
	"github.com/gorilla/mux"

	"github.com/go-devs-ua/octagon/app/entities"
	"github.com/go-devs-ua/octagon/lgr"
	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestUserHandler_GetAllUsers(t *testing.T) {
	logger, err := lgr.New("INFO")
	if err != nil {
		t.FailNow()
	}

	tests := map[string]struct {
		params                entities.QueryParams
		expectedStatusCode    int
		expectedResponsetBody string
		usecaseBuilder        func(ctrl *gomock.Controller, params entities.QueryParams) UserUsecase
	}{
		"succes": {
			params: entities.QueryParams{
				Offset: "0",
				Limit:  "5",
			},
			usecaseBuilder: func(ctrl *gomock.Controller, params entities.QueryParams) UserUsecase {
				mock := NewMockUserUsecase(ctrl)

				mock.EXPECT().GetAll(params).Return(
					[]entities.User{
						{
							ID:        "931add34-1f6d-4c06-b0e8-c37ac1ca614c",
							FirstName: "John1",
							LastName:  "Doe",
							Email:     "john1@examlpe.com",
							Password:  "qwerty",
							CreatedAt: "2022-11-05T22:28:36.679554Z",
						},
						{
							ID:        "4fddf9a4-fbd1-4083-98aa-e4d0e584e7bb",
							FirstName: "John2",
							LastName:  "Doe",
							Email:     "john2@examlpe.com",
							Password:  "qwerty",
							CreatedAt: "2022-11-05T22:28:36.679554Z",
						},
					}, nil).Times(1)

				return mock
			},
			expectedStatusCode: 200,
			expectedResponsetBody: `{
				"results": [
					{
						"id": "931add34-1f6d-4c06-b0e8-c37ac1ca614c",
						"email": "john1@examlpe.com",
						"first_name": "John1",
						"last_name": "Doe",
						"created_at": "2022-11-05T22:28:36.679554Z"
					},
					{
						"id": "4fddf9a4-fbd1-4083-98aa-e4d0e584e7bb",
						"email": "john2@examlpe.com",
						"first_name": "John2",
						"last_name": "Doe",
						"created_at": "2022-11-05T22:28:36.679554Z"
					}
				]
			}`,
		},
		"invalid-offset": {
			params: entities.QueryParams{
				Offset: "bad",
				Limit:  "5",
			},
			usecaseBuilder: func(ctrl *gomock.Controller, params entities.QueryParams) UserUsecase {
				return nil
			},
			expectedStatusCode:    400,
			expectedResponsetBody: `{"details":"offset argument has to be a number", "message":"Bad request"}`,
		},
		"extralong-offset": {
			params: entities.QueryParams{
				Offset: "10000000000000000000000000000000000000000",
				Limit:  "5",
			},
			usecaseBuilder: func(ctrl *gomock.Controller, params entities.QueryParams) UserUsecase {
				return nil
			},
			expectedStatusCode:    400,
			expectedResponsetBody: `{"details":"offset argument has to be a number", "message":"Bad request"}`,
		},
		"negative-offset": {
			params: entities.QueryParams{
				Offset: "-1",
				Limit:  "5",
			},
			usecaseBuilder: func(ctrl *gomock.Controller, params entities.QueryParams) UserUsecase {
				return nil
			},
			expectedStatusCode:    400,
			expectedResponsetBody: `{"details":"offset argument has to be a positive number", "message":"Bad request"}`,
		},
		"invalid-Limit": {
			params: entities.QueryParams{
				Offset: "0",
				Limit:  "bad",
			},
			usecaseBuilder: func(ctrl *gomock.Controller, params entities.QueryParams) UserUsecase {
				return nil
			},
			expectedStatusCode:    400,
			expectedResponsetBody: `{"details":"limit argument has to be a number", "message":"Bad request"}`,
		},
		"negative-Limit": {
			params: entities.QueryParams{
				Offset: "0",
				Limit:  "-1",
			},
			usecaseBuilder: func(ctrl *gomock.Controller, params entities.QueryParams) UserUsecase {
				return nil
			},
			expectedStatusCode:    400,
			expectedResponsetBody: `{"details":"limit argument has to be a positive number", "message":"Bad request"}`,
		},
		"extralong-Limit": {
			params: entities.QueryParams{
				Offset: "0",
				Limit:  "1000000000000000000000000000000000000000000000",
			},
			usecaseBuilder: func(ctrl *gomock.Controller, params entities.QueryParams) UserUsecase {
				return nil
			},
			expectedStatusCode:    400,
			expectedResponsetBody: `{"details":"limit argument has to be a number", "message":"Bad request"}`,
		},
		"bad-sorting": {
			params: entities.QueryParams{
				Offset: "10",
				Limit:  "5",
				Sort:   "first_name_bad",
			},
			usecaseBuilder: func(ctrl *gomock.Controller, params entities.QueryParams) UserUsecase {
				return nil
			},
			expectedStatusCode:    400,
			expectedResponsetBody: `{"details": "sort argument '_bad' does not fit list: [first_name last_name created_at ,]", "message": "Bad request"}`,
		},
		"internal-server-error": {
			params: entities.QueryParams{
				Offset: "0",
				Limit:  "1",
			},
			usecaseBuilder: func(ctrl *gomock.Controller, params entities.QueryParams) UserUsecase {
				mock := NewMockUserUsecase(ctrl)
				mock.EXPECT().GetAll(params).Return(nil, errors.New("Internal error"))

				return mock
			},
			expectedStatusCode: 500,
			expectedResponsetBody: `{
				"message": "Internal server error",
				"details": "could not fetch users"
			}`,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			// Init vars.
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			usecase := tt.usecaseBuilder(ctrl, tt.params)

			uh := NewUserHandler(usecase, logger)

			response := httptest.NewRecorder()

			request := httptest.NewRequest("GET", path.Join("/users"), nil)

			q := request.URL.Query()

			q.Add("offset", tt.params.Offset)
			q.Add("limit", tt.params.Limit)
			q.Add("sort", tt.params.Sort)

			request.URL.RawQuery = q.Encode()

			// Do testing.
			uh.GetAllUsers(response, request)

			// Check results of testing.
			require.Equal(t, tt.expectedStatusCode, response.Code)
			require.JSONEq(t, tt.expectedResponsetBody, response.Body.String())
		})
	}
}

func TestUserHandler_GetUserByID(t *testing.T) {
	logger, err := lgr.New(lgr.InfoLevel)
	if err != nil {
		t.Fatal("cannot initialize logger")
	}

	tests := map[string]struct {
		id                    string
		usecaseBuilder        func(ctrl *gomock.Controller) UserUsecase
		expectedStatusCode    int
		expectedResponsetBody string
	}{
		"success": {
			id: "91e3dcf7-34a6-4646-bd37-383cc949da93",
			usecaseBuilder: func(ctrl *gomock.Controller) UserUsecase {
				mock := NewMockUserUsecase(ctrl)

				mock.EXPECT().GetByID("91e3dcf7-34a6-4646-bd37-383cc949da93").Return(&entities.User{
					ID:        "91e3dcf7-34a6-4646-bd37-383cc949da93",
					FirstName: "John",
					LastName:  "Dou",
					Email:     "j.dou@test.com",
					Password:  "12345678Qwerty",
					CreatedAt: "2022-01-01 00:00:10",
				}, nil).Times(1)

				return mock
			},
			expectedStatusCode: http.StatusOK,
			expectedResponsetBody: `{"id":"91e3dcf7-34a6-4646-bd37-383cc949da93", "first_name":"John", "last_name":"Dou", 
									"email":"j.dou@test.com", "created_at":"2022-01-01 00:00:10"}`,
		},
		"invalid_uuid": {
			id:                    "00000000--000-0000-0000-000000000000",
			usecaseBuilder:        func(ctrl *gomock.Controller) UserUsecase { return nil },
			expectedStatusCode:    http.StatusBadRequest,
			expectedResponsetBody: `{"message": "Bad request", "details": "invalid UUID format"}`,
		},
		"user_not_found": {
			id: "00000000-0000-0000-0000-000000000000",
			usecaseBuilder: func(ctrl *gomock.Controller) UserUsecase {
				mock := NewMockUserUsecase(ctrl)

				mock.EXPECT().GetByID("00000000-0000-0000-0000-000000000000").Return(nil,
					globals.ErrNotFound).Times(1)

				return mock
			},
			expectedStatusCode:    http.StatusNotFound,
			expectedResponsetBody: `{"message": "Not found", "details": "no user found in DB"}`,
		},
		"internal_error": {
			id: "10000000-0000-0000-0000-000000000000",
			usecaseBuilder: func(ctrl *gomock.Controller) UserUsecase {
				mock := NewMockUserUsecase(ctrl)

				mock.EXPECT().GetByID("10000000-0000-0000-0000-000000000000").Return(nil,
					errors.New("internal error while processing test")).Times(1)

				return mock
			},
			expectedStatusCode:    http.StatusInternalServerError,
			expectedResponsetBody: `{"message": "Internal server error", "details": ""}`,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			uh := UserHandler{
				usecase: tt.usecaseBuilder(ctrl),
				logger:  logger,
			}

			response := httptest.NewRecorder()

			request := httptest.NewRequest(http.MethodGet, "*", nil)
			request = mux.SetURLVars(request, map[string]string{"id": tt.id})

			uh.GetUserByID(response, request)

			require.Equal(t, tt.expectedStatusCode, response.Code)
			require.JSONEq(t, tt.expectedResponsetBody, response.Body.String())
		})
	}
}

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
