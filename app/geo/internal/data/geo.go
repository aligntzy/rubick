package data

import (
	"github.com/aligntzy/rubick/app/geo/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type geoRepo struct {
	data *Data
	log  *log.Helper
}

func NewGEORepo(data *Data, logger log.Logger) biz.GEORepo {
	return geoRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
