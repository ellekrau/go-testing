package sqlclient

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"os"
)

var (
	dbClient sqlClientInterface
)

const (
	environment = "ENVIRONMENT"
	production  = "PRODUCTION"
)

type (
	sqlClientInterface interface {
		Query(query string, args ...any) (rowsInterface, error)
	}

	sqlClient struct {
		db *sql.DB
	}
)

func isProduction() bool {
	return os.Getenv(environment) == production
}

func Open(driverName, connectionString string) (sqlClientInterface, error) {
	if isMocked && !isProduction() {
		dbClient = &sqlClientMock{}
		return dbClient, nil
	}

	if driverName == "" {
		return nil, fmt.Errorf("sql-client driver name input is empty")
	}

	db, err := sql.Open(driverName, connectionString)
	if err != nil {
		return nil, fmt.Errorf("sql client open error: %s", err.Error())
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("sql client ping error: %s", err.Error())
	}

	dbClient = &sqlClient{db: db}
	return dbClient, nil
}

func (sc *sqlClient) Query(query string, args ...any) (rowsInterface, error) {
	resultRows, err := sc.db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("sql client query error: %s", err.Error())
	}

	return &rows{sqlRows: resultRows}, nil
}
