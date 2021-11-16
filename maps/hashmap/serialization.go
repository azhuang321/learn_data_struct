package hashmap

import (
	"encoding/json"

	"datastruct/containers"
	"datastruct/utils"
)

func assertSerializationImplementation() {
	var _ containers.JSONSerializer = (*Map)(nil)
	var _ containers.JSONDeserializer = (*Map)(nil)
}

// ToJSON 输出地图的 JSON 表示。
func (m *Map) ToJSON() ([]byte, error) {
	elements := make(map[string]interface{})
	for key, value := range m.m {
		elements[utils.ToString(key)] = value
	}
	return json.Marshal(&elements)
}

// FromJSON 从输入的 JSON 表示填充地图。
func (m *Map) FromJSON(data []byte) error {
	elements := make(map[string]interface{})
	err := json.Unmarshal(data, &elements)
	if err == nil {
		m.Clear()
		for key, value := range elements {
			m.m[key] = value
		}
	}
	return err
}
