package keyvaluestore

import (
	"fmt"
	"testing"
)

func TestInsertStringKey(t *testing.T) {
	store := GetNewKVStore()

	k := "Test"
	v := int64(20)

	if store.Insert(k, v) != string(k) {
		t.Errorf("expected %s, got %s", string(k), store.Insert(k, v).(string))
	}
}

func TestInsertIntegerKey(t *testing.T) {
	store := GetNewKVStore()

	k := 123
	v := int64(20)

	if store.Insert(fmt.Sprintf("%d", k), v) != fmt.Sprintf("%d", k) {
		t.Errorf("expected %s, got %s", fmt.Sprintf("%d", k), store.Insert(fmt.Sprintf("%d", k), v).(string))
	}
}
