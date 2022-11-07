package rest

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-devs-ua/octagon/app/entities"
	"github.com/go-devs-ua/octagon/lgr"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestUserHandler_GetAllUsers(t *testing.T) {
	logger, err := lgr.New("INFO")
	if err != nil {
		t.FailNow()
	}

	type UserUsecase func(u *MockUserUsecase, params entities.QueryParams, users []entities.User)

	tests := map[string]struct {
		params                entities.QueryParams
		GetAll                UserUsecase
		expectedStatusCode    int
		expectedResponsetBody string
	}{
		"succes": {
			params: entities.QueryParams{
				Offset: "0",
				Limit:  "5",
			},
			GetAll: func(u *MockUserUsecase, params entities.QueryParams, users []entities.User) {
				u.EXPECT().GetAll(params).Return(
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
					}, nil)
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
			expectedStatusCode:    400,
			expectedResponsetBody: `{"details":"offset argument has to be a number", "message":"Bad request"}`,
		},
		"extralong-offset": {
			params: entities.QueryParams{
				Offset: "10000000000000000000000000000000000000000",
				Limit:  "5",
			},
			expectedStatusCode:    400,
			expectedResponsetBody: `{"details":"offset argument has to be a number", "message":"Bad request"}`,
		},
		"negative-offset": {
			params: entities.QueryParams{
				Offset: "-1",
				Limit:  "5",
			},
			expectedStatusCode:    400,
			expectedResponsetBody: `{"details":"offset argument has to be a number", "message":"Bad request"}`,
		},
		"invalid-Limit": {
			params: entities.QueryParams{
				Offset: "0",
				Limit:  "bad",
			},
			expectedStatusCode:    400,
			expectedResponsetBody: `{"details":"limit argument has to be a number", "message":"Bad request"}`,
		},
		"negative-Limit": {
			params: entities.QueryParams{
				Offset: "0",
				Limit:  "-1",
			},
			expectedStatusCode:    400,
			expectedResponsetBody: `{"details":"limit argument has to be a number", "message":"Bad request"}`,
		},
		"extralong-Limit": {
			params: entities.QueryParams{
				Offset: "0",
				Limit:  "1000000000000000000000000000000000000000000000",
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
			expectedStatusCode:    400,
			expectedResponsetBody: "{\"details\": \"sort argument `_bad` does not fit list: [first_name last_name created_at ,]\", \"message\": \"Bad request\"}",
		},
		"sorting-by-first_name": {
			params: entities.QueryParams{
				Offset: "10",
				Limit:  "5",
				Sort:   "first_name",
			},
			GetAll: func(u *MockUserUsecase, params entities.QueryParams, users []entities.User) {
				u.EXPECT().GetAll(params).Return(
					[]entities.User{
						{
							ID:        "931add34-1f6d-4c06-b0e8-c37ac1ca614c",
							FirstName: "B",
							LastName:  "Doe",
							Email:     "john1@examlpe.com",
							Password:  "qwerty",
							CreatedAt: "2022-11-05T22:28:36.679554Z",
						},
						{
							ID:        "4fddf9a4-fbd1-4083-98aa-e4d0e584e7bb",
							FirstName: "A",
							LastName:  "Doe",
							Email:     "john2@examlpe.com",
							Password:  "qwerty",
							CreatedAt: "2022-11-05T22:28:36.679554Z",
						},
					}, nil)
			},
			expectedStatusCode: 400,
			expectedResponsetBody: `{
				"results": [
					{
						"id": "931add34-1f6d-4c06-b0e8-c37ac1ca614c",
						"email": "john1@examlpe.com",
						"first_name": "A",
						"last_name": "Doe",
						"created_at": "2022-11-05T22:28:36.679554Z"
					},
					{
						"id": "4fddf9a4-fbd1-4083-98aa-e4d0e584e7bb",
						"email": "john2@examlpe.com",
						"first_name": "B",
						"last_name": "Doe",
						"created_at": "2022-11-05T22:28:36.679554Z"
					}
				]
			}`,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			// Init vars.
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			usecase := NewMockUserUsecase(ctrl)
			users := []entities.User{}

			if tt.GetAll != nil {
				tt.GetAll(usecase, tt.params, users)
			}

			uh := UserHandler{
				usecase: usecase,
				logger:  logger,
			}

			resp := httptest.NewRecorder()

			url := "/users?offset=" + tt.params.Offset + "&limit=" + tt.params.Limit + "&sort=" + tt.params.Sort

			// Do testing.
			uh.GetAllUsers(resp, httptest.NewRequest(http.MethodGet, url, nil))

			// Check results of testing.
			require.Equal(t, tt.expectedStatusCode, resp.Code)
			require.JSONEq(t, tt.expectedResponsetBody, resp.Body.String())
		})
	}
}
