package data

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/infigaming-com/go-common/lock"
	"github.com/infigaming-com/meepo-api/app/user/service/internal/biz"
)

var _ biz.UserLoginRepo = (*userLoginRepo)(nil)

type userLoginRepo struct {
	lg   *log.Helper
	data *Data
}

func NewUserLoginRepo(lg log.Logger, data *Data) biz.UserLoginRepo {
	return &userLoginRepo{
		lg:   log.NewHelper(log.With(lg, "module", "data/user_login")),
		data: data,
	}
}

func (r *userLoginRepo) Lock(ctx context.Context, ul *biz.UserLogin, timeout time.Duration) (func(), error) {
	r.lg.WithContext(ctx).Infof("Lock: %v, timeout: %s", ul, timeout)
	unlocker, err := r.data.lock.Lock(ctx, ul.LockerKey(), lock.WithTimeout(timeout))
	if err != nil {
		return nil, err
	}
	return func() {
		unlocker(ctx)
	}, nil
}

func (r *userLoginRepo) Create(ctx context.Context, ul *biz.UserLogin) (*biz.UserLogin, error) {
	r.lg.WithContext(ctx).Infof("Create: %v", ul)
	return ul, nil
}

func (r *userLoginRepo) GetByLoginTypeAndLoginId(ctx context.Context, loginType string, loginId string) (*biz.UserLogin, error) {
	r.lg.WithContext(ctx).Infof("GetByLoginTypeAndLoginId: %s, %s", loginType, loginId)
	return nil, nil
}
