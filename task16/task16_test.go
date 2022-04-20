package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSort(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	QuickSort(data)
	assert.Equal(t, []int{1, 2, 3, 4, 5}, data)

	data = []int{4, 3, 2, 1}
	QuickSort(data)
	assert.Equal(t, []int{1, 2, 3, 4}, data)

	data = []int{3, 3, 0, 0, 5, 5, -123, -123, -44, 545}
	QuickSort(data)
	assert.Equal(t, []int{-123, -123, -44, 0, 0, 3, 3, 5, 5, 545}, data)

	data = []int{1, 1, 1, 1}
	QuickSort(data)
	assert.Equal(t, []int{1, 1, 1, 1}, data)
}
