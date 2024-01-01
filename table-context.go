package main

type TableContext struct {
	table     DbTable
	dbContext DbContext
}

func (tableCtx TableContext) Add(entity any) {
	tableCtx.table.records = append(tableCtx.table.records, entity)
}
