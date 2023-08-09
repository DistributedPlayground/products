package main

import (
	env "github.com/DistributedPlayground/go-lib/config"
	"github.com/DistributedPlayground/products/config"
)

func main() {
	err := env.LoadEnv(&config.Var)
	if err != nil {
		panic(err)
	}
}
