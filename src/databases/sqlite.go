package databases

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/shabarishramaswamy/key-value-store/src/models"
)

func RunInitDBLoop() (*sql.DB, error) {
	db, err := OpenNewConnection()
	if err != nil {
		return nil, err
	}

	err = CreateNewDatabase(db)
	return db, err
}

func OpenNewConnection() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", models.SQLITEDBPath)
	// defer db.Close()
	return db, err
}

func CreateNewDatabase(db *sql.DB) error {
	_, err := db.Exec(
		`CREATE TABLE IF NOT EXISTS kvs (
    	key TEXT NOT NULL PRIMARY KEY,
    	value BLOB
		);`)

	return err
}

func GetPreparedStatementFrom(db *sql.DB, sqlString string) (*sql.Stmt, *sql.Tx, error) {
	tx, err := db.Begin()
	if err != nil {
		return nil, nil, err
	}

	stmt, err := tx.Prepare(sqlString)
	return stmt, tx, err
}
