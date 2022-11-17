package rest

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-devs-ua/octagon/lgr"
	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

// func TestUserHandler_GetAllUsers(t *testing.T) {
// 	logger, err := lgr.New("INFO")
// 	if err != nil {
// 		t.FailNow()
// 	}

// 	tests := map[string]struct {
// 		requestJSON          string
// 		expectedStatusCode   int
// 		expectedResponseBody string
// 		usecaseBuilder       func(ctrl *gomock.Controller, params entities.QueryParams) UserUsecase
// 	}{
// 		"succes": {
// 			params: entities.QueryParams{
// 				Offset: "0",
// 				Limit:  "5",
// 			},
// 			usecaseBuilder: func(ctrl *gomock.Controller, params entities.QueryParams) UserUsecase {
// 				mock := NewMockUserUsecase(ctrl)

// 				mock.EXPECT().GetAll(params).Return(
// 					[]entities.User{
// 						{
// 							ID:        "931add34-1f6d-4c06-b0e8-c37ac1ca614c",
// 							FirstName: "John1",
// 							LastName:  "Doe",
// 							Email:     "john1@examlpe.com",
// 							Password:  "qwerty",
// 							CreatedAt: "2022-11-05T22:28:36.679554Z",
// 						},
// 						{
// 							ID:        "4fddf9a4-fbd1-4083-98aa-e4d0e584e7bb",
// 							FirstName: "John2",
// 							LastName:  "Doe",
// 							Email:     "john2@examlpe.com",
// 							Password:  "qwerty",
// 							CreatedAt: "2022-11-05T22:28:36.679554Z",
// 						},
// 					}, nil).Times(1)

// 				return mock
// 			},
// 			expectedStatusCode: 200,
// 			expectedResponseBody: `{
// 				"results": [
// 					{
// 						"id": "931add34-1f6d-4c06-b0e8-c37ac1ca614c",
// 						"email": "john1@examlpe.com",
// 						"first_name": "John1",
// 						"last_name": "Doe",
// 						"created_at": "2022-11-05T22:28:36.679554Z"
// 					},
// 					{
// 						"id": "4fddf9a4-fbd1-4083-98aa-e4d0e584e7bb",
// 						"email": "john2@examlpe.com",
// 						"first_name": "John2",
// 						"last_name": "Doe",
// 						"created_at": "2022-11-05T22:28:36.679554Z"
// 					}
// 				]
// 			}`,
// 		},
// 		"invalid-offset": {
// 			params: entities.QueryParams{
// 				Offset: "bad",
// 				Limit:  "5",
// 			},
// 			usecaseBuilder: func(ctrl *gomock.Controller, params entities.QueryParams) UserUsecase {
// 				return nil
// 			},
// 			expectedStatusCode:   400,
// 			expectedResponseBody: `{"details":"offset argument has to be a number", "message":"Bad request"}`,
// 		},
// 		"extralong-offset": {
// 			params: entities.QueryParams{
// 				Offset: "10000000000000000000000000000000000000000",
// 				Limit:  "5",
// 			},
// 			usecaseBuilder: func(ctrl *gomock.Controller, params entities.QueryParams) UserUsecase {
// 				return nil
// 			},
// 			expectedStatusCode:   400,
// 			expectedResponseBody: `{"details":"offset argument has to be a number", "message":"Bad request"}`,
// 		},
// 		"negative-offset": {
// 			params: entities.QueryParams{
// 				Offset: "-1",
// 				Limit:  "5",
// 			},
// 			usecaseBuilder: func(ctrl *gomock.Controller, params entities.QueryParams) UserUsecase {
// 				return nil
// 			},
// 			expectedStatusCode:   400,
// 			expectedResponseBody: `{"details":"offset argument has to be a positive number", "message":"Bad request"}`,
// 		},
// 		"invalid-Limit": {
// 			params: entities.QueryParams{
// 				Offset: "0",
// 				Limit:  "bad",
// 			},
// 			usecaseBuilder: func(ctrl *gomock.Controller, params entities.QueryParams) UserUsecase {
// 				return nil
// 			},
// 			expectedStatusCode:   400,
// 			expectedResponseBody: `{"details":"limit argument has to be a number", "message":"Bad request"}`,
// 		},
// 		"negative-Limit": {
// 			params: entities.QueryParams{
// 				Offset: "0",
// 				Limit:  "-1",
// 			},
// 			usecaseBuilder: func(ctrl *gomock.Controller, params entities.QueryParams) UserUsecase {
// 				return nil
// 			},
// 			expectedStatusCode:   400,
// 			expectedResponseBody: `{"details":"limit argument has to be a positive number", "message":"Bad request"}`,
// 		},
// 		"extralong-Limit": {
// 			params: entities.QueryParams{
// 				Offset: "0",
// 				Limit:  "1000000000000000000000000000000000000000000000",
// 			},
// 			usecaseBuilder: func(ctrl *gomock.Controller, params entities.QueryParams) UserUsecase {
// 				return nil
// 			},
// 			expectedStatusCode:   400,
// 			expectedResponseBody: `{"details":"limit argument has to be a number", "message":"Bad request"}`,
// 		},
// 		"bad-sorting": {
// 			params: entities.QueryParams{
// 				Offset: "10",
// 				Limit:  "5",
// 				Sort:   "first_name_bad",
// 			},
// 			usecaseBuilder: func(ctrl *gomock.Controller, params entities.QueryParams) UserUsecase {
// 				return nil
// 			},
// 			expectedStatusCode:   400,
// 			expectedResponseBody: `{"details": "sort argument '_bad' does not fit list: [first_name last_name created_at ,]", "message": "Bad request"}`,
// 		},
// 		"internal-server-error": {
// 			params: entities.QueryParams{
// 				Offset: "0",
// 				Limit:  "1",
// 			},
// 			usecaseBuilder: func(ctrl *gomock.Controller, params entities.QueryParams) UserUsecase {
// 				mock := NewMockUserUsecase(ctrl)
// 				mock.EXPECT().GetAll(params).Return(nil, errors.New("Internal error"))

// 				return mock
// 			},
// 			expectedStatusCode: 500,
// 			expectedResponseBody: `{
// 				"message": "Internal server error",
// 				"details": "could not fetch users"
// 			}`,
// 		},
// 	}

// 	for name, tt := range tests {
// 		t.Run(name, func(t *testing.T) {
// 			// Init vars.
// 			ctrl := gomock.NewController(t)
// 			defer ctrl.Finish()

// 			usecase := tt.usecaseBuilder(ctrl, tt.params)

// 			uh := NewUserHandler(usecase, logger)

// 			response := httptest.NewRecorder()

// 			request := httptest.NewRequest("GET", path.Join("/users"), nil)

// 			q := request.URL.Query()

// 			q.Add("offset", tt.params.Offset)
// 			q.Add("limit", tt.params.Limit)
// 			q.Add("sort", tt.params.Sort)

// 			request.URL.RawQuery = q.Encode()

// 			// Do testing.
// 			uh.GetAllUsers(response, request)

// 			// Check results of testing.
// 			require.Equal(t, tt.expectedStatusCode, response.Code)
// 			require.JSONEq(t, tt.expectedResponseBody, response.Body.String())
// 		})
// 	}
// }

func Test_New_Handler_CreateUser(t *testing.T) {

	logger, err := lgr.New(lgr.InfoLevel)
	if err != nil {
		t.Fatal("cannot initialize logger")
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			uh := NewUserHandler(tc.usecase, logger)

			buf := new(bytes.Buffer)
			json.NewEncoder(buf).Encode(tc.body)
			req := httptest.NewRequest(http.MethodPost, "/", buf)
			resp := httptest.NewRecorder()

			// Run origin handler.
			uh.CreateUser(resp, req)

			// Check results of testing.
			require.Equal(t, tc.expectedStatusCode, resp.Code)

			if name != "success" {
				require.Contains(t, resp.Body.String(), tc.expectedErrorMessage)
			}
		})
	}

}
