// Package linkedliststack Reference:https://zh.wikipedia.org/wiki/%E5%A0%86%E6%A0%88#%E4%B8%B2%E5%88%97%E5%A0%86%E7%96%8A
package linkedliststack

import (
	"fmt"
	"strings"

	"datastruct/lists/singlylinkedlist"
	"datastruct/stacks"
)

func assertStackImplementation() {
	var _ stacks.Stack = (*Stack)(nil)
}

// Stack 将元素保存在单向链表中
type Stack struct {
	list *singlylinkedlist.List
}

// New new实例化一个新的空栈
func New() *Stack {
	return &Stack{list: &singlylinkedlist.List{}}
}

// Push 在栈顶添加一个值
func (stack *Stack) Push(value interface{}) {
	stack.list.Prepend(value)
}

// Pop 移除栈顶元素并返回，如果栈为空则返回 nil。
// 第二个返回参数为真，除非堆栈是空的并且没有任何东西可以弹出。
func (stack *Stack) Pop() (value interface{}, ok bool) {
	value, ok = stack.list.Get(0)
	stack.list.Remove(0)
	return
}

// Peek 返回栈顶元素而不移除它，如果栈为空，则返回 nil。
// 第二个返回参数为真，除非堆栈为空并且没有什么可查看的。
func (stack *Stack) Peek() (value interface{}, ok bool) {
	return stack.list.Get(0)
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
	return stack.list.Values()
}

// String 返回容器的字符串表示
func (stack *Stack) String() string {
	str := "LinkedListStack\n"
	var values []string
	for _, value := range stack.list.Values() {
		values = append(values, fmt.Sprintf("%v", value))
	}
	str += strings.Join(values, ", ")
	return str
}

// withinRange 检查索引是否在列表范围内
func (stack *Stack) withinRange(index int) bool {
	return index >= 0 && index < stack.list.Size()
}
