package usecase

import (
	"testing"

	"github.com/go-devs-ua/octagon/app/entities"
	"github.com/go-devs-ua/octagon/app/globals"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestUser_GetByID(t *testing.T) {
	tests := map[string]struct {
		id           string
		expectedUser *entities.User
		expectedErr  error
	}{
		"success": {
			id: "a6703a6b-c054-4ebd-b782-2a96bd4f771a",
			expectedUser: &entities.User{
				ID:        "a6703a6b-c054-4ebd-b782-2a96bd4f771a",
				FirstName: "Jon",
				LastName:  "Doe",
				Email:     "john@emasadfasilsdf.com",
				CreatedAt: "2022-11-01T16:31:40.400825Z"},
			expectedErr: nil,
		},
		"fail": {
			id:           "a6703a6b-c054-4ebd-b782-2a96bd4f771b",
			expectedUser: nil,
			expectedErr:  globals.ErrNotFound,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			usecase := NewMockRepository(ctrl)
			usecase.EXPECT().
				FindUser(gomock.Eq(tt.id)).
				Return(tt.expectedUser, tt.expectedErr).
				Times(1)

			repo := User{Repo: usecase}

			user, err := repo.GetByID(tt.id)
			require.Equal(t, tt.expectedUser, user)
			require.ErrorIs(t, err, tt.expectedErr)
		})
	}
}
