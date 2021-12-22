package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"userServer/internal/biz"

	pb "userServer/api/userApi/v1"
)

type UserQueryService struct {
	pb.UnimplementedUserQueryServer

	uc  *biz.GreeterUsecase
	log *log.Helper
}

func NewUserQueryService(uc *biz.GreeterUsecase, logger log.Logger) *UserQueryService {
	return &UserQueryService{uc: uc, log: log.NewHelper(logger)}
}

func (s *UserQueryService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	//return &pb.GetUserReply{}, nil
	s.log.WithContext(ctx).Infof("GetUser ID Received: %d", req.GetId())
	if req.Id == 3 {
		return &pb.GetUserReply{
			Name: "zhangsan",
		}, nil
	} else {
		return &pb.GetUserReply{
			Name: "lisi",
		}, nil
	}
}
