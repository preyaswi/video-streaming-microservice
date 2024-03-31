// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"videoservice/pkg/api"
	"videoservice/pkg/api/service"
	"videoservice/pkg/config"
	"videoservice/pkg/db"
	"videoservice/pkg/repository"
)

// Injectors from wire.go:

func InitializeServe(c *config.Config) (*api.Server, error) {
	gormDB, err := db.Initdb(c)
	if err != nil {
		return nil, err
	}
	videoRepo := repository.NewVideoRepo(gormDB)
	videoServiceServer := service.NewVideoServer(videoRepo)
	server, err := api.NewgrpcServe(c, videoServiceServer)
	if err != nil {
		return nil, err
	}
	return server, nil
}
