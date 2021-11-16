package arraylist

import "datastruct/containers"

func assertIteratorImplementation() {
	var _ containers.ReverseIteratorWithIndex = (*Iterator)(nil)
}

// Iterator 保持迭代器的状态
type Iterator struct {
	list  *List
	index int
}

// Iterator 返回一个有状态的迭代器，其值可以通过索引获取。
func (list *List) Iterator() Iterator {
	return Iterator{list: list, index: -1}
}

// Prev 将迭代器移动到前一个元素，如果容器中有前一个元素，则返回 true。如果 Prev() 返回 true，则可以通过 Index() 和 Value() 检索前一个元素的索引和值。修改迭代器的状态。
func (iterator *Iterator) Prev() bool {
	if iterator.index >= 0 {
		iterator.index--
	}
	return iterator.list.withinRange(iterator.index)
}

// Next 将迭代器移动到下一个元素，如果容器中有下一个元素，则返回 true。如果 Next() 返回 true，则可以通过 Index() 和 Value() 检索下一个元素的索引和值。如果 Next() 是第一次调用，那么它会将迭代器指向第一个元素（如果存在）。修改迭代器的状态。
func (iterator *Iterator) Next() bool {
	if iterator.index < iterator.list.size {
		iterator.index++
	}
	return iterator.list.withinRange(iterator.index)
}

// Value 返回当前元素的值。不修改迭代器的状态。
func (iterator *Iterator) Value() interface{} {
	return iterator.list.elements[iterator.index]
}

// Index 返回当前元素的索引。不修改迭代器的状态。
func (iterator *Iterator) Index() int {
	return iterator.index
}

// Begin 将迭代器重置为其初始状态（一个在第一个之前）调用 Next() 以获取第一个元素（如果有）。
func (iterator *Iterator) Begin() {
	iterator.index = -1
}

// End 将迭代器移过最后一个元素（one-past-the-end）。调用 Prev() 以获取最后一个元素（如果有）。
func (iterator *Iterator) End() {
	iterator.index = iterator.list.size
}

// First 将迭代器移动到第一个元素，如果容器中有第一个元素，则返回 true。如果 First() 返回 true，则可以通过 Index() 和 Value() 检索第一个元素的索引和值。修改迭代器的状态。
func (iterator *Iterator) First() bool {
	iterator.Begin()
	return iterator.Next()
}

// Last 将迭代器移动到最后一个元素，如果容器中有最后一个元素，则返回 true。如果 Last() 返回 true，则可以通过 Index() 和 Value() 检索最后一个元素的索引和值。修改迭代器的状态。
func (iterator *Iterator) Last() bool {
	iterator.End()
	return iterator.Prev()
}
