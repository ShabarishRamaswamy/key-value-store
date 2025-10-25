package insertHandler

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/shabarishramaswamy/key-value-store/src/models"
)

func InsertIntoDB(kv models.KeyValuePair) error {
	db, err := sql.Open("sqlite3", "databases/kvs.db")
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS kvs (
    	key TEXT NOT NULL PRIMARY KEY,
    	value BLOB
		);`)
	if err != nil {
		return err
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(
		"INSERT INTO kvs(key, value) values (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(kv.Key, []byte(fmt.Sprintf("%+v", kv.Value)))
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
