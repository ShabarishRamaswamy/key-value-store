package keyvaluestore

import (
	"github.com/ShabarishRamaswamy/key-value-store.git/interfaces"
)

func (s Store) Get(key interfaces.Key) interfaces.Value {
	// fmt.Println("Getting the Value for: ", key, s[key])
	return s[key]
}
