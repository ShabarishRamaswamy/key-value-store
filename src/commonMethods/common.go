package commonMethods

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/shabarishramaswamy/key-value-store/src/models"
)

var kvInMemoryStore models.GlobalKVStore = make(models.GlobalKVStore)

func GetNewKVPair() *models.GlobalKVStore {
	return &kvInMemoryStore
}

func DebugWhatsInTheDB() {
	db, err := sql.Open("sqlite3", "databases/kvs.db")
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM kvs")
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer rows.Close()

	for rows.Next() {
		var key string
		var value []byte

		err = rows.Scan(&key, &value)
		if err != nil {
			log.Println(err.Error())
			return
		}
		log.Println(key, string(value))
	}
}
