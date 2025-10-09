package keyvaluestore_test

import (
	"testing"

	keyvaluestore "github.com/ShabarishRamaswamy/key-value-store.git/src"
)

func TestGet_EmptyKey(t *testing.T) {
	kv := keyvaluestore.GetNewKVStore()
	if kv.Get("") != nil {
		t.Errorf(`Get("") = %q, want nil`, kv.Get(""))
	}
}

func TestGet_KeyExists(t *testing.T) {
	kv := keyvaluestore.GetNewKVStore()
	key := keyvaluestore.GetKeyFrom(100)
	value := "Krishna"

	kv.Insert(key, value)
	if kv.Get(key) != value {
		t.Errorf(`Get("") = %q, want %s`, kv.Get(""), value)
	}
}
