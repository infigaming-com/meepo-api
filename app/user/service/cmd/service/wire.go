//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/infigaming-com/meepo-api/app/user/service/internal/biz"
	"github.com/infigaming-com/meepo-api/app/user/service/internal/conf"
	"github.com/infigaming-com/meepo-api/app/user/service/internal/data"
	"github.com/infigaming-com/meepo-api/app/user/service/internal/server"
	"github.com/infigaming-com/meepo-api/app/user/service/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Biz, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
