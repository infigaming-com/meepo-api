//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/infigaming-com/meepo-api/internal/biz"
	"github.com/infigaming-com/meepo-api/internal/conf"
	"github.com/infigaming-com/meepo-api/internal/data"
	"github.com/infigaming-com/meepo-api/internal/server"
	"github.com/infigaming-com/meepo-api/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
