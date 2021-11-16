package linkedliststack

import "datastruct/containers"

func assertSerializationImplementation() {
	var _ containers.JSONSerializer = (*Stack)(nil)
	var _ containers.JSONDeserializer = (*Stack)(nil)
}

// ToJSON 输出堆栈的 JSON 表示。
func (stack *Stack) ToJSON() ([]byte, error) {
	return stack.list.ToJSON()
}

// FromJSON 从输入的 JSON 表示填充堆栈。
func (stack *Stack) FromJSON(data []byte) error {
	return stack.list.FromJSON(data)
}
