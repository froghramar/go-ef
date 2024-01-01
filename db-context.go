package main

import (
	"github.com/thoas/go-funk"
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

func (ctx DbContext) Table(entity any) TableContext {
	tableName := reflect.TypeOf(entity).Name()
	table := funk.Find(ctx.tables, func(table DbTable) bool {
		return table.tableName == tableName
	})

	if table == nil {
		panic("Table not found")
	}

	return TableContext{dbContext: ctx, table: table.(DbTable)}
}
