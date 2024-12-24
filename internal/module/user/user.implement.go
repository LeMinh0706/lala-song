package user

import (
	"context"

	"github.com/LeMinh0706/lala-song/internal/db"
	"github.com/LeMinh0706/lala-song/util"
)

type UserService struct {
	q *db.Queries
}

// Login implements IUserService.
func (u *UserService) Login(ctx context.Context, username string, password string) (db.User, error) {
	var res db.User
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
func (u *UserService) Register(ctx context.Context, arg db.RegisterParams) (db.User, error) {
	panic("unimplemented")
}

func NewUserService(q *db.Queries) IUserService {
	return &UserService{
		q: q,
	}
}
