// Package arraystack Reference:https://zh.wikipedia.org/wiki/%E5%A0%86%E6%A0%88#%E9%99%A3%E5%88%97%E5%A0%86%E7%96%8A
package arraystack

import (
	"datastruct/lists/arraylist"
	"datastruct/stacks"
	"fmt"
	"strings"
)

func assertArrayStackImplementation() {
	var _ stacks.Stack = (*Stack)(nil)
}

// Stack 保存数组列表中的元素
type Stack struct {
	list *arraylist.List
}

// New 实例化一个新的空栈
func New() *Stack {
	return &Stack{list: arraylist.New()}
}

// Push 在栈顶添加一个值
func (stack *Stack) Push(value interface{}) {
	stack.list.Add(value)
}

// Pop 移除栈顶元素并返回，如果栈为空则返回 nil。第二个返回参数为真，除非堆栈是空的并且没有任何东西可以弹出。
func (stack *Stack) Pop() (interface{}, bool) {
	topElement := stack.list.Size() - 1
	value, ok := stack.list.Get(topElement)
	stack.list.Remove(topElement)
	return value, ok
}

// Peek 返回栈顶元素而不移除它，如果栈为空，则返回 nil。第二个返回参数为真，除非堆栈为空并且没有什么可查看的。
func (stack *Stack) Peek() (interface{}, bool) {
	return stack.list.Get(stack.list.Size() - 1)
}

// Empty 如果堆栈不包含任何元素，则返回 true。
func (stack *Stack) Empty() bool {
	return stack.list.Empty()
}

// Size 返回堆栈中的元素数。
func (stack *Stack) Size() int {
	return stack.list.Size()
}

// Clear 从堆栈中删除所有元素。
func (stack *Stack) Clear() {
	stack.list.Clear()
}

// Values 返回堆栈中的所有元素（LIFO 顺序）。
func (stack *Stack) Values() []interface{} {
	size := stack.list.Size()
	elements := make([]interface{}, size)
	for i := 1; i <= size; i++ {
		elements[size-i], _ = stack.list.Get(i - 1)
	}
	return elements
}

// String 返回容器的字符串表示
func (stack *Stack) String() string {
	str := "ArrayStack\n"
	var values []string
	for _, value := range stack.list.Values() {
		values = append(values, fmt.Sprintf("%v", value))
	}
	str += strings.Join(values, ", ")
	return str
}

// Check 索引在列表的范围内
func (stack *Stack) withinRange(index int) bool {
	return index >= 0 && index < stack.list.Size()
}
