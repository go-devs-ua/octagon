package rest

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"path"
	"testing"

	"github.com/go-devs-ua/octagon/app/entities"
	"github.com/go-devs-ua/octagon/app/globals"
	"github.com/go-devs-ua/octagon/lgr"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
)

func TestGetUserByID(t *testing.T) {
	logger, err := lgr.New("INFO")
	if err != nil {
		t.Fail()
	}

	tests := map[string]struct {
		id                   string
		expectedStatusCode   int
		expectedResponseBody string
		expectedUser         *entities.User
		expectedErr          error
	}{
		"success": {
			id:                 "a6703a6b-c054-4ebd-b782-2a96bd4f771a",
			expectedStatusCode: http.StatusOK,
			expectedResponseBody: `{ 
				"id": "a6703a6b-c054-4ebd-b782-2a96bd4f771a", 
				"first_name": "Jon", 
				"last_name":  "Doe", 
				"email":    "john@emasadfasilsdf.com", 
				"created_at": "2022-11-01T16:31:40.400825Z"	
			}`,
			expectedUser: &entities.User{
				ID:        "a6703a6b-c054-4ebd-b782-2a96bd4f771a",
				FirstName: "Jon",
				LastName:  "Doe",
				Email:     "john@emasadfasilsdf.com",
				CreatedAt: "2022-11-01T16:31:40.400825Z"},
			expectedErr: nil,
		},
		"invalid_short_id": {
			id:                 "1234",
			expectedStatusCode: http.StatusBadRequest,
			expectedResponseBody: `{
				"message": "Bad request", 
				"details": "invalid UUID length: 4"
			}`,
			expectedUser: nil,
			expectedErr:  nil,
		},
		"invalid_long_id": {
			id:                 "a6703a6b-c054-4ebd-b782-2a96bd4fasfd771asdfadfsfadsaa6703a6b-c054-4ebd-b782-2a96bd4f771aa6703a6b-c054-4ebd-b782-2a96bd4f771aa670",
			expectedStatusCode: http.StatusBadRequest,
			expectedResponseBody: `{
				"message": "Bad request", 
				"details": "invalid UUID length: 128"
			}`,
			expectedUser: nil,
			expectedErr:  nil,
		},
		"zero_id": {
			id:                 "",
			expectedStatusCode: http.StatusBadRequest,
			expectedResponseBody: `{ 
				"message": "Bad request", 
				"details": "invalid UUID length: 0" 
			}`,
			expectedUser: nil,
			expectedErr:  nil,
		},
		"wrong_id": {
			id:                 "a6703a6b-c054-4ebd-b782-2a96bd4f771a",
			expectedStatusCode: http.StatusNotFound,
			expectedResponseBody: `{
				"message": "Not found",
				"details": "no user found in DB"
			} `,
			expectedUser: nil,
			expectedErr:  globals.ErrNotFound,
		},
		"unexpected_error": {
			id:                 "a6703a6b-c054-4ebd-b782-2a96bd4f771a",
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponseBody: `{ 
				"message": "Internal server error", 
				"details": ""
			}`,
			expectedUser: nil,
			expectedErr:  errors.New("unexpected error"),
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			usecase := NewMockUserUsecase(ctrl)
			usecase.EXPECT().
				GetByID(gomock.Eq(tt.id)).
				Return(tt.expectedUser, tt.expectedErr).
				Times(func() int {
					if tt.expectedUser == nil && tt.expectedErr == nil {
						return 0
					}
					return 1
				}())

			uh := UserHandler{
				usecase: usecase,
				logger:  logger,
			}

			resp := httptest.NewRecorder()
			defer func() {
				if err := resp.Result().Body.Close(); err != nil {
					t.Logf("Failed closing response body: %v", err)
				}
			}()

			req := httptest.NewRequest(http.MethodGet, path.Join("/users", tt.id), nil)
			req = mux.SetURLVars(req, map[string]string{"id": tt.id})

			uh.GetUserByID(resp, req)

			require.Equal(t, resp.Code, tt.expectedStatusCode)

			require.JSONEq(t, tt.expectedResponseBody, resp.Body.String())
		})
	}
}
