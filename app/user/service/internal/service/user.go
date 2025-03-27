package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/infigaming-com/meepo-api/api/user/service/v1"
	"github.com/infigaming-com/meepo-api/app/user/service/internal/biz"
)

type UserService struct {
	pb.UnimplementedUserServer
	log *log.Helper
	uc  *biz.UserUsecase
}

func NewUserService(lg log.Logger, uc *biz.UserUsecase) *UserService {
	return &UserService{
		log: log.NewHelper(log.With(lg, "module", "service/user")),
		uc:  uc,
	}
}

func (s *UserService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	s.log.Infof("Register: %v", req.GetEmail())
	userLoginInfo, err := s.uc.Register(ctx, req.GetEmail(), req.GetPassword())
	if err != nil {
		return nil, err
	}
	return &pb.RegisterResponse{
		AccessToken: userLoginInfo.AccessToken,
		UserInfo: &pb.UserInfo{
			UserId: userLoginInfo.UserId,
		},
	}, nil
}

func (s *UserService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	s.log.Infof("Login: %v", req.GetEmail())
	return &pb.LoginResponse{}, nil
}

func (s *UserService) OAuthLogin(ctx context.Context, req *pb.OAuthLoginRequest) (*pb.OAuthLoginResponse, error) {
	s.log.Infof("OAuthLogin: %v", req)
	return &pb.OAuthLoginResponse{}, nil
}

func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	s.log.Infof("GetUser: %v", req)
	return &pb.GetUserResponse{}, nil
}
