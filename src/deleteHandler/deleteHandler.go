package deleteHandler

import (
	"context"
	"database/sql"
	"errors"

	_ "github.com/mattn/go-sqlite3"
	"github.com/shabarishramaswamy/key-value-store/src/commonMethods"
	"github.com/shabarishramaswamy/key-value-store/src/models"
)

func Delete(ctx context.Context, db *sql.DB, kvPair models.KeyValuePair) (models.KeyValuePair, error) {
	if ctx.Value(models.GlobalKVStoreName) == nil {
		return models.KeyValuePair{}, errors.New(models.ErrNoGlovalKVStore)
	}

	kvStore := ctx.Value(models.GlobalKVStoreName).(*models.GlobalKVStore)
	currentValueFromKVStore := (*kvStore)[kvPair.Key]

	delete(*kvStore, kvPair.Key)
	err := DeleteFromDB(db, kvPair)
	if err != nil {
		return models.KeyValuePair{}, err
	}
	commonMethods.DebugWhatsInTheDB()

	return models.KeyValuePair{Key: kvPair.Key, Value: currentValueFromKVStore}, nil
}
