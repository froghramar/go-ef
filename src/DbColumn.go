package src

type DbColumn struct {
	name       string
	columnType DbColumnType
}

type DbColumnType string

const (
	Integer DbColumnType = "int"
)
