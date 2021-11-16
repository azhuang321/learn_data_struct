package singlylinkedlist

import (
	"datastruct/containers"
	"encoding/json"
)

func assertSerializationImplementation() {
	var _ containers.JSONSerializer = (*List)(nil)
	var _ containers.JSONDeserializer = (*List)(nil)
}

// ToJSON 输出列表元素的 JSON 表示。
func (list *List) ToJSON() ([]byte, error) {
	return json.Marshal(list.Values())
}

// FromJSON 从输入的 JSON 表示中填充列表的元素。
func (list *List) FromJSON(data []byte) error {
	var elements []interface{}
	err := json.Unmarshal(data, &elements)
	if err == nil {
		list.Clear()
		list.Add(elements...)
	}
	return err
}
