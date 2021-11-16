package hashbidimap

import (
	"encoding/json"

	"datastruct/containers"
)

func assertSerializationImplementation() {
	var _ containers.JSONSerializer = (*Map)(nil)
	var _ containers.JSONDeserializer = (*Map)(nil)
}

// ToJSON 输出map的 JSON 表示。
func (m *Map) ToJSON() ([]byte, error) {
	return m.forwardMap.ToJSON()
}

// FromJSON 从输入的 JSON 表示填充map。
func (m *Map) FromJSON(data []byte) error {
	elements := make(map[string]interface{})
	err := json.Unmarshal(data, &elements)
	if err == nil {
		m.Clear()
		for key, value := range elements {
			m.Put(key, value)
		}
	}
	return err
}
