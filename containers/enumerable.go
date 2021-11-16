package containers

// EnumerableWithIndex 为可以通过索引获取值的有序容器提供函数。
type EnumerableWithIndex interface {
	// Each 为每个元素调用一次给定的函数，传递该元素的索引和值。
	Each(func(index int, value interface{}))

	// Map 为每个元素调用一次给定的函数，并返回一个包含给定函数返回值的容器。
	// TODO 非常感谢有关如何在容器中强制执行此操作的帮助（链接时不想键入断言）
	// Map(func(index int, value interface{}) interface{}) Container

	// Select 返回一个新容器，其中包含给定函数返回真值的所有元素。
	// TODO 需要有关如何在容器中强制执行此操作的帮助（链接时不想键入断言）
	// Select(func(index int, value interface{}) bool) Container

	// Any 将容器的每个元素传递给给定的函数，如果函数对任何元素返回 true，则返回 true。
	Any(func(index int, value interface{}) bool) bool

	// All 将容器的每个元素传递给给定的函数，如果该函数对所有元素都返回 true，则返回 true。
	All(func(index int, value interface{}) bool) bool

	// Find 将容器的每个元素传递给给定的函数，并返回函数为真的第一个 (index,value) 或 -1,nil，否则如果没有元素符合条件。
	Find(func(index int, value interface{}) bool) (int, interface{})
}

// EnumerableWithKey 为元素为键值对的值的有序容器提供函数。
type EnumerableWithKey interface {
	// Each calls the given function once for each element, passing that element's key and value.
	Each(func(key interface{}, value interface{}))

	// Map 为每个元素调用给定函数一次，并返回一个容器，其中包含给定函数返回的值作为键值对。
	// TODO 需要有关如何在容器中强制执行此操作的帮助（链接时不想键入断言）
	// Map(func(key interface{}, value interface{}) (interface{}, interface{})) Container

	// Select 返回一个新容器，其中包含给定函数返回真值的所有元素。
	// TODO 需要有关如何在容器中强制执行此操作的帮助（链接时不想键入断言）
	// Select(func(key interface{}, value interface{}) bool) Container

	// Any 将容器的每个元素传递给给定的函数，如果函数对任何元素返回 true，则返回 true。
	Any(func(key interface{}, value interface{}) bool) bool

	// All 将容器的每个元素传递给给定的函数，如果该函数对所有元素都返回 true，则返回 true。
	All(func(key interface{}, value interface{}) bool) bool

	// Find 将容器的每个元素传递给给定的函数，并返回函数为 true 或 nil,nil 的第一个 (key,value)，否则如果没有元素符合条件。
	Find(func(key interface{}, value interface{}) bool) (interface{}, interface{})
}
