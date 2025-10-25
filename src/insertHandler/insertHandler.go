package insertHandler

import (
	"context"
	"errors"
	"log"

	"github.com/shabarishramaswamy/key-value-store/src/commonMethods"
	"github.com/shabarishramaswamy/key-value-store/src/models"
)

func Insert(ctx context.Context, kvPair models.KeyValuePair) (models.KeyValuePair, error) {
	if ctx.Value(models.GlobalKVStoreName) == nil {
		return models.KeyValuePair{}, errors.New(models.ErrNoGlovalKVStore)
	}

	kvStore := ctx.Value(models.GlobalKVStoreName).(*models.GlobalKVStore)

	// Add kvPair to Memory
	(*kvStore)[kvPair.Key] = kvPair.Value

	// This should be a go func.
	// go InsertIntoDB(kvPair)
	err := InsertIntoDB(kvPair)
	if err != nil {
		log.Println(err.Error())
	}
	commonMethods.DebugWhatsInTheDB()

	// TODO: Store kvPair to Disk
	return models.KeyValuePair{Key: kvPair.Key, Value: (*kvStore)[kvPair.Key]}, nil
}
