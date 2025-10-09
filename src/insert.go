package keyvaluestore

import (
	"fmt"
	"strconv"

	"github.com/ShabarishRamaswamy/key-value-store.git/interfaces"
)

type Store map[interfaces.Key]interfaces.Value

func GetNewKVStore() *Store {
	return &Store{}
}

func GetKeyFrom(key interface{}) interfaces.Key {
	switch key := key.(type) {
	case int:
		// buf := make([]byte, 8) // int64 is 8 bytes
		// unsignedKey := uint64(key)
		// binary.BigEndian.PutUint64(buf, unsignedKey)
		return interfaces.Key(strconv.Itoa(key))

	case string:
		return interfaces.Key(key)

	default:
		return interfaces.Key("")
	}
}

// func GetKeyFrom()

func (s Store) Insert(key interfaces.Key, value interfaces.Value) interfaces.Value {
	s[key] = value
	fmt.Println("Stored: ", key, s, s[key])
	return value
}
