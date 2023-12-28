package main

import (
	"go-ef/entities"
	"go-ef/src"
)

func main() {
	ctx := src.CreateDbContext()
	src.RegisterTable(ctx, entities.User{})
}
