//go:build wireinject
// +build wireinject

package main

import (
	"github.com/Scopics/architecture-lab-3/server/restaurant"
	"github.com/google/wire"
)

func ComposeApiServer(port int) (*APIServer, error) {
	wire.Build(
		NewDbConnection,
		restaurant.Providers,
		wire.Struct(new(APIServer), "Port", "Handler"),
	)
	return nil, nil
}
