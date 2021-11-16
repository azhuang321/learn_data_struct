package containers

// JSONSerializer 提供 JSON 序列化
type JSONSerializer interface {
	// ToJSON 输出容器元素的 JSON 表示。
	ToJSON() ([]byte, error)
}

// JSONDeserializer 提供 JSON 反序列化
type JSONDeserializer interface {
	// FromJSON 从输入的 JSON 表示中填充容器的元素。
	FromJSON([]byte) error
}
