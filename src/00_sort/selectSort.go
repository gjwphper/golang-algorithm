package main

import "fmt"

//
//选择排序就好比第一个数字站在擂台上，大吼一声：“还有谁比我小？”。剩余数字来挨个打擂，如果出现比第一个数字小的数，则新的擂主产生。每轮打擂结束都会找出一个最小的数，将其交换至首位。经过 n-1 轮打擂，所有的数字就按照从小到大排序完成了。

//
//现在让我们思考一下，冒泡排序和选择排序有什么异同？
//
//相同点：
//
//都是两层循环，时间复杂度都为 O(n^2)O(n
//2
//);
//都只使用有限个变量，空间复杂度 O(1)O(1)。
//不同点：
//
//冒泡排序在比较过程中就不断交换；而选择排序增加了一个变量保存最小值 / 最大值的下标，遍历完成后才交换，减少了交换次数。
//事实上，冒泡排序和选择排序还有一个非常重要的不同点，那就是：
//
//冒泡排序法是稳定的，选择排序法是不稳定的。

//
//作者：力扣 (LeetCode)
//链接：https://leetcode-cn.com/leetbook/read/sort-algorithms/ev1l5g/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
func main() {
	list := []int{6, 3, 5, 2, 4, 1}
	selectSort(list)
	fmt.Println(list)
}

/**
1、2层循环，里面的作比较
2、list构建
3、交换
4、j:=i+1
5、记录最小的下标，当前第i个，为选出来的最小的
*/

func selectSort(list []int) {

	for i := 0; i < len(list)-1; i++ {
		minIndex := i //先初始化假设为最小的
		for j := i + 1; j < len(list); j++ {
			if list[minIndex] > list[j] {
				minIndex = j
			}
		}

		temp := list[i]
		list[i] = list[minIndex]
		list[minIndex] = temp
		fmt.Println(list)

	}

}
