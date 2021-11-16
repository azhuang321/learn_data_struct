// Package 容器为数据结构提供核心接口和功能。
//
// Container 是所有要实现的数据结构的基本接口。
//
// Iterators 提供有状态的迭代器。
//
// Enumerable 提供受 Ruby 启发的（each、select、map、find、any? 等）容器函数。
//
// Serialization 提供序列化器（marshalers）和反序列化器（unmarshalers）。
package containers

import "datastruct/utils"

type Container interface {
	Empty() bool
	Size() int
	Clear()
	Values() []interface{}
}

// GetSortedValues 返回相对于传递的比较器排序的容器元素。不影响容器内元素的顺序。
func GetSortedValues(container Container, comparator utils.Comparator) []interface{} {
	values := container.Values()
	if len(values) < 2 {
		return values
	}
	utils.Sort(values, comparator)
	return values
}
