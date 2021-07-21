package main

import "fmt"

func main() {

	list := []int{6, 3, 5, 2, 4, 1}

	// i = 0 第一遍排序
	// 6,3,5,2,4,1
	// 3,6,5,2,4,1
	// 3,5,6,2,4,1
	// 3,5,2,6,4,1
	// 3,5,2,4,6,1
	// 3,5,2,4,1,6
	// i = 1
	// 、、、第二遍
	bubbleSort(list)
	fmt.Println(list)

	// 不引入变量，交换变量值
	a := 3
	b := 4
	//b = b + a
	//a = b - a
	//b = b - a

	a = a ^ b

	//011   3   = 1*2的0次方 + 1*2的一次方
	//100	4 = 1*2的2次方
	//111

	b = b ^ a
	a = a ^ b

	c := 3
	d := 4

	//c = c & d//0

	//c = c | d//7

	c = ^1

	fmt.Println(a, b, c)

	fmt.Println(d)

}

/**
1、2层循环，里面的作比较
2、list构建
3、交换
*/
func bubbleSort(list []int) {

	for i := 0; i < len(list); i++ {
		for j := 1; j < len(list); j++ {
			if list[j-1] > list[j] {
				//temp := list[j-1]
				//list[j-1] = list[j]
				//list[j] = temp
				swap(list, j-1, j)
			}
		}
	}
	//return list
}

//
//func swap(list []int,m int,n int){
//	temp := list[m]
//	list[m] = list[n]
//	list[n] = temp
//}

func swap(arr []int, i int, j int) {
	arr[i] = arr[i] ^ arr[j]
	arr[j] = arr[j] ^ arr[i]
	arr[i] = arr[i] ^ arr[j]
}
