package deleteHandler

import (
	"context"
	"errors"

	"github.com/shabarishramaswamy/key-value-store/src/models"
)

func Delete(ctx context.Context, kvPair models.KeyValuePair) (models.KeyValuePair, error) {
	if ctx.Value(models.GlobalKVStoreName) == nil {
		return models.KeyValuePair{}, errors.New(models.ErrNoGlovalKVStore)
	}

	kvStore := ctx.Value(models.GlobalKVStoreName).(*models.GlobalKVStore)
	currentValueFromKVStore := (*kvStore)[kvPair.Key]
	delete(*kvStore, kvPair.Key)
	// fmt.Println("%+v", *kvStore)

	return models.KeyValuePair{Key: kvPair.Key, Value: currentValueFromKVStore}, nil
}
