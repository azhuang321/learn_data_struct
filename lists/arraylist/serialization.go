package arraylist

import (
	"datastruct/containers"
	"encoding/json"
)

func assertSerializationImplementation() {
	var _ containers.JSONSerializer = (*List)(nil)
	var _ containers.JSONDeserializer = (*List)(nil)
}

func (list *List) ToJSON() ([]byte, error) {
	return json.Marshal(list.elements[:list.size])
}

func (list *List) FromJSON(data []byte) error {
	err := json.Unmarshal(data, &list.elements)
	if err == nil {
		list.size = len(list.elements)
	}
	return err
}
