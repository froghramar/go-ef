package main

import (
	"fmt"
	"github.com/thoas/go-funk"
	"log"
	"reflect"
	"strings"
)

type DbContext struct {
	tables []DbTable
}

func CreateDbContext() DbContext {
	return DbContext{
		tables: make([]DbTable, 0),
	}
}

func (ctx *DbContext) RegisterTable(entity any) {
	tableName := reflect.TypeOf(entity).Name()
	ctx.tables = append(ctx.tables, DbTable{
		tableName: tableName,
	})
}

func (ctx *DbContext) Add(entity any) {
	tableName := reflect.TypeOf(entity).Name()
	tableIndex := funk.IndexOf(ctx.tables, func(table DbTable) bool {
		return table.tableName == tableName
	})

	if tableIndex == -1 {
		panic("Table not found")
	}

	ctx.tables[tableIndex].records = append(ctx.tables[tableIndex].records, entity)
}

func (ctx *DbContext) BuildQuery() string {
	var sb strings.Builder
	for _, table := range ctx.tables {
		if len(table.records) == 0 {
			continue
		}
		query := generateQuery(table)
		sb.WriteString(query)
	}
	return sb.String()
}

func (ctx *DbContext) Save() {
	query := ctx.BuildQuery()
	log.Println(query)
}

func generateQuery(table DbTable) string {
	if len(table.records) == 0 {
		panic("No records")
	}

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("INSERT INTO %s", table.tableName))
	return sb.String()
}
