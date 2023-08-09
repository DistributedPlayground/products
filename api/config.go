package api

import (
	"github.com/DistributedPlayground/products/pkg/repository"
	"github.com/DistributedPlayground/products/pkg/service"
)

func NewRepos(config APIConfig) repository.Repositories {
	return repository.Repositories{
		Collection: repository.NewCollection(config.DB),
		Product:    repository.NewProduct(config.DB),
	}
}

func NewServices(config APIConfig, repos repository.Repositories) service.Services {
	return service.Services{
		Collection: service.NewCollection(repos.Collection),
		Product:    service.NewProduct(repos.Product),
	}
}
