package sorttest

import (
	"math/rand"
	"testing"
	"time"
)

func intList(n int) []int {
	list := make([]int, n)
	for i := 0; i < n; i++ {
		list[i] = i + 1
	}
	return list
}

// TestInt is a helper for testing functions that sort integer slices.
func TestInt(t *testing.T, sortFn func([]int)) {
	seed := time.Now().UnixNano()
	t.Log("Seed for random cases:", seed)
	rand.Seed(seed)

	for name, list := range map[string][]int{
		"sorted":         []int{1, 2, 3, 4},
		"reverse":        []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		"duplicates":     []int{3, 5, 3, 5, 3, 5},
		"random-len10":   rand.Perm(10),
		"random-len20":   rand.Perm(20),
		"random-len50":   rand.Perm(50),
		"random-len100":  rand.Perm(100),
		"random-len1000": rand.Perm(1000),
		// "random-len100000": rand.Perm(100000),
		// "sorted-len100000": intList(100000),
	} {
		t.Run(name, func(t *testing.T) {
			sortFn(list)
			errorCount := 0
			for i := 0; i < len(list)-1; i++ {
				if errorCount >= 5 {
					t.Fatalf("additional errors omitted for brevity")
				}
				if list[i] > list[i+1] {
					errorCount++
					t.Errorf("%d > %d; index %d", list[i], list[i+1], i)
				}
			}
		})
	}
}
