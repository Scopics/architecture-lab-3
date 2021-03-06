// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/Scopics/architecture-lab-3/server/restaurant"
)

// Injectors from modules.go:

func ComposeApiServer(port int) (*APIServer, error) {
	db, err := NewDbConnection()
	if err != nil {
		return nil, err
	}
	store := restaurant.NewStore(db)
	httpHandlerFunc := restaurant.HttpHandler(store)
	apiServer := &APIServer{
		Port:    port,
		Handler: httpHandlerFunc,
	}
	return apiServer, nil
}
