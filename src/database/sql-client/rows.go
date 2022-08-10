package sqlclient

import "database/sql"

type (
	rows struct {
		sqlRows *sql.Rows
	}

	rowsInterface interface {
		HasNext() bool
		Close() error
		Scan(target ...interface{}) error
	}
)

func (rows *rows) HasNext() bool {
	return rows.sqlRows.Next()
}

func (rows *rows) Close() error {
	return rows.sqlRows.Close()
}

func (rows *rows) Scan(target ...interface{}) error {
	return rows.sqlRows.Scan(target...)
}
