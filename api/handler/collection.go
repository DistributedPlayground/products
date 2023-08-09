package handler

import (
	"github.com/DistributedPlayground/products/pkg/service"
	"github.com/labstack/echo/v4"
)

type Collection interface {
	Create(c echo.Context) error
	Update(c echo.Context) error
	RegisterRoutes(g *echo.Group, ms ...echo.MiddlewareFunc)
}

type collection struct {
	service service.Collection
	Group   *echo.Group
}

func NewCollection(service service.Collection) Collection {
	return &collection{service: service}
}
