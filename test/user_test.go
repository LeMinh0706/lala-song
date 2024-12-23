package test

import (
	"context"
	"testing"

	"github.com/LeMinh0706/lala-song/internal/db"
	"github.com/LeMinh0706/lala-song/util"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestRegister(t *testing.T) {
	hash, err := util.HashPassword("kocanpass")
	require.NoError(t, err)
	gender := util.RandomGender()
	id, _ := uuid.NewRandom()

	require.NoError(t, err)
	user, err := testQueries.Register(context.Background(), db.RegisterParams{
		ID:       id,
		Username: util.RandomString(6),
		Password: hash,
		Fullname: util.RandomString(10),
		Gender:   gender,
		Avt:      util.RandomAvatar(gender),
	})
	require.NoError(t, err)
	require.NotEmpty(t, user)
}
