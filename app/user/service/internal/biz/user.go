package biz

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/infigaming-com/meepo-api/app/user/service/internal/conf"
	"github.com/infigaming-com/meepo-api/app/user/service/internal/pkg/jwt"
	"github.com/infigaming-com/meepo-api/app/user/service/internal/pkg/util"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUserNotFound          = errors.New("user not found")
	ErrUserAlreadyRegistered = errors.New("user already registered")
	ErrInvalidPassword       = errors.New("invalid password")
)

const (
	LoginTypeEmail  = "email"
	LoginTypeMobile = "mobile"
	LoginTypeGoogle = "google"
)

type UserLogin struct {
	LoginType string
	LoginId   string
	UserId    string
	LoginData any
}

type UserLoginDataEmail struct {
	Password string
}

type UserLoginDataMobile struct{}

type UserLoginDataGoogle struct {
	GoogleId string
}

type UserLoginInfo struct {
	UserId      string
	AccessToken string
}

type User struct {
	Id          string
	OperatorId  string
	Username    string
	Enabled     bool
	Email       string
	Firstname   string
	Lastname    string
	Nickname    string
	Country     string
	DateOfBirth string
	BanLogin    bool
	BanGame     bool
	BanWithdraw bool
	MfaEnabled  bool
	MfsSecret   string
}

type UserLoginRepo interface {
	Lock(context.Context, *UserLogin, time.Duration) (func(), error)
	Create(context.Context, *UserLogin) (*UserLogin, error)
	GetByLoginTypeAndLoginId(context.Context, string, string) (*UserLogin, error)
}

type UserRepo interface {
	Create(context.Context, *User) (*User, error)
	Update(context.Context, *User) (*User, error)
	GetById(context.Context, string) (*User, error)
}

type UserUsecase struct {
	lg            *log.Helper
	userLoginRepo UserLoginRepo
	userRepo      UserRepo
	jwtSecret     string
	jwtExpiry     time.Duration
}

func NewUserUsecase(c *conf.Biz, lg log.Logger, userLoginRepo UserLoginRepo, userRepo UserRepo) *UserUsecase {
	return &UserUsecase{
		lg:            log.NewHelper(log.With(lg, "module", "biz/user")),
		userLoginRepo: userLoginRepo,
		userRepo:      userRepo,
		jwtSecret:     c.User.JwtSecret,
		jwtExpiry:     c.User.JwtExpiry.AsDuration(),
	}
}

func (uc *UserUsecase) Register(ctx context.Context, email, password string) (*UserLoginInfo, error) {
	uc.lg.WithContext(ctx).Infof("Register: %v", email)

	existingUserLogin, err := uc.userLoginRepo.GetByLoginTypeAndLoginId(ctx, LoginTypeEmail, email)
	if err != nil {
		return nil, err
	}

	if existingUserLogin != nil {
		return nil, ErrUserAlreadyRegistered
	}

	unlock, err := uc.userLoginRepo.Lock(
		ctx,
		&UserLogin{
			LoginType: LoginTypeEmail,
			LoginId:   email,
		},
		10*time.Second,
	)
	if err != nil {
		return nil, err
	}
	defer unlock()

	existingUserLogin, err = uc.userLoginRepo.GetByLoginTypeAndLoginId(ctx, LoginTypeEmail, email)
	if err != nil {
		return nil, err
	}

	if existingUserLogin != nil {
		return nil, ErrUserAlreadyRegistered
	}

	hashedPassword, err := util.HashPassword(password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	userLogin := &UserLogin{
		LoginType: LoginTypeEmail,
		LoginId:   email,
		LoginData: &UserLoginDataEmail{
			Password: hashedPassword,
		},
	}

	userLogin, err = uc.userLoginRepo.Create(ctx, userLogin)
	if err != nil {
		return nil, fmt.Errorf("failed to create user login record: %w", err)
	}

	user := &User{
		Id:         userLogin.UserId,
		OperatorId: "", // TODO: operator_id should be included in context
		Email:      email,
		Enabled:    true,
	}
	user, err = uc.userRepo.Create(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("failed to create user record: %w", err)
	}

	accessToken, err := jwt.CreateToken(
		jwt.UserInfo{
			UserId: user.Id,
		},
		uc.jwtSecret,
		uc.jwtExpiry,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create access token: %w", err)
	}

	return &UserLoginInfo{
		UserId:      user.Id,
		AccessToken: accessToken,
	}, nil
}

func (uc *UserUsecase) Login(ctx context.Context, email, password string) (*UserLoginInfo, error) {
	uc.lg.WithContext(ctx).Infof("Login: %v", email)

	userLogin, err := uc.userLoginRepo.GetByLoginTypeAndLoginId(ctx, LoginTypeEmail, email)
	if err != nil {
		return nil, err
	}

	if userLogin == nil {
		return nil, ErrUserNotFound
	}

	userLoginDataEmail, ok := userLogin.LoginData.(*UserLoginDataEmail)
	if !ok {
		return nil, fmt.Errorf("invalid user login data: %v", userLogin.LoginData)
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(userLoginDataEmail.Password),
		[]byte(password),
	); err != nil {
		return nil, ErrInvalidPassword
	}

	accessToken, err := jwt.CreateToken(
		jwt.UserInfo{
			UserId: userLogin.UserId,
		},
		uc.jwtSecret,
		uc.jwtExpiry,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create access token: %w", err)
	}

	return &UserLoginInfo{
		UserId:      userLogin.UserId,
		AccessToken: accessToken,
	}, nil
}

func (uc *UserUsecase) UpdateUser(ctx context.Context, u *User) (*User, error) {
	uc.lg.WithContext(ctx).Infof("Update: %v", u)
	return uc.userRepo.Update(ctx, u)
}

func (uc *UserUsecase) GetUser(ctx context.Context, id string) (*User, error) {
	uc.lg.WithContext(ctx).Infof("Get: %s", id)
	return uc.userRepo.GetById(ctx, id)
}

func (ul *UserLogin) Key() string {
	return fmt.Sprintf("user_login:%s:%s", ul.LoginType, ul.LoginId)
}

func (ul *UserLogin) LockerKey() string {
	return fmt.Sprintf("locker:%s", ul.Key())
}
