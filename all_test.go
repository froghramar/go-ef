package main

import (
	"go-ef/entities"
	"testing"
)

func TestAll(t *testing.T) {
	ctx := CreateDbContext()
	ctx.RegisterTable(entities.User{})
}
