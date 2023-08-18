package handler

import (
	"net/http"

	"github.com/DistributedPlayground/go-lib/common"
	"github.com/DistributedPlayground/go-lib/httperror"
	"github.com/DistributedPlayground/products/pkg/model"
	"github.com/DistributedPlayground/products/pkg/service"
	"github.com/labstack/echo/v4"
)

type Product interface {
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	RegisterRoutes(g *echo.Group, ms ...echo.MiddlewareFunc)
}

type product struct {
	service service.Product
	Group   *echo.Group
}

func NewProduct(service service.Product) Product {
	return &product{service: service}
}

func (p product) Create(ctx echo.Context) error {
	body := model.ProductUpsert{}
	err := ctx.Bind(&body)
	if err != nil {
		common.LogDPError(ctx, err, "Product Create: Bad Request")
		return httperror.BadRequest400(ctx)
	}

	product, err := p.service.Create(ctx.Request().Context(), body)
	if err != nil {
		common.LogDPError(ctx, err, "Product Create: Internal Error")
		return httperror.Internal500(ctx)
	}
	return ctx.JSON(http.StatusCreated, product)
}

func (p product) Update(ctx echo.Context) error {
	body := model.ProductUpsert{}
	err := ctx.Bind(&body)
	if err != nil {
		common.LogDPError(ctx, err, "Product Update: Bad Request")
		return httperror.BadRequest400(ctx)
	}

	product, err := p.service.Update(ctx.Request().Context(), ctx.Param("id"), body)
	if err != nil {
		common.LogDPError(ctx, err, "Product Update: Internal Error")
		return httperror.Internal500(ctx)
	}
	return ctx.JSON(http.StatusOK, product)
}

func (p product) RegisterRoutes(g *echo.Group, ms ...echo.MiddlewareFunc) {
	if g == nil {
		panic("no group attached to the product handler")
	}
	p.Group = g
	g.Use(ms...)
	g.POST("", p.Create)
	g.PUT("/:id", p.Update)
}
