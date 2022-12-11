package data

import (
	"context"

	"github.com/aligntzy/rubick/app/geo/internal/biz"
	"github.com/aligntzy/rubick/app/geo/internal/pkg/location"
	"github.com/go-kratos/kratos/v2/log"
)

type geoRepo struct {
	data *Data
	log  *log.Helper
}

func NewGEORepo(data *Data, logger log.Logger) biz.GEORepo {
	return &geoRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *geoRepo) Location(ctx context.Context, ip string) (*biz.Location, error) {
	res, err := location.Location(ip)
	if err != nil {
		return nil, err
	}
	return &biz.Location{
		IP:      res.Data.IP,
		Country: res.Data.Country,
		City:    res.Data.City,
	}, nil
}
