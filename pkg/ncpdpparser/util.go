package ncpdpparser

import (
	"github.com/emirpasic/gods/maps/linkedhashmap"
)

func GetFromMap(p linkedhashmap.Map, key string) string {
	fieldVal, ok := p.Get(key)
	if !ok {
		return ""
	}
	return fieldVal.(string)
}

func GetMapKeys(v map[string]string) []string {
	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}
	return keys
}

func ContainsInArray(arr []string, str string) int {

	for i, a := range arr {
		if a == str {
			return i
		}
	}
	return -1
}
