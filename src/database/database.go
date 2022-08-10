package database

import (
	"fmt"
	sqlclient "go-testing/database/sql-client"
	"go-testing/domain/user"
)

func StartDB() error {
	sqlclient.StartMockSqlClient()

	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		"admin", "admin", "localhost", "5432", "go_testing")
	db, err := sqlclient.Open("postgres", connectionString)
	if err != nil {
		return fmt.Errorf("StartDB error: %s", err.Error())
	}

	sqlclient.AddMock(sqlclient.Mock{
		Query: "select * from users where ID = 1;",
		Args:  nil,
		//Error:   fmt.Errorf("error creating query"),
		Columns: []string{"id", "email"},
		Rows: [][]interface{}{
			{1, "Elle mock", "ellemock@elle.com"},
		},
	})

	rows, err := db.Query("select * from users where ID = 1;")
	if err != nil {
		return fmt.Errorf("execute simple query error: %s", err.Error())
	}
	defer rows.Close()

	var user user.User
	for rows.HasNext() {
		if err = rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return fmt.Errorf("sacan user error: %s", err.Error())
		}
	}

	fmt.Println(user)

	return nil
}
