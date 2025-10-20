package insertHandler

import (
	"context"
	"errors"
	"fmt"

	"github.com/shabarishramaswamy/key-value-store/src/models"
)

func Insert(ctx context.Context, kvPair models.KeyValuePair) (models.KeyValuePair, error) {
	if ctx.Value(models.GlobalKVStoreName) == nil {
		return models.KeyValuePair{}, errors.New(models.ErrNoGlovalKVStore)
	}

	kvStore := ctx.Value(models.GlobalKVStoreName).(*models.GlobalKVStore)
	fmt.Println(kvStore)

	// Add kvPair to Memory
	(*kvStore)[kvPair.Key] = kvPair.Value

	// TODO: Store kvPair to Disk
	return models.KeyValuePair{Key: kvPair.Key, Value: (*kvStore)[kvPair.Key]}, nil
}
