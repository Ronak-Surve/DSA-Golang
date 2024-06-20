// optimized O(N)
package main

import (
	"fmt"
	"math"
)

func findMaxAverage(nums []int, k int) float64 {

	maxAvg := 0.0
	sum := 0.0

	for i := 0; i < len(nums); i++ {

		sum = sum + float64(nums[i])

		if i >= k-1 {

			avg := sum / float64(k)
			if maxAvg == 0 {
				maxAvg = avg
			}
			maxAvg = math.Max(avg, maxAvg)
			sum = sum - float64(nums[i-k+1])
		}
	}

	return maxAvg
}

func main() {

	var arr = []int{1, 12, -5, -6, 50, 3}
	var k int = 4
	var result float64 = findMaxAverage(arr, k)
	fmt.Println(result)
}
