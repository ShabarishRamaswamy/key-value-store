package keyvaluestore

import (
	"fmt"

	"github.com/ShabarishRamaswamy/key-value-store.git/interfaces"
)

type Store map[interfaces.Key]interfaces.Value

// func GetKeyFrom()

func (s Store) Insert(key interfaces.Key, value interfaces.Value) interfaces.Value {
	s[key] = value
	fmt.Println("Stored: ", key, s, s[key])
	return value
}
