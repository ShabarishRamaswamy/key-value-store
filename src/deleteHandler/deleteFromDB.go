package deleteHandler

import (
	"database/sql"

	"github.com/shabarishramaswamy/key-value-store/src/databases"
	"github.com/shabarishramaswamy/key-value-store/src/models"
)

func DeleteFromDB(db *sql.DB, kvPair models.KeyValuePair) error {
	stmt, tx, err := databases.GetPreparedStatementFrom(db, "DELETE FROM kvs WHERE key = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(kvPair.Key)
	if err != nil {
		return err
	}

	return tx.Commit()
}
