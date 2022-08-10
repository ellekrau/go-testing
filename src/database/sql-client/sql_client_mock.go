package sqlclient

import "fmt"

var (
	isMocked = false
)

type (
	sqlClientMock struct {
		mocks map[string]Mock
	}

	Mock struct {
		Query   string
		Args    []interface{}
		Error   error
		Columns []string
		Rows    [][]interface{}
	}
)

func (scm *sqlClientMock) Query(query string, args ...any) (rowsInterface, error) {
	mock, exists := scm.mocks[query]
	if !exists {
		return nil, fmt.Errorf("no mock found for query '%s'", query)
	}

	if mock.Error != nil {
		return nil, mock.Error
	}

	return &sqlRowsMock{
		Columns: mock.Columns,
		Rows:    mock.Rows,
	}, nil
}

func StartMockSqlClient() {
	isMocked = true
}

func StopMockSqlClient() {
	isMocked = false
}

func AddMock(mock Mock) {
	client := dbClient.(*sqlClientMock)
	if client.mocks == nil {
		client.mocks = make(map[string]Mock, 0)
	}

	client.mocks[mock.Query] = mock
}
