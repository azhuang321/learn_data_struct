package arraylist

import "datastruct/containers"

func assertEnumerableImplementation() {
	var _ containers.EnumerableWithIndex = (*List)(nil)
}

// Each 为每个元素调用一次给定的函数，传递该元素的索引和值。
func (list *List) Each(f func(index int, value interface{})) {
	iterator := list.Iterator()
	for iterator.Next() {
		f(iterator.Index(), iterator.Value())
	}
}

// Map 为每个元素调用一次给定的函数，并返回一个包含给定函数返回值的容器。
func (list *List) Map(f func(index int, value interface{}) interface{}) *List {
	newList := &List{}
	iterator := list.Iterator()
	for iterator.Next() {
		newList.Add(f(iterator.Index(), iterator.Value()))
	}
	return newList
}

// Select 返回一个新容器，其中包含给定函数返回真值的所有元素。
func (list *List) Select(f func(index int, value interface{}) bool) *List {
	newList := &List{}
	iterator := list.Iterator()
	for iterator.Next() {
		if f(iterator.Index(), iterator.Value()) {
			newList.Add(iterator.Value())
		}
	}
	return newList
}

// Any 将集合的每个元素传递给给定的函数，如果该函数对任何元素返回 true，则返回 true。
func (list *List) Any(f func(index int, value interface{}) bool) bool {
	iterator := list.Iterator()
	for iterator.Next() {
		if f(iterator.Index(), iterator.Value()) {
			return true
		}
	}
	return false
}

// All 将集合的每个元素传递给给定的函数，如果该函数对所有元素都返回 true，则返回 true。
func (list *List) All(f func(index int, value interface{}) bool) bool {
	iterator := list.Iterator()
	for iterator.Next() {
		if !f(iterator.Index(), iterator.Value()) {
			return false
		}
	}
	return true
}

// Find 将容器的每个元素传递给给定的函数，并返回函数为真的第一个 (index,value) 或 -1,nil，否则如果没有元素符合条件。
func (list *List) Find(f func(index int, value interface{}) bool) (int, interface{}) {
	iterator := list.Iterator()
	for iterator.Next() {
		if f(iterator.Index(), iterator.Value()) {
			return iterator.Index(), iterator.Value()
		}
	}
	return -1, nil
}
