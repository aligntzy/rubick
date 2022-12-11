package service

import (
	"context"

	pb "github.com/aligntzy/rubick/api/geo/v1"
	"github.com/aligntzy/rubick/app/geo/internal/biz"
)

type GEOService struct {
	pb.UnimplementedGEOServer

	uc *biz.GEOUseCase
}

func NewGEOService(uc *biz.GEOUseCase) *GEOService {
	return &GEOService{uc: uc}
}

func (s *GEOService) Location(ctx context.Context, req *pb.LocationRequest) (*pb.LocationReply, error) {
	return &pb.LocationReply{}, nil
}
