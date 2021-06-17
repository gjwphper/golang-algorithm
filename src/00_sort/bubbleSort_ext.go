package main

import "fmt"

func main() {

	list := []int {0,12,1,0,3,2}
	moveZeroes(list)
	fmt.Println(list)

}


//方法一：双指针
//思路及解法
//
//使用双指针，左指针指向当前已经处理好的序列的尾部，右指针指向待处理序列的头部。
//
//右指针不断向右移动，每次右指针指向非零数，则将左右指针对应的数交换，同时左指针右移。
//
//注意到以下性质：
//
//左指针左边均为非零数；
//
//右指针左边直到左指针处均为零。
//
//因此每次交换，都是将左指针的零与右指针的非零数交换，且非零数的相对顺序并未改变。
//
//作者：LeetCode-Solution
//链接：https://leetcode-cn.com/problems/move-zeroes/solution/yi-dong-ling-by-leetcode-solution/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。


//复杂度分析
//
//时间复杂度：O(n)O(n)，其中 nn 为序列长度。每个位置至多被遍历两次。
//
//空间复杂度：O(1)O(1)。只需要常数的空间存放若干变量。

func moveZeroes(nums []int) {
	left, right, n := 0, 0, len(nums)
	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
		fmt.Println(nums)
	}
}



