package keyvaluestore

import (
	"fmt"

	"github.com/ShabarishRamaswamy/key-value-store.git/interfaces"
)

type Store map[string]interface{}

func GetNewKVStore() *Store {
	return &Store{}
}

// func GetKeyFrom()

func (s Store) Insert(key string, value interface{}) interface{} {
	fmt.Println("Stored: ", key, s, s[key])
	return value
}

func Get[Key interfaces.Key, Value interfaces.Value](key Key) interface{} {
	fmt.Println("Getting the Value for: ", key)
	v := "Not Yet Implemented"
	return v
}
