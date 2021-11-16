package linkedliststack

import "datastruct/containers"

func assertIteratorImplementation() {
	var _ containers.IteratorWithIndex = (*Iterator)(nil)
}

// Iterator 返回一个有状态的迭代器，其值可以通过索引获取。
type Iterator struct {
	stack *Stack
	index int
}

// Iterator 返回一个有状态的迭代器，其值可以通过索引获取。
func (stack *Stack) Iterator() Iterator {
	return Iterator{stack: stack, index: -1}
}

// Next 将迭代器移动到下一个元素，如果容器中有下一个元素，则返回 true。
// 如果 Next() 返回 true，则可以通过 Index() 和 Value() 检索下一个元素的索引和值。
// 如果 Next() 是第一次调用，那么它会将迭代器指向第一个元素（如果存在）。
// 修改迭代器的状态。
func (iterator *Iterator) Next() bool {
	if iterator.index < iterator.stack.Size() {
		iterator.index++
	}
	return iterator.stack.withinRange(iterator.index)
}

// Value 返回当前元素的值。
// 不修改迭代器的状态。
func (iterator *Iterator) Value() interface{} {
	value, _ := iterator.stack.list.Get(iterator.index) // in reverse (LIFO)
	return value
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
}

// First 将迭代器移动到第一个元素，如果容器中有第一个元素，则返回 true。
// 如果 First() 返回 true，则可以通过 Index() 和 Value() 检索第一个元素的索引和值。
// 修改迭代器的状态。
func (iterator *Iterator) First() bool {
	iterator.Begin()
	return iterator.Next()
}
