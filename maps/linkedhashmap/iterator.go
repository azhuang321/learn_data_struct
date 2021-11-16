package linkedhashmap

import (
	"datastruct/containers"
	"datastruct/lists/doublylinkedlist"
)

func assertIteratorImplementation() {
	var _ containers.ReverseIteratorWithKey = (*Iterator)(nil)
}

// Iterator 保持迭代器的状态
type Iterator struct {
	iterator doublylinkedlist.Iterator
	table    map[interface{}]interface{}
}

// Iterator 返回一个有状态的迭代器，其元素是键值对。
func (m *Map) Iterator() Iterator {
	return Iterator{
		iterator: m.ordering.Iterator(),
		table:    m.table,
	}
}

// Next 将迭代器移动到下一个元素，如果容器中有下一个元素，则返回 true。
// 如果 Next() 返回 true，则可以通过 Key() 和 Value() 检索下一个元素的键和值。
// 如果 Next() 是第一次调用，那么它会将迭代器指向第一个元素（如果存在）。
// 修改迭代器的状态。
func (iterator *Iterator) Next() bool {
	return iterator.iterator.Next()
}

// Prev 将迭代器移动到前一个元素，如果容器中有前一个元素，则返回 true。
// 如果 Prev() 返回 true，则可以通过 Key() 和 Value() 检索前一个元素的键和值。
// 修改迭代器的状态。
func (iterator *Iterator) Prev() bool {
	return iterator.iterator.Prev()
}

// Value 返回当前元素的值。
// 不修改迭代器的状态。
func (iterator *Iterator) Value() interface{} {
	key := iterator.iterator.Value()
	return iterator.table[key]
}

// Key 返回当前元素的键。
// 不修改迭代器的状态。
func (iterator *Iterator) Key() interface{} {
	return iterator.iterator.Value()
}

// Begin 将迭代器重置为其初始状态（one-before-first）
// 调用 Next() 以获取第一个元素（如果有）。
func (iterator *Iterator) Begin() {
	iterator.iterator.Begin()
}

// End 将迭代器移过最后一个元素（one-past-the-end）。
// 调用 Prev() 以获取最后一个元素（如果有）。
func (iterator *Iterator) End() {
	iterator.iterator.End()
}

// First 将迭代器移动到第一个元素，如果容器中有第一个元素，则返回 true。
// 如果 First() 返回 true，则 Key() 和 Value() 可以检索第一个元素的键和值。
// 修改迭代器的状态
func (iterator *Iterator) First() bool {
	return iterator.iterator.First()
}

// Last 将迭代器移动到最后一个元素，如果容器中有最后一个元素，则返回 true。
// 如果 Last() 返回 true，则可以通过 Key() 和 Value() 检索最后一个元素的键和值。
// 修改迭代器的状态。
func (iterator *Iterator) Last() bool {
	return iterator.iterator.Last()
}
