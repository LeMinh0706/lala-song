package test

import (
	"context"
	"testing"

	"github.com/LeMinh0706/lala-song/internal/db"
	"github.com/LeMinh0706/lala-song/util"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) {
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

func createSinger(t *testing.T) {
	singer, err := testQueries.CreateSinger(context.Background(), db.CreateSingerParams{
		Fullname: util.RandomString(6),
		ImageUrl: util.RandomAvatar(util.RandomGender()),
	})
	require.NoError(t, err)
	require.NotEmpty(t, singer)
}

// // Testing
func TestRegister(t *testing.T) {
	createRandomUser(t)
}

func TestSinger(t *testing.T) {
	createSinger(t)
}

func TestSelect(t *testing.T) {
	for i := 0; i < 10; i++ {
		createSinger(t)
	}
	list, err := testQueries.GetSinger(context.Background(), db.GetSingerParams{Limit: 10, Offset: 3})
	require.NoError(t, err)
	require.NotEmpty(t, list)
}
