package commonMethods

import "github.com/shabarishramaswamy/key-value-store/src/models"

var kvInMemoryStore models.GlobalKVStore = make(models.GlobalKVStore)

func GetNewKVPair() *models.GlobalKVStore {
	return &kvInMemoryStore
}
