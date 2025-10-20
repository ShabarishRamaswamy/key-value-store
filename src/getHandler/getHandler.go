package getHandler

import (
	"context"
	"errors"

	"github.com/shabarishramaswamy/key-value-store/src/models"
)

func Get(ctx context.Context, kvPair models.KeyValuePair) (models.KeyValuePair, error) {
	if ctx.Value(models.GlobalKVStoreName) == nil {
		return models.KeyValuePair{}, errors.New(models.ErrNoGlovalKVStore)
	}
	kvStore := ctx.Value(models.GlobalKVStoreName).(*models.GlobalKVStore)

	return models.KeyValuePair{Key: kvPair.Key, Value: (*kvStore)[kvPair.Key]}, nil
}
