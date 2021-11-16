// Reference: http://en.wikipedia.org/wiki/Associative_array
package linkedhashmap

import (
	"datastruct/lists/doublylinkedlist"
	"datastruct/maps"
	"fmt"
	"strings"
)

func assertLinkedHashMapImplementation() {
	var _ maps.Map = (*Map)(nil)
}

type Map struct {
	table    map[interface{}]interface{}
	ordering *doublylinkedlist.List
}

func New() *Map {
	return &Map{
		table:    make(map[interface{}]interface{}),
		ordering: doublylinkedlist.New(),
	}
}

func (m *Map) Put(key interface{}, value interface{}) {
	if _, contains := m.table[key]; !contains {
		m.ordering.Append(key)
	}
	m.table[key] = value
}

func (m *Map) Get(key interface{}) (value interface{}, found bool) {
	value = m.table[key]
	found = value != nil
	return
}

func (m *Map) Remove(key interface{}) {
	if _, contains := m.table[key]; contains {
		delete(m.table, key)
		index := m.ordering.IndexOf(key)
		m.ordering.Remove(index)
	}
}

// Empty returns true if map does not contain any elements
func (m *Map) Empty() bool {
	return m.Size() == 0
}

// Size returns number of elements in the map.
func (m *Map) Size() int {
	return m.ordering.Size()
}

// Keys returns all keys in-order
func (m *Map) Keys() []interface{} {
	return m.ordering.Values()
}

// Values returns all values in-order based on the key.
func (m *Map) Values() []interface{} {
	values := make([]interface{}, m.Size())
	count := 0
	it := m.Iterator()
	for it.Next() {
		values[count] = it.Value()
		count++
	}
	return values
}

// Clear removes all elements from the map.
func (m *Map) Clear() {
	m.table = make(map[interface{}]interface{})
	m.ordering.Clear()
}

// String returns a string representation of container
func (m *Map) String() string {
	str := "LinkedHashMap\nmap["
	it := m.Iterator()
	for it.Next() {
		str += fmt.Sprintf("%v:%v ", it.Key(), it.Value())
	}
	return strings.TrimRight(str, " ") + "]"

}
