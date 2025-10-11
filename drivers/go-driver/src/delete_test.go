package keyvaluestore_test

import (
	"testing"

	keyvaluestore "github.com/ShabarishRamaswamy/key-value-store.git/src"
)

func Test_Delete_ValueExists(t *testing.T) {
	s := keyvaluestore.GetNewKVStore()

	key := keyvaluestore.GetKeyFrom("brahmarpanam")
	value := "brahmahavair"

	s.Insert(key, value)

	delRes := s.Delete(key)
	if delRes != value {
		t.Errorf(`s.Delete(%s) = %s, want %s`, key, delRes, value)
	}
}
