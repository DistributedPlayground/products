package handler

import (
	"net/http"

	"github.com/DistributedPlayground/go-lib/common"
	"github.com/DistributedPlayground/go-lib/httperror"
	"github.com/DistributedPlayground/products/pkg/model"
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

func (c collection) Create(ctx echo.Context) error {
	body := model.CollectionUpsert{}
	err := ctx.Bind(&body)
	if err != nil {
		common.LogDPError(ctx, err, "Collection Create: Bad Request")
		return httperror.BadRequest400(ctx)
	}

	collection, err := c.service.Create(ctx.Request().Context(), body)
	if err != nil {
		common.LogDPError(ctx, err, "Collection Create: Internal Error")
		return httperror.Internal500(ctx)
	}
	return ctx.JSON(http.StatusCreated, collection)
}

func (c collection) Update(ctx echo.Context) error {
	body := model.CollectionUpsert{}
	err := ctx.Bind(&body)
	if err != nil {
		common.LogDPError(ctx, err, "Collection Update: Bad Request")
		return httperror.BadRequest400(ctx)
	}

	collection, err := c.service.Update(ctx.Request().Context(), ctx.Param("id"), body)
	if err != nil {
		common.LogDPError(ctx, err, "Collection Update: Internal Error")
		return httperror.Internal500(ctx)
	}
	return ctx.JSON(http.StatusOK, collection)
}

func (c collection) RegisterRoutes(g *echo.Group, ms ...echo.MiddlewareFunc) {
	if g == nil {
		panic("no group attached to the collection handler")
	}
	c.Group = g
	g.Use(ms...)
	g.POST("", c.Create)
	g.PUT("/:id", c.Update)
}
