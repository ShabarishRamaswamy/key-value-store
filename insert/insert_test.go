package keyvaluestore

import (
	"testing"

	"github.com/ShabarishRamaswamy/key-value-store.git/interfaces"
)

func TestInsert(t *testing.T) {
	k := interfaces.Key("Test")
	v := int64(20)

	if Insert(k, v) != string(k) {
		t.Errorf("expected %s, got %s", string(k), string(Insert(k, v)))
	}
}
