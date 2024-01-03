package main

import (
	"github.com/stretchr/testify/assert"
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
	query := ctx.BuildQuery()
	assert.Equal(t, "INSERT INTO User (FirstName, LastName, Age) VALUES ('Feroz', 'Ahmmed', 29);", query)
}
