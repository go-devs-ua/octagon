package pg

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-devs-ua/octagon/app/entities"
	"github.com/stretchr/testify/require"
)

// Move const to separate file
func TestRepo_DeleteUser(t *testing.T) {
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
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("failed opening database connection %v", err)
			}
			defer db.Close()

			const SQL = `
			SELECT id, first_name, last_name, email, created_at 
			FROM "user" 
			WHERE id = (.+) 
			AND deleted_at is null;`

			rows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "email", "created_at"}).
				AddRow(
					tt.expectedUser.ID,
					tt.expectedUser.FirstName,
					tt.expectedUser.LastName,
					tt.expectedUser.Email,
					tt.expectedUser.CreatedAt)
			mock.ExpectQuery(SQL).
				WithArgs(tt.id).
				WillReturnRows(rows)

			r := Repo{DB: db}

			user, err := r.FindUser(tt.id)
			require.Equal(t, err, tt.expectedErr)
			require.Equal(t, tt.expectedUser, user)

		})
	}
}
