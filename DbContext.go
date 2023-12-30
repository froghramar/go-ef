package main

import (
	"reflect"
)

type DbContext struct {
	tables []DbTable
}

func CreateDbContext() DbContext {
	return DbContext{
		tables: make([]DbTable, 0),
	}
}

func (ctx DbContext) RegisterTable(entity any) {
	tableName := reflect.TypeOf(entity).Name()
	ctx.tables = append(ctx.tables, DbTable{
		tableName: tableName,
	})
}
