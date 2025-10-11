package keyvaluestore_test

import (
	"fmt"
	"testing"

	keyvaluestore "github.com/ShabarishRamaswamy/key-value-store.git/src"
)

func TestInsertStringKey(t *testing.T) {
	store := keyvaluestore.GetNewKVStore()

	k := keyvaluestore.GetKeyFrom("Test")
	v := int64(20)

	// TODO: Add a swtich to the type
	if store.Insert(k, v) != v {
		t.Errorf("expected %s, got %s", string(k), store.Insert(k, v).(string))
	}
}

func TestInsertIntegerKey(t *testing.T) {
	store := keyvaluestore.GetNewKVStore()

	k := keyvaluestore.GetKeyFrom(123)
	v := int64(20)

	if store.Insert(k, v) != v {
		t.Errorf("expected %s, got %s", fmt.Sprintf("%+v", k), store.Insert(k, v).(string))
	}
}
