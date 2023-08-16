package main

import (
	"os"

	env "github.com/DistributedPlayground/go-lib/config"
	"github.com/DistributedPlayground/products/api"
	"github.com/DistributedPlayground/products/config"
	"github.com/DistributedPlayground/products/pkg/store"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

func main() {
	err := env.LoadEnv(&config.Var)
	if err != nil {
		panic(err)
	}

	lg := zerolog.New(os.Stdout)

	port := config.Var.PORT

	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	db := store.MustNewPG()

	kp := store.MustNewKafka()

	api.Start(api.APIConfig{
		DB:     db,
		KP:     kp,
		Port:   port,
		Logger: &lg,
	})
}
