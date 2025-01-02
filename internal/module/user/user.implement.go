package user

import (
	"context"

	"github.com/LeMinh0706/lala-song/internal/db"
	"github.com/LeMinh0706/lala-song/util"
	"github.com/google/uuid"
)

type UserService struct {
	q *db.Queries
}

// GetMe implements IUserService.
func (u *UserService) GetMe(ctx context.Context, username string) (db.GetMeRow, error) {
	user, err := u.q.GetMe(ctx, username)
	if err != nil {
		return db.GetMeRow{}, err
	}
	return user, nil
}

// Login implements IUserService.
func (u *UserService) Login(ctx context.Context, username string, password string) (db.LoginRow, error) {
	var res db.LoginRow
	user, err := u.q.Login(ctx, username)
	if err != nil {
		return res, err
	}
	err = util.CheckPassword(password, user.Password)
	if err != nil {
		return res, err
	}
	return user, nil
}

// Register implements IUserService.
func (u *UserService) Register(ctx context.Context, req Register) (db.User, error) {
	var res db.User
	tokenId, _ := uuid.NewRandom()
	hash, _ := util.HashPassword(req.Password)
	arg := db.RegisterParams{
		ID:       tokenId,
		Username: req.Username,
		Password: hash,
		Fullname: req.Fullname,
		Gender:   req.Gender,
		Avt:      util.RandomAvatar(req.Gender),
	}
	user, err := u.q.Register(ctx, arg)
	if err != nil {
		return res, err
	}
	return user, nil
}

func NewUserService(q *db.Queries) IUserService {
	return &UserService{
		q: q,
	}
}
