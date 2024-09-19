package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AppendMap(t *testing.T) {
	assert.Equal(t, appendMap(), map[string]int{
		"shoes": 5,
		"phone": 10,
		"glass": 3,
	})
}
