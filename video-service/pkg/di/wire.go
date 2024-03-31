//go:build wireinject
// +build wireinject

package di

import (
	"videoservice/pkg/api"
	"videoservice/pkg/api/service"
	"videoservice/pkg/config"
	"videoservice/pkg/db"
	"videoservice/pkg/repository"

	"github.com/google/wire"
)

func InitializeServe(c *config.Config) (*api.Server, error) {
	wire.Build(db.Initdb, repository.NewVideoRepo, service.NewVideoServer, api.NewgrpcServe)
	return &api.Server{}, nil
}
