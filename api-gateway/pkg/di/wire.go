//go:build wireinject
// +build wireinject

package di

import (
	"videogateway/pkg/api"
	"videogateway/pkg/api/handler"
	"videogateway/pkg/client"
	"videogateway/pkg/config"

	"github.com/google/wire"
)

func InitializeAPI(c *config.Config) (*api.Server, error) {
	wire.Build(client.InitClient, client.NewVideoClient, handler.NewVideoHandler, api.NewServeHTTP)
	return &api.Server{}, nil
}
