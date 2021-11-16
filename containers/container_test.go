package containers

import (
	"datastruct/utils"
	"testing"
)

type ContainerTest struct {
	values []interface{}
}

func (container ContainerTest) Empty() bool {
	return len(container.values) == 0
}

func (container ContainerTest) Size() int {
	return len(container.values)
}

func (container ContainerTest) Clear() {
	container.values = make([]interface{}, 0)
}

func (container ContainerTest) Values() []interface{} {
	return container.values
}

func TestGetSortedValuesInts(t *testing.T) {
	container := ContainerTest{}
	container.values = []interface{}{5, 1, 3, 2, 4}
	values := GetSortedValues(container, utils.IntComparator)
	for i := 1; i < container.Size(); i++ {
		if values[i-1].(int) > values[i].(int) {
			t.Errorf("Not sorted!")
		}
	}
}

func TestGetSortedValuesStrings(t *testing.T) {
	container := ContainerTest{}
	container.values = []interface{}{"g", "a", "d", "e", "f", "c", "b"}
	values := GetSortedValues(container, utils.StringComparator)
	for i := 1; i < container.Size(); i++ {
		if values[i-1].(string) > values[i].(string) {
			t.Errorf("Not sorted!")
		}
	}
}
