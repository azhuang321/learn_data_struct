// Package arraylist Reference:https://zh.wikipedia.org/wiki/%E9%93%BE%E8%A1%A8
package arraylist

import (
	"datastruct/lists"
	"datastruct/utils"
	"fmt"
	"strings"
)

func assertListImplementation() {
	var _ lists.List = (*List)(nil)
}

// List 将元素保存在切片中
type List struct {
	elements []interface{}
	size     int
}

const (
	growthFactor = float32(2.0)  // 100% 时增长
	shrinkFactor = float32(0.25) //当大小为容量的 25% 时收缩（0 表示永不收缩）
)

// New 实例化一个新列表并将传递的值（如果有）添加到列表中
func New(values ...interface{}) *List {
	list := new(List)
	if len(values) > 0 {
		list.Add(values...)
	}
	return list
}

// Add 在列表末尾附加一个值
func (list *List) Add(values ...interface{}) {
	list.growBy(len(values))
	for _, value := range values {
		list.elements[list.size] = value
		list.size++
	}
}

// Get 返回索引处的元素。
// 如果索引在数组范围内且数组不为空，则第二个返回参数为真，否则为假。
func (list *List) Get(index int) (interface{}, bool) {
	if !list.withinRange(index) {
		return nil, false
	}
	return list.elements[index], true
}

// Remove 从列表中删除给定索引处的元素。
func (list *List) Remove(index int) {
	if !list.withinRange(index) {
		return
	}
	list.elements[index] = nil                                    // 清理引用
	copy(list.elements[index:], list.elements[index+1:list.size]) //左移一位（操作慢，需要优化的方法）
	list.size--

	list.shrink()
}

// Contains 检查集合中是否存在元素（一个或多个）。
//所有元素都必须存在于集合中，方法才能返回 true。
//n^2 的性能时间复杂度。
//如果根本没有传递任何参数，则返回 true，即 set 始终是空集的超集。
func (list *List) Contains(values ...interface{}) bool {
	for _, searchVal := range values {
		found := false
		for _, element := range list.elements {
			if element == searchVal {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

// Values 返回列表中的所有元素。
func (list *List) Values() []interface{} {
	allVal := make([]interface{}, list.size)
	copy(allVal, list.elements[:list.size])
	return allVal
}

//IndexOf 返回所提供元素的索引
func (list *List) IndexOf(value interface{}) int {
	if list.size == 0 {
		return -1
	}

	for i, element := range list.elements {
		if element == value {
			return i
		}
	}
	return -1
}

// Empty 如果列表不包含任何元素，则返回 true。
func (list *List) Empty() bool {
	return list.size == 0
}

// Size 返回列表中的元素数。
func (list *List) Size() int {
	return list.size
}

// Clear 从列表中删除所有元素。
func (list *List) Clear() {
	list.size = 0
	list.elements = []interface{}{}
}

// Sort 使用排序值（就地）。
func (list *List) Sort(comparator utils.Comparator) {
	if len(list.elements) < 2 {
		return
	}
	utils.Sort(list.elements[:list.size], comparator)
}

// Swap 交换指定位置的两个值。
func (list *List) Swap(i, j int) {
	if list.withinRange(i) && list.withinRange(j) {
		list.elements[i], list.elements[j] = list.elements[j], list.elements[i]
	}
}

// Insert 在指定的索引位置插入值，将该位置的值（如果有）和任何后续元素向右移动。
//如果位置为负或大于列表的大小，则不执行任何操作
//注意：位置等于列表的大小是有效的，即追加。
func (list *List) Insert(index int, values ...interface{}) {
	if !list.withinRange(index) {
		//append
		if index == list.size {
			list.Add(values...)
		}
		return
	}

	l := len(values)
	list.growBy(l)
	list.size += l
	copy(list.elements[index+l:], list.elements[index:list.size-1])
	copy(list.elements[index:], values)
}

// Set 如果位置为负或大于列表的大小，则指定索引处的值不执行任何操作
//注意：位置等于列表的大小是有效的，即追加。
func (list *List) Set(index int, value interface{}) {
	if !list.withinRange(index) {
		//append
		if index == list.size {
			list.Add(value)
		}
		return
	}
	list.elements[index] = value
}

func (list *List) String() string {
	str := "ArrayList\n"
	values := make([]string, list.size)
	for _, value := range list.elements[:list.size] {
		values = append(values, fmt.Sprintf("%v", value))
	}
	str += strings.Join(values, ", ")
	return str
}

// withinRange 检查索引是否在列表范围内
func (list *List) withinRange(index int) bool {
	return index >= 0 && index < list.size
}

// growBy 如有必要，扩展数组，即如果我们添加 n 个元素将达到容量
func (list *List) growBy(n int) {
	currentCapacity := cap(list.elements)
	// 当达到容量时，以增长因子的因子增长并添加元素数量
	if list.size+n >= currentCapacity {
		newCapacity := int(growthFactor * float32(currentCapacity+n))
		list.resize(newCapacity)
	}
}

func (list *List) resize(cap int) {
	newElements := make([]interface{}, cap, cap)
	copy(newElements, list.elements)
	list.elements = newElements
}

// Shrink 如有必要，数组，即当大小为当前容量的 shrinkFactor 百分比时
func (list *List) shrink() {
	if shrinkFactor == 0.0 {
		return
	}
	// 当尺寸处于收缩因子容量时收缩
	currentCapacity := cap(list.elements)
	if list.size <= int(float32(currentCapacity)*shrinkFactor) {
		list.resize(list.size)
	}
}
