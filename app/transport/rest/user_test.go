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

	type mockBehaviour func(u *MockUserUsecase, params entities.QueryParams, users []entities.User)

	tests := map[string]struct {
		params                entities.QueryParams
		mockBehaviour         mockBehaviour
		expectedStatusCode    int
		expectedResponsetBody string
	}{
		"succes": {
			params: entities.QueryParams{
				Offset: "10000",
				Limit:  "5",
				Sort:   "first_name",
			},
			mockBehaviour: func(u *MockUserUsecase, params entities.QueryParams, users []entities.User) {
				u.EXPECT().GetAll(params).Return(users, nil)
			},
			expectedResponsetBody: `{"results":[]}`,
			expectedStatusCode:    200,
		},
		"invalid-offset": {
			params: entities.QueryParams{
				Offset: "bad",
				Limit:  "5",
				Sort:   "first_name",
			},
			mockBehaviour:         nil,
			expectedStatusCode:    400,
			expectedResponsetBody: `{"details":"offset argument has to be a number", "message":"Bad request"}`,
		},
		"invalid-Limit": {
			params: entities.QueryParams{
				Offset: "10",
				Limit:  "bad",
				Sort:   "first_name",
			},
			mockBehaviour:         nil,
			expectedStatusCode:    400,
			expectedResponsetBody: `{"details":"limit argument has to be a number", "message":"Bad request"}`,
		},
		"bad-sorting": {
			params: entities.QueryParams{
				Offset: "10",
				Limit:  "5",
				Sort:   "first_name_bad",
			},
			mockBehaviour:         nil,
			expectedStatusCode:    400,
			expectedResponsetBody: "{\"details\": \"sort argument `_bad` does not fit list: [first_name last_name created_at ,]\", \"message\": \"Bad request\"}",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			// Init vars.
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			usecase := NewMockUserUsecase(ctrl)
			users := []entities.User{}

			if tt.mockBehaviour != nil {
				tt.mockBehaviour(usecase, tt.params, users)
			}

			uh := UserHandler{
				usecase: usecase,
				logger:  logger,
			}

			resp := httptest.NewRecorder()

			url := "/users?offset=" + tt.params.Offset + "&limit=" + tt.params.Limit + "&sort=" + tt.params.Sort

			// Test server.
			uh.GetAllUsers(resp, httptest.NewRequest(http.MethodGet, url, nil))
			// fmt.Println("----------------", resp.Body.String())
			require.Equal(t, tt.expectedStatusCode, resp.Code)
			require.JSONEq(t, tt.expectedResponsetBody, resp.Body.String())
		})
	}
}
