package linkedhashmap

import "datastruct/containers"

func assertEnumerableImplementation() {
	var _ containers.EnumerableWithKey = (*Map)(nil)
}

// Each 为每个元素调用一次给定的函数，传递该元素的键和值。
func (m *Map) Each(f func(key interface{}, value interface{})) {
	iterator := m.Iterator()
	for iterator.Next() {
		f(iterator.Key(), iterator.Value())
	}
}

// Map 为每个元素调用一次给定的函数并返回一个容器
// 包含由给定函数返回的值作为键值对。
func (m *Map) Map(f func(key1 interface{}, value1 interface{}) (interface{}, interface{})) *Map {
	newMap := New()
	iterator := m.Iterator()
	for iterator.Next() {
		key2, value2 := f(iterator.Key(), iterator.Value())
		newMap.Put(key2, value2)
	}
	return newMap
}

// Select 返回一个新容器，其中包含给定函数返回真值的所有元素。
func (m *Map) Select(f func(key interface{}, value interface{}) bool) *Map {
	newMap := New()
	iterator := m.Iterator()
	for iterator.Next() {
		if f(iterator.Key(), iterator.Value()) {
			newMap.Put(iterator.Key(), iterator.Value())
		}
	}
	return newMap
}

// Any 将容器的每个元素传递给给定的函数，如果函数对任何元素返回 true，则返回 true。
func (m *Map) Any(f func(key interface{}, value interface{}) bool) bool {
	iterator := m.Iterator()
	for iterator.Next() {
		if f(iterator.Key(), iterator.Value()) {
			return true
		}
	}
	return false
}

// All 将容器的每个元素传递给给定的函数，如果该函数对所有元素都返回 true，则返回 true。
func (m *Map) All(f func(key interface{}, value interface{}) bool) bool {
	iterator := m.Iterator()
	for iterator.Next() {
		if !f(iterator.Key(), iterator.Value()) {
			return false
		}
	}
	return true
}

// Find 将容器的每个元素传递给给定的函数，并返回函数为 true 或 nil,nil 的第一个 (key,value)，否则如果没有元素符合条件。
func (m *Map) Find(f func(key interface{}, value interface{}) bool) (interface{}, interface{}) {
	iterator := m.Iterator()
	for iterator.Next() {
		if f(iterator.Key(), iterator.Value()) {
			return iterator.Key(), iterator.Value()
		}
	}
	return nil, nil
}
