package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/infigaming-com/meepo-api/app/user/service/internal/biz"
)

var _ biz.UserRepo = (*userRepo)(nil)

type userRepo struct {
	lg   *log.Helper
	data *Data
}

func NewUserRepo(lg log.Logger, data *Data) biz.UserRepo {
	return &userRepo{
		lg:   log.NewHelper(log.With(lg, "module", "data/user")),
		data: data,
	}
}

func (r *userRepo) Create(ctx context.Context, u *biz.User) (*biz.User, error) {
	r.lg.WithContext(ctx).Infof("Create: %v", u)
	return u, nil
}

func (r *userRepo) Update(ctx context.Context, u *biz.User) (*biz.User, error) {
	r.lg.WithContext(ctx).Infof("Update: %v", u)
	return u, nil
}

func (r *userRepo) GetById(ctx context.Context, id string) (*biz.User, error) {
	r.lg.WithContext(ctx).Infof("GetById: %v", id)
	return nil, nil
}
