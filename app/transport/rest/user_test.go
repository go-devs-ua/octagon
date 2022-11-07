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
		getByIDFunc          UserUsecase
	}{
		"success": {
			id:                 "a6703a6b-c054-4ebd-b782-2a96bd4f771a",
			expectedStatusCode: http.StatusOK,
			expectedResponseBody: `
									{
										"id": "a6703a6b-c054-4ebd-b782-2a96bd4f771a",
										"first_name": "Jon",
										"last_name":  "Doe",
										"email":    "john@emasadfasilsdf.com",
										"created_at": "2022-11-01T16:31:40.400825Z"	
									}
								   `,
			getByIDFunc: &UserUsecaseMock{
				GetByIDFunc: func(string) (*entities.User, error) {
					return &entities.User{
						ID:        "a6703a6b-c054-4ebd-b782-2a96bd4f771a",
						FirstName: "Jon",
						LastName:  "Doe",
						Email:     "john@emasadfasilsdf.com",
						CreatedAt: "2022-11-01T16:31:40.400825Z"}, nil
				}},
		},
		"invalid_short_id": {
			id:                 "1234",
			expectedStatusCode: http.StatusBadRequest,
			expectedResponseBody: `
									{
										"message": "Bad request",
										"details": "invalid UUID length: 4"
									}
								   `,
			getByIDFunc: nil,
		},
		"invalid_long_id": {
			id:                 "a6703a6b-c054-4ebd-b782-2a96bd4fasfd771asdfadfsfadsaa6703a6b-c054-4ebd-b782-2a96bd4f771aa6703a6b-c054-4ebd-b782-2a96bd4f771aa670",
			expectedStatusCode: http.StatusBadRequest,
			expectedResponseBody: `
									{
										"message": "Bad request",
										"details": "invalid UUID length: 128"
									}
								   `,
			getByIDFunc: nil,
		},
		"zero_id": {
			id:                 "",
			expectedStatusCode: http.StatusBadRequest,
			expectedResponseBody: `
									{
										"message": "Bad request",
										"details": "invalid UUID length: 0"
									}
								   `,
			getByIDFunc: nil,
		},
		"wrong_id": {
			id:                 "a6703a6b-c054-4ebd-b782-2a96bd4f771a",
			expectedStatusCode: http.StatusNotFound,
			expectedResponseBody: `
									{
										"message": "Not found",
										"details": "no user found in DB"
									}
								   `,
			getByIDFunc: &UserUsecaseMock{
				GetByIDFunc: func(string) (*entities.User, error) {
					return nil, globals.ErrNotFound
				}},
		},
		"unexpected_error": {
			id:                 "a6703a6b-c054-4ebd-b782-2a96bd4f771a",
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponseBody: `
									{
										"message": "Internal server error",
										"details": ""
									}
								   `,
			getByIDFunc: &UserUsecaseMock{
				GetByIDFunc: func(string) (*entities.User, error) {
					return nil, errors.New("unexpected error")
				}},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			uh := UserHandler{
				usecase: tt.getByIDFunc,
				logger:  logger,
			}

			resp := httptest.NewRecorder()
			defer resp.Result().Body.Close()

			req := httptest.NewRequest(http.MethodGet, path.Join("/users", tt.id), nil)
			req = mux.SetURLVars(req, map[string]string{"id": tt.id})

			uh.GetUserByID(resp, req)

			require.Equal(t, resp.Code, tt.expectedStatusCode)

			require.JSONEq(t, tt.expectedResponseBody, resp.Body.String())
		})
	}
}
