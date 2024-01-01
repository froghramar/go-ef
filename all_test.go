package main

import (
	"go-ef/entities"
	"testing"
)

func TestAll(t *testing.T) {
	ctx := CreateDbContext()
	ctx.RegisterTable(entities.User{})
	ctx.Add(entities.User{
		Id:        0,
		FirstName: "Feroz",
		LastName:  "Ahmmed",
		Age:       29,
	})
}
