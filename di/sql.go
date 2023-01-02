package sql

import "database/sql"

// type in Go can implement many interface at the same time
// interface could also satisfy many type at the same time
type DB interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
}

// use DB interface instead of sql.DB
// func execQuery(db *sql.DB, query string, args ...interface{}) (int64, error) {
func execQuery(db DB, query string, args ...interface{}) (int64, error) {
	// need to inject with interface that has .Exec
	res, err := db.Exec(query, args...)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
