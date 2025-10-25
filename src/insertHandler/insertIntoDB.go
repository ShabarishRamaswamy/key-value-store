package insertHandler

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/shabarishramaswamy/key-value-store/src/databases"
	"github.com/shabarishramaswamy/key-value-store/src/models"
)

func InsertIntoDB(db *sql.DB, kv models.KeyValuePair) error {
	stmt, tx, err := databases.GetPreparedStatementFrom(db, "INSERT INTO kvs(key, value) values(?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(kv.Key, []byte(fmt.Sprintf("%+v", kv.Value)))
	if err != nil {
		return err
	}

	return tx.Commit()
}
