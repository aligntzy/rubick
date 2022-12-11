package biz

import "github.com/go-kratos/kratos/v2/log"

type Location struct{}

type GEORepo interface{}

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
