package util_test

import (
	"testing"

	"github.com/slavamuravey/wbppc/pkg/assert"
	"github.com/slavamuravey/wbppc/pkg/util"
)

func TestCreateIndex(t *testing.T) {
	data := [][]string{
		{"a", "b", "c"},
		{"d", "e"},
		{"f"},
		{"g", "h", "i"},
	}

	a := util.CreateIndex(data, 1)

	e := map[string][]string{
		"b": {"a", "b", "c"},
		"e": {"d", "e"},
		"h": {"g", "h", "i"},
	}

	assert.Equal(t, e, a, "indices should be equal")
}
