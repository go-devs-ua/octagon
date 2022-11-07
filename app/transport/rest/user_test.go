package rest

import (
	"net/http"
	"net/http/httptest"
	"path"
	"testing"

	"github.com/go-devs-ua/octagon/app/entities"
	"github.com/go-devs-ua/octagon/lgr"
	"github.com/stretchr/testify/assert"

	"github.com/gorilla/mux"
)

func TestGetUserByID(t *testing.T) {
	logger, err := lgr.New("INFO")
	if err != nil {
		t.Fail()
	}

	tests := map[string]struct {
		id                   string
		expectedStatusCode   int
		expectedErr          error
		expectedResponseBody string
		getByIDFunc          UserUsecase
	}{
		"success": {
			id:                 "a6703a6b-c054-4ebd-b782-2a96bd4f771a",
			expectedStatusCode: http.StatusOK,
			expectedErr:        nil,
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
						CreatedAt: "2022-11-01T16:31:40.400825Z",
					}, nil
				}},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			uh := UserHandler{
				usecase: tt.getByIDFunc,
				logger:  logger,
			}

			rec := httptest.NewRecorder()
			resp := rec.Result()
			defer resp.Body.Close()

			req := httptest.NewRequest(http.MethodGet, path.Join("/users", tt.id), nil)
			req = mux.SetURLVars(req, map[string]string{"id": tt.id})

			uh.GetUserByID(rec, req)

			assert.Equal(t, resp.StatusCode, tt.expectedStatusCode)

			assert.JSONEq(t, tt.expectedResponseBody, rec.Body.String())
		})
	}
}
