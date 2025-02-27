package db

import (
	"context"
	"database/sql"
	"testing"

	// "github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	arg := CreateUserParams{
		FirstName:   sql.NullString{String: "Nguyen", Valid: true},
		LastName:    sql.NullString{String: "Nguyen", Valid: true},
		Email:       "hnguy5@gmail.com",
		Password:    "123n123n",
		PhoneNumber: sql.NullString{String: "123456789", Valid: true},
		IsActive:    true,
	}

	user, err := testQueries.CreateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.NotZero(t, user.ID)
	require.Equal(t, arg.FirstName, user.FirstName)
	require.Equal(t, arg.LastName, user.LastName)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.Password, user.Password)
	require.Equal(t, arg.PhoneNumber, user.PhoneNumber)
	require.Equal(t, arg.IsActive, user.IsActive)
	require.NotZero(t, user.UpdatedAt)
	require.NotZero(t, user.CreatedAt)

}
