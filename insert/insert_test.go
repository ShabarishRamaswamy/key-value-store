package keyvaluestore

import (
	"encoding/binary"
	"fmt"
	"testing"
)

func TestInsertStringKey(t *testing.T) {
	store := GetNewKVStore()

	k := "Test"
	v := int64(20)

	// TODO: Add a swtich to the type
	if store.Insert(k, v) != v {
		t.Errorf("expected %s, got %s", string(k), store.Insert(k, v).(string))
	}
}

func TestInsertIntegerKey(t *testing.T) {
	store := GetNewKVStore()

	k := 123
	v := int64(20)

	// Use of int directly
	if store.Insert(fmt.Sprintf("%d", k), v) != v {
		t.Errorf("expected %s, got %s", fmt.Sprintf("%d", k), store.Insert(fmt.Sprintf("%d", k), v).(string))
	}

	buf := make([]byte, 8) // int64 is 8 bytes
	unsignedKey := uint64(k)
	binary.BigEndian.PutUint64(buf, unsignedKey)

	// Use of binary arrays
	if store.Insert(string(buf), v) != v {
		t.Errorf("expected %s, got %s", fmt.Sprintf("%d", k), store.Insert(fmt.Sprintf("%d", k), v).(string))
	}
}
