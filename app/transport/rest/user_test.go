package rest

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"path"
	"testing"

	"github.com/go-devs-ua/octagon/app/globals"
	"github.com/gorilla/mux"

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
	const (
		//If we use const (nameMask, mailMask,passMask) from entities.user in our expectedResponseBody we got an error
		//"JSON parsing error: invalid character 'p' in string escape code" (where chacter 'p' is some character after symbol \
		//in regex). That`s why we create and use escaped version of this const
		nameMask = `^[\\p{L}&\\s-\\\\'’.]{2,256}$`
		mailMask = `(?i)^(?:[a-z\\d!#$%&'*+/=?^_\\x60{|}~-]+(?:\\.[a-z\\d!#$%&'*+/=?^_\\x60{|}~-]+)*)@(?:(?:[a-z\\d](?:[a-z\\d-]*[a-z\\d])?\\.)+[a-z\\d](?:[a-z\\d-]*[a-z\\d])?)$`
		passMask = `^[[:graph:]]{8,256}$`
		//maybe for test cases with too long domain or local part in email we should create separate variable for requestBody?
		longLocalPartEmail       = "1234567890123456789012345678901234567890123456789012345678901234567890@gmail.com"
		lenOfLongLocalPartEmail  = "70"
		longDomainPartEmail      = "1234567890@12345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890.com"
		lenOfLongDomainPartEmail = "264"
	)

	logger, err := lgr.New("INFO")
	if err != nil {
		t.FailNow()
	}

	tests := map[string]struct {
		requestBody          string
		usecaseConstructor   func(ctrl *gomock.Controller) UserUsecase
		expectedStatusCode   int
		expectedResponseBody string
	}{
		"success": {
			requestBody: `{"email":"johngold@gmail.com","first_name":"John","last_name":"Gold","password":"123456789Qw"}`,
			usecaseConstructor: func(ctrl *gomock.Controller) UserUsecase {
				mock := NewMockUserUsecase(ctrl)

				mock.EXPECT().SignUp(entities.User{
					Email:     "johngold@gmail.com",
					FirstName: "John",
					LastName:  "Gold",
					Password:  "123456789Qw",
				}).Return("91e3dcf7-34a6-4646-bd37-383cc949da93", nil).Times(1)

				return mock
			},
			expectedStatusCode:   http.StatusCreated,
			expectedResponseBody: `{"id":"91e3dcf7-34a6-4646-bd37-383cc949da93"}`,
		},
		"failed_decoding": {
			requestBody: `{"email":"johngold@gmail.com","first_name":"John","last_name":"Gold","password":"123456789Qw}`,
			usecaseConstructor: func(ctrl *gomock.Controller) UserUsecase {
				return nil
			},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"message":"Bad request","details":"unexpected EOF"}`,
		},
		"first_name_validation_failed": {
			requestBody: `{"email":"johngold@gmail.com","first_name":"J","last_name":"Gold","password":"123456789Qw"}`,
			usecaseConstructor: func(ctrl *gomock.Controller) UserUsecase {

				return nil
			},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"message":"Bad request","details":"invalid first name: name does not match with regex: ` + "`" + nameMask + "`" + `"}`,
		},
		"last_name_validation_failed": {
			requestBody: `{"email":"johngold@gmail.com","first_name":"John","last_name":"**","password":"123456789Qw"}`,
			usecaseConstructor: func(ctrl *gomock.Controller) UserUsecase {

				return nil
			},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"message":"Bad request","details":"invalid last name: name does not match with regex: ` + "`" + nameMask + "`" + `"}`,
		},
		"too_long_local_part_of_email": {
			requestBody: `{"email":"` + longLocalPartEmail + `","first_name":"John","last_name":"Gold","password":"123456789Qw"}`,
			usecaseConstructor: func(ctrl *gomock.Controller) UserUsecase {

				return nil
			},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"message":"Bad request","details":"local part of email contains too many bytes: ` + lenOfLongLocalPartEmail + `"}`,
		},
		"too_long_domain_part_of_email": {
			requestBody: `{"email":"` + longDomainPartEmail + `","first_name":"John","last_name":"Gold","password":"123456789Qw"}`,
			usecaseConstructor: func(ctrl *gomock.Controller) UserUsecase {

				return nil
			},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"message":"Bad request","details":"domain part of email contains too many bytes: ` + lenOfLongDomainPartEmail + `"}`,
		},
		"email_does_not_match_with_regex": {
			requestBody: `{"email":"johngoldgmail.com","first_name":"John","last_name":"Gold","password":"123456789Qw"}`,
			usecaseConstructor: func(ctrl *gomock.Controller) UserUsecase {

				return nil
			},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"message":"Bad request","details":"email does not match with regex: ` + "`" + mailMask + "`" + `"}`,
		},
		"password_validation_failed": {
			requestBody: `{"email":"john@goldgmail.com","first_name":"John","last_name":"Gold","password":"gotcha!"}`,
			usecaseConstructor: func(ctrl *gomock.Controller) UserUsecase {

				return nil
			},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"message":"Bad request","details":"password does not match with regex: ` + "`" + passMask + "`" + `"}`,
		},
		"email_duplicate": {
			requestBody: `{"email":"johngold@gmail.com","first_name":"John","last_name":"Gold","password":"123456789Qw"}`,
			usecaseConstructor: func(ctrl *gomock.Controller) UserUsecase {
				mock := NewMockUserUsecase(ctrl)

				mock.EXPECT().SignUp(entities.User{
					Email:     "johngold@gmail.com",
					FirstName: "John",
					LastName:  "Gold",
					Password:  "123456789Qw",
				}).Return("", globals.ErrDuplicateEmail).Times(1)

				return mock
			},
			expectedStatusCode:   http.StatusConflict,
			expectedResponseBody: `{"message":"Bad request","details":"email is already taken"}`,
		},
		"internal_server_error": {
			requestBody: `{"email":"johngold@gmail.com","first_name":"John","last_name":"Gold","password":"123456789Qw"}`,
			usecaseConstructor: func(ctrl *gomock.Controller) UserUsecase {
				mock := NewMockUserUsecase(ctrl)

				mock.EXPECT().SignUp(entities.User{
					Email:     "johngold@gmail.com",
					FirstName: "John",
					LastName:  "Gold",
					Password:  "123456789Qw",
				}).Return("", errors.New("Internal error")).Times(1)

				return mock
			},
			expectedStatusCode:   http.StatusInternalServerError,
			expectedResponseBody: `{"message":"Internal server error","details":""}`,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			uh := UserHandler{
				usecase: tt.usecaseConstructor(ctrl),
				logger:  logger,
			}

			response := httptest.NewRecorder()
			request := httptest.NewRequest(http.MethodPost, "/users", nil)
			uh.CreateUser(response, request)
			require.Equal(t, tt.expectedStatusCode, response.Code)
			require.JSONEq(t, tt.expectedResponseBody, response.Body.String())
		})
	}
}
