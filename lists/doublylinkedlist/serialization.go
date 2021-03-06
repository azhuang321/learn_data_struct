package doublylinkedlist

import (
	"datastruct/containers"
	"encoding/json"
)

func assertSerializationImplementation() {
	var _ containers.JSONSerializer = (*List)(nil)
	var _ containers.JSONDeserializer = (*List)(nil)
}

func (list *List) ToJSON() ([]byte, error) {
	return json.Marshal(list.Values())
}

func (list *List) FromJSON(data []byte) error {
	var element []interface{}
	err := json.Unmarshal(data, &element)
	if err == nil {
		list.Clear()
		list.Add(element...)
	}
	return err
}
