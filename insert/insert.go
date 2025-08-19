package keyvaluestore

import (
	"fmt"

	"github.com/ShabarishRamaswamy/key-value-store.git/interfaces"
)

func Insert[V interfaces.Value](key interfaces.Key, value V) string {
	fmt.Println("Stored: ", key)
	return string(key)
}
