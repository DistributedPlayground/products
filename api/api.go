package api

import (
	"net/http"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"

	"github.com/DistributedPlayground/products/api/handler"
	"github.com/DistributedPlayground/products/pkg/service"
)

type APIConfig struct {
	DB     *sqlx.DB
	KP     *kafka.Producer
	Logger *zerolog.Logger
	Port   string
}

func heartbeat(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func Start(config APIConfig) {
	e := echo.New()
	e.Logger.Fatal(e.Start(":" + config.Port))

	e.GET("/heartbeat", heartbeat)

	repos := NewRepos(config.DB)
	messages := NewMessages(config.KP)
	services := NewServices(repos, messages)

	collectionRoute(services, e)
	productRoute(services, e)

	e.Logger.Fatal(e.Start(":" + config.Port))
}

// func baseMiddleware(logger *zerolog.Logger, e *echo.Echo) {
// 	e.Use(libmiddleware.Recover())
// 	e.Use(libmiddleware.RequestId())
// 	e.Use(libmiddleware.Logger(logger))
// 	e.Use(libmiddleware.LogRequest())
// }

func collectionRoute(services service.Services, e *echo.Echo) {
	handler := handler.NewCollection(services.Collection)
	handler.RegisterRoutes(e.Group("/collection"))
}

func productRoute(services service.Services, e *echo.Echo) {
	handler := handler.NewProduct(services.Product)
	handler.RegisterRoutes(e.Group("/product"))
}
