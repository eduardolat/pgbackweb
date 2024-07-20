package maputil

import "sort"

// GetSortedStringKeys returns the keys of a map sorted in lexicographical order.
func GetSortedStringKeys[T any](m map[string]T) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
