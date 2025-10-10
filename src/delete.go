package keyvaluestore

import "github.com/ShabarishRamaswamy/key-value-store.git/interfaces"

func (s Store) Delete(key interfaces.Key) interfaces.Value {
	currentValue := s.Get(key)
	delete(s, key)

	return currentValue
}
