package orderedmap

type OrderedMap[V any] struct {
	keys []string
	data map[string]V
}

func NewOrderedMap[V any]() *OrderedMap[V] {
	return &OrderedMap[V]{
		keys: make([]string, 0),
		data: make(map[string]V),
	}
}

func (m *OrderedMap[V]) Put(key string, value V) {
	if _, ok := m.data[key]; !ok {
		m.keys = append(m.keys, key)
	}
	m.data[key] = value
}

func (m *OrderedMap[V]) Get(key string) *V {
	val, ok := m.data[key]
	if !ok {
		return nil
	}
	return &val
}
func (m *OrderedMap[V]) GetAsValue(key string) (V, bool) {
	val, ok := m.data[key]
	return val, ok
}

func (m *OrderedMap[V]) Keys() []string {
	return m.keys
}
