package service

import (
	"context"

	pb "github.com/infigaming-com/meepo-api/api/reward/service/v1"
)

type RewardService struct {
	pb.UnimplementedRewardServer
}

func NewRewardService() *RewardService {
	return &RewardService{}
}

func (s *RewardService) CreateReward(ctx context.Context, req *pb.CreateRewardRequest) (*pb.CreateRewardReply, error) {
	return &pb.CreateRewardReply{}, nil
}
func (s *RewardService) UpdateReward(ctx context.Context, req *pb.UpdateRewardRequest) (*pb.UpdateRewardReply, error) {
	return &pb.UpdateRewardReply{}, nil
}
func (s *RewardService) DeleteReward(ctx context.Context, req *pb.DeleteRewardRequest) (*pb.DeleteRewardReply, error) {
	return &pb.DeleteRewardReply{}, nil
}
func (s *RewardService) GetReward(ctx context.Context, req *pb.GetRewardRequest) (*pb.GetRewardReply, error) {
	return &pb.GetRewardReply{}, nil
}
func (s *RewardService) ListReward(ctx context.Context, req *pb.ListRewardRequest) (*pb.ListRewardReply, error) {
	return &pb.ListRewardReply{}, nil
}
