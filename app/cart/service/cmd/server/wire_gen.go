// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/go-kratos/beer-shop/app/cart/service/internal/biz"
	"github.com/go-kratos/beer-shop/app/cart/service/internal/conf"
	"github.com/go-kratos/beer-shop/app/cart/service/internal/data"
	"github.com/go-kratos/beer-shop/app/cart/service/internal/server"
	"github.com/go-kratos/beer-shop/app/cart/service/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, error) {
	dataData, err := data.NewData(confData, logger)
	if err != nil {
		return nil, err
	}
	cartRepo := data.NewCartRepo(dataData, logger)
	cartUseCase := biz.NewCartUseCase(cartRepo, logger)
	cartService := service.NewCartService(cartUseCase, logger)
	httpServer := server.NewHTTPServer(confServer, cartService)
	grpcServer := server.NewGRPCServer(confServer, cartService)
	registrar := server.NewRegistrar()
	app := newApp(logger, httpServer, grpcServer, registrar)
	return app, nil
}
