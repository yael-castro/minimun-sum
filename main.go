package main

import (
	"math"
	"sort"
)

// minSum
//
// Complexity O(k * n)
//
// n - is the length of num
func minSum(num []int32, k int32) (sum int32) {
	if len(num) == 0 {
		return
	}

	sort.SliceStable(num, func(x, y int) bool {
		return num[x] > num[y]
	})

	sortNum := func() {
		for i := 0; i < len(num)-1; i++ {
			if num[i] < num[i+1] {
				num[i], num[i+1] = num[i+1], num[i]
				continue
			}
			break
		}
	}

	for k > 0 {
		if len(num) > 1 {
			if num[0] < num[1] { // Check if numbers are not ordered
				sortNum()
			}
		}

		num[0] = int32(math.Ceil(float64(num[0]) / 2.0))
		k--
	}

	for _, v := range num {
		sum += v
	}

	return
}
