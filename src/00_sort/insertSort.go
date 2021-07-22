package main

import "fmt"

//相邻两个交互，有点类似冒泡
func main() {

	list := []int{6, 3, 5, 2, 4, 1}
	insertSort(list)
	fmt.Println(list)

}

func insertSort(list []int) {
	// 从第二个数开始，往前插入数字
	for i := 1; i < len(list); i++ {
		// j 记录当前数字下标
		var j = i
		fmt.Println(list)
		for j >= 1 && list[j] < list[j-1] {

			insertSwap(list, j, j-1)
			j--

		}
	}
}

func insertSwap(list []int, m int, n int) {
	temp := list[m]
	list[m] = list[n]
	list[n] = temp
}
