package logger

import "github.com/eduardolat/pgbackweb/internal/util/maputil"

// KV is a record of key-value pair to be logged
type KV map[string]any

// kvToArgs converts a slice of KV to a slice of any
func kvToArgs(kv ...KV) []any {
	pickedKv := KV{}
	if len(kv) > 0 {
		pickedKv = kv[0]
	}

	sortedKeys := maputil.GetSortedStringKeys(pickedKv)
	args := make([]any, 0, len(sortedKeys)*2)

	for _, k := range sortedKeys {
		args = append(args, k, pickedKv[k])
	}
	return args
}

// kvToArgsNs converts a slice of KV to a slice of any
// and adds a namespace to the resulting slice
func kvToArgsNs(ns string, kv ...KV) []any {
	pickedKv := KV{}
	if len(kv) > 0 {
		pickedKv = kv[0]
	}

	sortedKeys := maputil.GetSortedStringKeys(pickedKv)
	args := make([]any, 0, len(sortedKeys)*2+2)
	args = append(args, "ns", ns)
	for _, k := range sortedKeys {
		args = append(args, k, pickedKv[k])
	}
	return args
}
