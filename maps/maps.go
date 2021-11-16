// Reference: https://zh.wikipedia.org/wiki/%E5%85%B3%E8%81%94%E6%95%B0%E7%BB%84
package maps

import "datastruct/containers"

// Map 所有 map 实现的接口
type Map interface {
	Put(key interface{}, value interface{})
	Get(key interface{}) (value interface{}, found bool)
	Remove(key interface{})
	Keys() []interface{}

	containers.Container
}

// BidiMap 所有双向映射实现的接口（扩展 Map 接口）
type BidiMap interface {
	GetKey(value interface{}) (key interface{}, found bool)

	Map
}
