package rest

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-devs-ua/octagon/app/entities"
	"github.com/go-devs-ua/octagon/app/globals"
	"github.com/go-devs-ua/octagon/lgr"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestUserHandler_DeleteUser(t *testing.T) {
	logger, err := lgr.New("INFO")
	if err != nil {
		t.FailNow()
	}

	tests := map[string]struct {
		requestBody        string
		usecaseConstructor func(ctrl *gomock.Controller) UserUsecase
		expStatusCode      int
		expResponseBody    string
	}{
		"success": {
			requestBody: `{"id": "dca5947d-3dfc-49f1-bc09-dd53ce7e71cc"}`,
			usecaseConstructor: func(ctrl *gomock.Controller) UserUsecase {
				mock := NewMockUserUsecase(ctrl)

				mock.EXPECT().Delete(entities.User{
					ID: "dca5947d-3dfc-49f1-bc09-dd53ce7e71cc",
				}).Return(nil).Times(1)

				return mock
			},
			expResponseBody: "",
			expStatusCode:   http.StatusNoContent,
		},
		"invalid_notexisted_id": {
			requestBody: `{"id": "dca5947d-3dfc-49f1-bc09-dd53ce7e71cc"}`,
			usecaseConstructor: func(ctrl *gomock.Controller) UserUsecase {
				mock := NewMockUserUsecase(ctrl)
				mock.EXPECT().Delete(entities.User{
					ID: "dca5947d-3dfc-49f1-bc09-dd53ce7e71cc",
				}).Return(globals.ErrNotFound).Times(1)

				return mock
			},
			expResponseBody: `{"message":"Not found","details":"no user found in DB"}`,
			expStatusCode:   http.StatusNotFound,
		},
		"invalid_bad_id_too_long": {
			requestBody: `{"id": "1acef78b-55ef-43b1-8c2b-c9e36c9c11111"}`,
			usecaseConstructor: func(ctrl *gomock.Controller) UserUsecase {
				return nil
			},
			expResponseBody: `{"message":"Bad request","details":"invalid uuid: invalid UUID length: 37"}`,
			expStatusCode:   http.StatusBadRequest,
		},
		"invalid_bad_id_too_short": {
			requestBody: `{"id": "1acef78b-55ef-43b1-8c2b-c9e36c9c280"}`,
			usecaseConstructor: func(ctrl *gomock.Controller) UserUsecase {
				return nil
			},
			expResponseBody: `{"message":"Bad request","details":"invalid uuid: invalid UUID length: 35"}`,
			expStatusCode:   http.StatusBadRequest,
		},
		"invalid_bad_id_bad_symbpls": {
			requestBody: `{"id": "%acef78b-55ef-43b1-8c2b-c9e36c9c2807"}`,
			usecaseConstructor: func(ctrl *gomock.Controller) UserUsecase {
				return nil
			},
			expResponseBody: `{"message":"Bad request","details":"invalid uuid: invalid UUID format"}`,
			expStatusCode:   http.StatusBadRequest,
		},
		"invalid_internal_server_error": {
			requestBody: `{"id": "dca5947d-3dfc-49f1-bc09-dd53ce7e71cc"}`,
			usecaseConstructor: func(ctrl *gomock.Controller) UserUsecase {
				mock := NewMockUserUsecase(ctrl)
				mock.EXPECT().Delete(entities.User{
					ID: "dca5947d-3dfc-49f1-bc09-dd53ce7e71cc",
				}).Return(errors.New("Internal error")).Times(1)

				return mock
			},
			expResponseBody: `{"message":"Internal server error","details":""}`,
			expStatusCode:   http.StatusInternalServerError,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			uh := UserHandler{
				usecase: tt.usecaseConstructor(ctrl),
				logger:  logger,
			}
			defer ctrl.Finish()
			resp := httptest.NewRecorder()
			uh.DeleteUser(resp, httptest.NewRequest(http.MethodDelete, "*", strings.NewReader(tt.requestBody)))
			require.Equal(t, tt.expStatusCode, resp.Code)
			if len(tt.expResponseBody) > 0 {
				require.JSONEq(t, tt.expResponseBody, resp.Body.String())
			}
		})
	}
}
