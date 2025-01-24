package main

import (
	"context"

	"github.com/mvndaai/known-socially/internal/db"
	"github.com/mvndaai/known-socially/internal/router"
)

func main() {
	ctx := context.Background()

	db := db.Postgres{}
	err := db.Connect(ctx)
	if err != nil {
		panic(err)
	}
	defer db.Close(ctx)

	router.StartServer()
}
