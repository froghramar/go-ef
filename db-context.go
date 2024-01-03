package main

import (
	"fmt"
	"github.com/thoas/go-funk"
	"go-ef/utils"
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
	tableType := reflect.TypeOf(entity)
	tableName := tableType.Name()
	columns := make([]DbColumn, 0)
	for i := 0; i < tableType.NumField(); i++ {
		var column = tableType.Field(i)
		columns = append(columns, DbColumn{name: column.Name})
	}
	ctx.tables = append(ctx.tables, DbTable{
		tableName: tableName,
		columns:   columns,
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

	columnNames := utils.Select(table.columns, func(column DbColumn) string { return column.name })
	commaSeparatedColumnNames := strings.Join(columnNames[:], ", ")

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("INSERT INTO %s (%s)", table.tableName, commaSeparatedColumnNames))
	return sb.String()
}
