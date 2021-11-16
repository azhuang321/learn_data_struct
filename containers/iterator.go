package containers

// IteratorWithIndex 是有序容器的有状态迭代器，其值可以通过索引获取。
type IteratorWithIndex interface {
	// Next 将迭代器移动到下一个元素，如果容器中有下一个元素，则返回 true。如果 Next() 返回 true，则可以通过 Index() 和 Value() 检索下一个元素的索引和值。如果 Next() 是第一次调用，那么它会将迭代器指向第一个元素（如果存在）。修改迭代器的状态。
	Next() bool

	// Value 返回当前元素的值。不修改迭代器的状态。
	Value() interface{}

	// Index 返回当前元素的索引。不修改迭代器的状态。
	Index() int

	// Begin 将迭代器重置为其初始状态（一个在第一个之前）调用 Next() 以获取第一个元素（如果有）。
	Begin()

	// First 将迭代器移动到第一个元素，如果容器中有第一个元素，则返回 true。如果 First() 返回 true，则可以通过 Index() 和 Value() 检索第一个元素的索引和值。修改迭代器的状态。
	First() bool
}

// IteratorWithKey 是元素为键值对的有序容器的有状态迭代器。
type IteratorWithKey interface {
	// Next 将迭代器移动到下一个元素，如果容器中有下一个元素，则返回 true。如果 Next() 返回 true，则可以通过 Key() 和 Value() 检索下一个元素的键和值。如果 Next() 是第一次调用，那么它会将迭代器指向第一个元素（如果存在）。修改迭代器的状态。
	Next() bool

	// Value 返回当前元素的值。不修改迭代器的状态。
	Value() interface{}

	// Key 返回当前元素的键。不修改迭代器的状态。
	Key() interface{}

	// Begin 将迭代器重置为其初始状态（一个在第一个之前）调用 Next() 以获取第一个元素（如果有）。
	Begin()

	// First 将迭代器移动到第一个元素，如果容器中有第一个元素，则返回 true。如果 First() 返回 true，则 Key() 和 Value() 可以检索第一个元素的键和值。修改迭代器的状态。
	First() bool
}

// ReverseIteratorWithIndex 是有序容器的有状态迭代器，其值可以通过索引获取。
//
// 本质上它与 IteratorWithIndex 相同，但提供了额外的:
//
// Prev() 启用反向遍历的函数
//
// Last() 函数将迭代器移动到最后一个元素。
//
// End() 函数将迭代器移过最后一个元素（one-past-the-end）。
type ReverseIteratorWithIndex interface {
	// Prev 将迭代器移动到前一个元素，如果容器中有前一个元素，则返回 true。如果 Prev() 返回 true，则可以通过 Index() 和 Value() 检索前一个元素的索引和值。修改迭代器的状态。
	Prev() bool

	// End 将迭代器移过最后一个元素（one-past-the-end）。调用 Prev() 以获取最后一个元素（如果有）。
	End()

	// Last 将迭代器移动到最后一个元素，如果容器中有最后一个元素，则返回 true。如果 Last() 返回 true，则可以通过 Index() 和 Value() 检索最后一个元素的索引和值。修改迭代器的状态。
	Last() bool

	IteratorWithIndex
}

// ReverseIteratorWithKey 是元素为键值对的有序容器的有状态迭代器。
//
// 本质上它与 IteratorWithKey 相同，但提供了额外的:
//
// Prev() 启用反向遍历的函数
//
// Last() 函数将迭代器移动到最后一个元素。
type ReverseIteratorWithKey interface {
	// Prev 将迭代器移动到前一个元素，如果容器中有前一个元素，则返回 true。如果 Prev() 返回 true，则可以通过 Key() 和 Value() 检索前一个元素的键和值。修改迭代器的状态。
	Prev() bool

	// End 将迭代器移过最后一个元素（one-past-the-end）。调用 Prev() 以获取最后一个元素（如果有）。
	End()

	// Last 将迭代器移动到最后一个元素，如果容器中有最后一个元素，则返回 true。如果 Last() 返回 true，则可以通过 Key() 和 Value() 检索最后一个元素的键和值。修改迭代器的状态。
	Last() bool

	IteratorWithKey
}
