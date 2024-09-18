package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GreatestElement(t *testing.T) {
	assert.Equal(t, greatestElement([]int{1, 2, 3, 4, 5}), 5)

	assert.Equal(t, greatestElement([]int{1, 3, 7, 13, 2, 1}), 13)

	assert.Equal(t, greatestElement([]int{20, 20, 20, 20, 20}), 20)

	assert.Equal(t, greatestElement([]int{-1, -5, -10, -13, -2, -4}), -1)
}
