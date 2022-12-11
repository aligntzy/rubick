package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type Location struct {
	IP      string
	Country string
	City    string
}

type GEORepo interface {
	Location(ctx context.Context, ip string) (*Location, error)
}

type GEOUseCase struct {
	repo GEORepo
	log  *log.Helper
}

func NewGEOUseCase(repo GEORepo, logger log.Logger) *GEOUseCase {
	return &GEOUseCase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (uc *GEOUseCase) Location(ctx context.Context, ip string) (*Location, error) {
	return uc.repo.Location(ctx, ip)
}
