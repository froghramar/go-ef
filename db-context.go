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
		column := tableType.Field(i)
		columnType := getDbColumnType(column.Type)
		columns = append(columns, DbColumn{name: column.Name, columnType: columnType})
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
	for _, record := range table.records {
		structVal := reflect.ValueOf(record)
		values := utils.Select(table.columns, func(column DbColumn) string {
			fieldVal := structVal.FieldByName(column.name)
			return addQuotesIfNecessary(fmt.Sprint(fieldVal.Interface()), column.columnType)
		})
		sb.WriteString(fmt.Sprintf("(%s)", strings.Join(values, ", ")))
	}

	return fmt.Sprintf("INSERT INTO %s (%s) VALUES %s;", table.tableName, commaSeparatedColumnNames, sb.String())
}

func addQuotesIfNecessary(value string, columnType DbColumnType) string {
	if columnType == String {
		return addQuotes(value)
	}

	return value
}

func addQuotes(value string) string {
	return fmt.Sprintf("'%s'", value)
}

func getDbColumnType(columnType reflect.Type) DbColumnType {
	switch columnType.Kind() {
	default:
		panic("Unrecognized column type")
	case reflect.Int:
		return Integer
	case reflect.String:
		return String
	}
}
