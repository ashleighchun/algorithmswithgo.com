package module02

import (
	"algo/module02/sorttest"
	"testing"
)

// BubbleSortInt will sort a list of integers using the bubble sort algorithm.
func TestBubbleSortInt(t *testing.T) {
	sorttest.TestInt(t, BubbleSortInt)
}
