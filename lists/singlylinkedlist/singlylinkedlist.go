// Package singlylinkedlist Reference:https://zh.wikipedia.org/wiki/%E5%8D%95%E5%90%91%E9%93%BE%E8%A1%A8
package singlylinkedlist

import (
	"datastruct/lists"
	"datastruct/utils"
	"fmt"
	"strings"
)

func assertSinglyLinkedListImplementation() {
	var _ lists.List = (*List)(nil)
}

// List 保存元素，其中每个元素都指向下一个元素
type List struct {
	first *element
	last  *element
	size  int
}

type element struct {
	value interface{}
	next  *element
}

// New 实例化一个新列表并将传递的值（如果有）添加到列表中
func New(values ...interface{}) *List {
	list := &List{}
	if len(values) > 0 {
		list.Add(values...)
	}
	return list
}

// Add 在列表末尾附加一个值（一个或多个）（与 Append() 相同）
func (list *List) Add(values ...interface{}) {
	for _, value := range values {
		//fmt.Println("value",value,"size",list.size,"list",list.Values())
		newElement := &element{value: value}
		if list.size == 0 {
			list.first = newElement
			list.last = newElement
		} else {
			list.last.next = newElement
			list.last = newElement
		}
		list.size++
	}
}

// Append 在列表末尾附加一个值（一个或多个）（与 Add() 相同）
func (list *List) Append(values ...interface{}) {
	list.Add(values...)
}

// Prepend 前置一个值（或更多）
func (list *List) Prepend(values ...interface{}) {
	newList := New(values...)
	newList.last.next = list.first
	list.first = newList.first
	list.size += len(values)
}

// Get 返回索引处的元素。
// 如果索引在数组范围内且数组不为空，则第二个返回参数为真，否则为假。
func (list *List) Get(index int) (interface{}, bool) {
	if !list.withinRange(index) {
		return nil, false
	}
	element := list.first
	for e := 0; e != index; e, element = e+1, element.next {
	}
	return element.value, true
}

// Remove 从列表中删除给定索引处的元素。
func (list *List) Remove(index int) {
	if !list.withinRange(index) {
		return
	}

	if list.size == 1 {
		list.Clear()
		return
	}

	foundElement, preElement := list.first, list.first
	for e := 0; e != index; e, preElement, foundElement = e+1, foundElement, foundElement.next {
	}

	if foundElement == list.first {
		list.first = list.first.next
	} else {
		preElement.next = foundElement.next
		foundElement = nil
	}

	list.size--
}

// Contains 检查集合中是否存在值（一个或多个）。
// 所有值都必须存在于集合中，方法才能返回 true。
// n^2 的性能时间复杂度。
// 如果根本没有传递任何参数，则返回 true，即 set 始终是空集的超集。
func (list *List) Contains(values ...interface{}) bool {
	if len(values) == 0 {
		return true
	}
	if list.size == 0 {
		return false
	}
	for _, value := range values {
		found := false
		for element := list.first; element != nil; element = element.next {
			if element.value == value {
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
	values := make([]interface{}, list.size)
	for e, element := 0, list.first; element != nil; e, element = e+1, element.next {
		values[e] = element.value
	}
	return values
}

//IndexOf 返回所提供元素的索引
func (list *List) IndexOf(value interface{}) int {
	if list.size == 0 {
		return -1
	}
	for e, element := 0, list.first; element != nil; e, element = e+1, element.next {
		if element.value == value {
			return e
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
	list.first = nil
	list.last = nil
	list.size = 0
}

// Sort 排序值（就地）使用。
func (list *List) Sort(comparator utils.Comparator) {
	if list.size < 2 {
		return
	}
	values := list.Values()
	utils.Sort(values, comparator)
	list.Clear()
	list.Add(values...)
}

// Swap 交换给定索引处两个元素的值。
func (list *List) Swap(i, j int) {
	if !list.withinRange(i) || !list.withinRange(j) || i == j {
		return
	}
	var iElement, jElement *element
	for e, currentElement := 0, list.first; iElement == nil || jElement == nil; e, currentElement = e+1, currentElement.next {
		if e == i {
			iElement = currentElement
		}
		if e == j {
			jElement = currentElement
		}
	}
	iElement.value, jElement.value = jElement.value, iElement.value
}

// Insert 在指定的索引位置插入值，将该位置的值（如果有）和任何后续元素向右移动。
// 如果位置为负或大于列表大小，则不执行任何操作
// 注意：位置等于列表的大小是有效的，即追加。
func (list *List) Insert(index int, value ...interface{}) {
	if !list.withinRange(index) {
		if index == list.size {
			list.Add(value...)
		}
		return
	}
	currentElement, preElement := list.first, list.first
	for e := 0; index != e; e, preElement, currentElement = e+1, currentElement, currentElement.next {
	}
	newList := New(value...)
	if currentElement == list.first {
		newList.last.next = list.first
		list.first = newList.first
	} else {
		preElement.next = newList.first
		newList.last.next = currentElement
	}
	list.size += len(value)
}

// Set 指定索引处的值
// 如果位置为负或大于列表大小，则不执行任何操作
// 注意：位置等于列表的大小是有效的，即追加。
func (list *List) Set(index int, value interface{}) {
	if !list.withinRange(index) {
		if index == list.size {
			list.Add(value)
		}
		return
	}
	currentElement := list.first
	for e := 0; index != e; e, currentElement = e+1, currentElement.next {
	}
	currentElement.value = value
}

// String 返回容器的字符串表示
func (list *List) String() string {
	str := "SinglyLinkedList\n"
	var values []string
	for element := list.first; element != nil; element = element.next {
		values = append(values, fmt.Sprintf("%v", element.value))
	}
	str += strings.Join(values, ", ")
	return str
}

// Check 索引在列表的范围内
func (list *List) withinRange(index int) bool {
	return index >= 0 && index < list.size
}
