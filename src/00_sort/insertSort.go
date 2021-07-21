package main

import "fmt"

func main() {

	list := []int{6, 3, 5, 2, 4, 1}
	insertSort(list)
	fmt.Println(list)

}

func insertSort(list []int) {
	// 从第二个数开始，往前插入数字
	for i := 1; i < len(list); i++ {
		// j 记录当前数字下标
		for j := i; j >= i; j-- {
			if list[j] < list[j-1] {
				insertSwap(list, j, j-1)
			}
		}
	}
}

func insertSwap(list []int, m int, n int) {
	temp := list[m]
	list[m] = list[n]
	list[n] = temp
}
