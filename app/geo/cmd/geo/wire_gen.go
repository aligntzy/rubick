// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/aligntzy/rubick/app/geo/internal/biz"
	"github.com/aligntzy/rubick/app/geo/internal/conf"
	"github.com/aligntzy/rubick/app/geo/internal/data"
	"github.com/aligntzy/rubick/app/geo/internal/server"
	"github.com/aligntzy/rubick/app/geo/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	geoRepo := data.NewGEORepo(dataData, logger)
	geoUseCase := biz.NewGEOUseCase(geoRepo, logger)
	geoService := service.NewGEOService(geoUseCase)
	grpcServer := server.NewGRPCServer(confServer, geoService, logger)
	httpServer := server.NewHTTPServer(confServer, geoService, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
