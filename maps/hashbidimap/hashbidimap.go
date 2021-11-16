// Reference: https://en.wikipedia.org/wiki/Bidirectional_map
package hashbidimap

import (
	"datastruct/maps"
	"datastruct/maps/hashmap"
	"fmt"
)

func assertHashBidiMapImplementation() {
	var _ maps.Map = (*Map)(nil)
}

type Map struct {
	forwardMap hashmap.Map
	inverseMap hashmap.Map
}

func New() *Map {
	return &Map{
		forwardMap: *hashmap.New(),
		inverseMap: *hashmap.New(),
	}
}

func (m *Map) Put(key interface{}, value interface{}) {
	if valueByKey, ok := m.forwardMap.Get(key); ok {
		m.inverseMap.Remove(valueByKey)
	}

	if keyByValue, ok := m.inverseMap.Get(value); ok {
		m.forwardMap.Remove(keyByValue)
	}

	m.forwardMap.Put(key, value)
	m.inverseMap.Put(value, key)
}

func (m *Map) Get(key interface{}) (value interface{}, found bool) {
	return m.forwardMap.Get(key)
}

func (m *Map) GetKey(value interface{}) (key interface{}, found bool) {
	return m.inverseMap.Get(value)
}

func (m *Map) Remove(key interface{}) {
	if value, contains := m.forwardMap.Get(key); contains {
		m.forwardMap.Remove(key)
		m.inverseMap.Remove(value)
	}
}

// Empty returns true if map does not contain any elements
func (m *Map) Empty() bool {
	return m.Size() == 0
}

// Size returns number of elements in the map.
func (m *Map) Size() int {
	return m.forwardMap.Size()
}

// Keys returns all keys (random order).
func (m *Map) Keys() []interface{} {
	return m.forwardMap.Keys()
}

// Values returns all values (random order).
func (m *Map) Values() []interface{} {
	return m.inverseMap.Keys()
}

// Clear removes all elements from the map.
func (m *Map) Clear() {
	m.forwardMap.Clear()
	m.inverseMap.Clear()
}

// String returns a string representation of container
func (m *Map) String() string {
	str := "HashBidiMap\n"
	str += fmt.Sprintf("%v", m.forwardMap)
	return str
}
