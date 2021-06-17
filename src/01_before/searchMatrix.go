package main

import (
	"fmt"
	"sort"
)
/**
编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target 。该矩阵具有以下特性：

每行的元素从左到右升序排列。
每列的元素从上到下升序排列。

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions/xmlwi1/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
 */
func main() {

	outMatrix := [][]int {{1,2,3},{4,5,6},{7,8,9}}
	matrix1 := searchMatrix1(outMatrix, 9)
	matrix2 := searchMatrix2(outMatrix, 10)
	fmt.Println(matrix1)
	fmt.Println(matrix2)

}

func searchMatrix1(matrix [][]int,target int)  bool{
	var all []int
	for _,ints := range matrix{
		all = append(all,ints...)
	}


	sort.Ints(all)
	idx := sort.SearchInts(all,target)
	if idx < len(all) && all[idx] == target{
		return true
	}
	return false
}



func searchMatrix2(matrix [][]int,target int)  bool {
	row := len(matrix) - 1
	col := 0

	for row >= 0 && col < len(matrix[0]) {

		if matrix[row][col] > target {
			row--
		} else if (matrix[row][col] < target) {
			col++
		} else {
			return true
		}
	}


	return false

}
