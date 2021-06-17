package main

import (
	"fmt"
)

func main() {
	list := []int {6,3,5,2,4,1}
	largest := findKthLargest(list, 2)
	fmt.Println(largest)

}

func findKthLargest(nums []int,k int) int {

	max_index := 0
	for i := 0; i < len(nums) - 1; i++  {
		max_index = i
		for j := i + 1; j < len(nums); j++ {
			if (nums[j] > nums[max_index]) {
				max_index = j
			}
		}

		temp := nums[i]
		nums[i] = nums[max_index]
		nums[max_index]  = temp

		if i == k - 1 {
			break
		}
	}
	return nums[k - 1]


}





