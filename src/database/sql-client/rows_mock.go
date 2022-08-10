package sqlclient

import "fmt"

type (
	sqlRowsMock struct {
		Columns      []string
		Rows         [][]interface{}
		currentIndex int
	}
)

func (m *sqlRowsMock) HasNext() bool {
	return m.currentIndex < len(m.Rows)
}

func (m *sqlRowsMock) Close() error {
	return nil
}

func (m *sqlRowsMock) Scan(target ...interface{}) error {
	mockedRow := m.Rows[m.currentIndex]

	if len(mockedRow) != len(target) {
		return fmt.Errorf("rows mock invalid destinantion len")
	}

	for index, value := range mockedRow {
		target[index] = value
	}

	m.currentIndex++

	return nil
}
