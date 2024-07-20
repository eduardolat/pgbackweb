package logger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKvToArgsNoArgs(t *testing.T) {
	result := kvToArgs()
	assert.Equal(t, []any{}, result)
}

func TestKvToArgsOneArg(t *testing.T) {
	kv := KV{"key": "value"}
	result := kvToArgs(kv)
	assert.Equal(t, []any{"key", "value"}, result)
}

func TestKvToArgsMultipleArgs(t *testing.T) {
	kv1 := KV{"key1": "value1", "key2": "value2"}
	kv2 := KV{"key3": "value3"}
	result := kvToArgs(kv1, kv2)
	assert.Equal(t, []any{"key1", "value1", "key2", "value2"}, result)
}

func TestKvToArgsNsNoArgs(t *testing.T) {
	result := kvToArgsNs("namespace")
	assert.Equal(t, []any{"ns", "namespace"}, result)
}

func TestKvToArgsNsOneArg(t *testing.T) {
	kv := KV{"key": "value"}
	result := kvToArgsNs("namespace", kv)
	assert.Equal(t, []any{"ns", "namespace", "key", "value"}, result)
}

func TestKvToArgsNsMultipleArgs(t *testing.T) {
	kv1 := KV{"key1": "value1", "key2": "value2"}
	kv2 := KV{"key3": "value3"}
	result := kvToArgsNs("namespace", kv1, kv2)
	assert.Equal(t, []any{"ns", "namespace", "key1", "value1", "key2", "value2"}, result)
}

func TestKvToArgsPickOnlyFirst(t *testing.T) {
	kv1 := KV{"key1": "value1"}
	kv2 := KV{"key2": "value2"}
	result := kvToArgs(kv1, kv2)
	assert.Equal(t, []any{"key1", "value1"}, result)
}

func TestKvToArgsNsPickOnlyFirst(t *testing.T) {
	kv1 := KV{"key1": "value1"}
	kv2 := KV{"key2": "value2"}
	result := kvToArgsNs("namespace", kv1, kv2)
	assert.Equal(t, []any{"ns", "namespace", "key1", "value1"}, result)
}

func TestKvToArgsOrder(t *testing.T) {
	kv := KV{"z": "value1", "a": "value2"}
	result := kvToArgs(kv)
	assert.Equal(t, []any{"a", "value2", "z", "value1"}, result)
}

func TestKvToArgsNsOrder(t *testing.T) {
	kv := KV{"z": "value1", "a": "value2"}
	result := kvToArgsNs("namespace", kv)
	assert.Equal(t, []any{"ns", "namespace", "a", "value2", "z", "value1"}, result)
}
