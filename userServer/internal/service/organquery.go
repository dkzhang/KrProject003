package service

import (
	"context"

	pb "userServer/api/userApi/v1"
)

type OrganQueryService struct {
	pb.UnimplementedOrganQueryServer
}

func NewOrganQueryService() *OrganQueryService {
	return &OrganQueryService{}
}

func (s *OrganQueryService) GetUser(ctx context.Context, req *pb.GetOrganRequest) (*pb.GetOrganReply, error) {
	return &pb.GetOrganReply{}, nil
}
