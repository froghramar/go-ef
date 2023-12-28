package src

import "reflect"

type DbContext struct {
	tables []DbTable
}

func CreateDbContext() DbContext {
	return DbContext{
		tables: make([]DbTable, 0),
	}
}

func RegisterTable[T any](ctx DbContext, entity T) {
	tableName := reflect.TypeOf(entity).Name()
	ctx.tables = append(ctx.tables, DbTable{
		tableName: tableName,
	})
}
