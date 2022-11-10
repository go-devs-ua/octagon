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
		getByIdFunc          func(*gomock.Controller, string) *MockUserUsecase
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
			getByIdFunc: func(ctrl *gomock.Controller, id string) *MockUserUsecase {
				mock := NewMockUserUsecase(ctrl)
				mock.EXPECT().
					GetByID(gomock.Eq(id)).
					Return(&entities.User{
						ID:        "a6703a6b-c054-4ebd-b782-2a96bd4f771a",
						FirstName: "Jon",
						LastName:  "Doe",
						Email:     "john@emasadfasilsdf.com",
						CreatedAt: "2022-11-01T16:31:40.400825Z"}, nil).
					Times(1)

				return mock
			},
		},
		"invalid_short_id": {
			id:                 "1234",
			expectedStatusCode: http.StatusBadRequest,
			expectedResponseBody: `{
				"message": "Bad request", 
				"details": "invalid UUID length: 4"
			}`,
			getByIdFunc: func(ctrl *gomock.Controller, id string) *MockUserUsecase {
				return nil
			},
		},
		"invalid_long_id": {
			id:                 "a6703a6b-c054-4ebd-b782-2a96bd4fasfd771asdfadfsfadsaa6703a6b-c054-4ebd-b782-2a96bd4f771aa6703a6b-c054-4ebd-b782-2a96bd4f771aa670",
			expectedStatusCode: http.StatusBadRequest,
			expectedResponseBody: `{
				"message": "Bad request", 
				"details": "invalid UUID length: 128"
			}`,
			getByIdFunc: func(ctrl *gomock.Controller, id string) *MockUserUsecase {
				return nil
			},
		},
		"zero_id": {
			id:                 "",
			expectedStatusCode: http.StatusBadRequest,
			expectedResponseBody: `{ 
				"message": "Bad request", 
				"details": "invalid UUID length: 0" 
			}`,
			getByIdFunc: func(ctrl *gomock.Controller, id string) *MockUserUsecase {
				return nil
			},
		},
		"wrong_id": {
			id:                 "a6703a6b-c054-4ebd-b782-2a96bd4f771a",
			expectedStatusCode: http.StatusNotFound,
			expectedResponseBody: `{
				"message": "Not found",
				"details": "no user found in DB"
			} `,
			getByIdFunc: func(ctrl *gomock.Controller, id string) *MockUserUsecase {
				mock := NewMockUserUsecase(ctrl)
				mock.EXPECT().
					GetByID(gomock.Eq(id)).
					Return(nil, globals.ErrNotFound)

				return mock
			},
		},
		"unexpected_error": {
			id:                 "a6703a6b-c054-4ebd-b782-2a96bd4f771a",
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponseBody: `{ 
				"message": "Internal server error", 
				"details": ""
			}`,
			getByIdFunc: func(ctrl *gomock.Controller, id string) *MockUserUsecase {
				mock := NewMockUserUsecase(ctrl)
				mock.EXPECT().
					GetByID(gomock.Eq(id)).
					Return(nil, errors.New("unexpected error"))

				return mock
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			uh := UserHandler{
				usecase: tt.getByIdFunc(ctrl, tt.id),
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
