package doublylinkedlist

import (
	"datastruct/containers"
)

func assertIteratorImplementation() {
	var _ containers.ReverseIteratorWithIndex = (*Iterator)(nil)
}

// Iterator 保持迭代器的状态
type Iterator struct {
	list    *List
	index   int
	element *element
}

// Iterator 返回一个有状态的迭代器，其值可以通过索引获取。
func (list *List) Iterator() Iterator {
	return Iterator{list: list, index: -1}
}

// Next 将迭代器移动到下一个元素，如果容器中有下一个元素，则返回 true。
// 如果 Next() 返回 true，则可以通过 Index() 和 Value() 检索下一个元素的索引和值。
// 如果 Next() 是第一次调用，那么它会将迭代器指向第一个元素（如果存在）。
// 修改迭代器的状态。
func (iterator *Iterator) Next() bool {
	if iterator.index < iterator.list.size {
		iterator.index++
	}
	if !iterator.list.withinRange(iterator.index) {
		iterator.element = nil
		return false
	}

	if iterator.index == 0 {
		iterator.element = iterator.list.first
	} else {
		iterator.element = iterator.element.next
	}
	return true
}

// Prev 将迭代器移动到前一个元素，如果容器中有前一个元素，则返回 true。
// 如果 Prev() 返回 true，则可以通过 Index() 和 Value() 检索前一个元素的索引和值。
// 修改迭代器的状态。
func (iterator *Iterator) Prev() bool {
	if iterator.index >= 0 {
		iterator.index--
	}

	if !iterator.list.withinRange(iterator.index) {
		iterator.element = nil
		return false
	}

	if iterator.index == iterator.list.size-1 {
		iterator.element = iterator.list.last
	} else {
		iterator.element = iterator.element.prev
	}
	return true
}

// Value 返回当前元素的值。
// Does not modify the state of the iterator.
func (iterator *Iterator) Value() interface{} {
	return iterator.element.value
}

// Index 返回当前元素的索引。
// 不修改迭代器的状态。
func (iterator *Iterator) Index() int {
	return iterator.index
}

// Begin 将迭代器重置为其初始状态（one-before-first）
// 调用 Next() 以获取第一个元素（如果有）。
func (iterator *Iterator) Begin() {
	iterator.index = -1
	iterator.element = nil
}

// End 将迭代器移过最后一个元素（one-past-the-end）。
// 调用 Prev() 以获取最后一个元素（如果有）。
func (iterator *Iterator) End() {
	iterator.index = iterator.list.size
	iterator.element = iterator.list.last
}

// First 将迭代器移动到第一个元素，如果容器中有第一个元素，则返回 true。
// 如果 First() 返回 true，则可以通过 Index() 和 Value() 检索第一个元素的索引和值。
// 修改迭代器的状态。
func (iterator *Iterator) First() bool {
	iterator.Begin()
	return iterator.Next()
}

// Last 将迭代器移动到最后一个元素，如果容器中有最后一个元素，则返回 true。
// 如果 Last() 返回 true，则可以通过 Index() 和 Value() 检索最后一个元素的索引和值。
// 修改迭代器的状态。
func (iterator *Iterator) Last() bool {
	iterator.End()
	return iterator.Prev()
}
