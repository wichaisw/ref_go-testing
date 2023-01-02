package sql

import (
	"database/sql"
	"testing"
)

type mockDB struct {
	query        string
	lastInsertID int64
	rowsAffected int64
}

// added to satisfy sq.Result signature
func (m *mockDB) LastInsertId() (int64, error) {
	return m.lastInsertID, nil
}

// added to satisfy sq.Result signature
func (m *mockDB) RowsAffected() (int64, error) {
	return m.rowsAffected, nil
}

func (m *mockDB) Exec(query string, args ...interface{}) (sql.Result, error) {
	m.query = query // spy on query
	return m, nil
}

func TestExecQuery(t *testing.T) {
	mock := &mockDB{
		rowsAffected: 32,
	}

	r, _ := execQuery(mock, "SELECT * FROM sql")

	// mock as spy
	if mock.query != "SELECT * FROM sql" {
		t.Error("should have been call db.Exec with query but it not.")
	}

	if r != 32 {
		t.Errorf("should return row effect %d but it got %d.", 32, r)
	}

}
