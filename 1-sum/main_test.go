package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Sum(t *testing.T) {
	assert.Equal(t, sum([]int{1, 2, 3, 4, 5}), 15)

	assert.Equal(t, sum([]int{-1, 2, 3, 4, 5}), 13)

	assert.Equal(t, sum([]int{20, 20, 20, 20, 20}), 100)

	assert.Equal(t, sum([]int{20, 20, -20, 20, 20}), 60)
}
