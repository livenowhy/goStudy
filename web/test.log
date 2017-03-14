package main

import "fmt"

func arr_init() {
	// 长度为5的数组, 其元素值依次为：1, 2, 3, 0, 0
	// 在初始化时没有指定初值的元素将会赋值为其元素类型int的默认值0,string的默认值是""
	a1 := [5]int{1, 2, 3}
	fmt.Println(a1)

	// 长度为5的数组, 其长度是根据初始化时指定的元素个数决定的
	a2 := [...] int {1, 2, 3, 4, 5}
	fmt.Println(a2)

	a3 := [5] int {2:1, 4:2, 3:4}
	fmt.Println(a3)

	// 通过for遍历数组元素
	for index, value := range a2 {
		fmt.Println(index, value)
	}

	for index := 0; index < len(a2); index++ {
		fmt.Println(a2[index])
	}
	// 输出
	//[1 2 3 0 0]
	//[1 2 3 4 5]
	//[0 0 1 4 2]
	//0 1
	//1 2
	//2 3
	//3 4
	//4 5
	//1
	//2
	//3
	//4
	//5
}

func slice_init()  {
	// 直接初始化切片, []表示是切片类型, {1,2,3}初始化值依次是1,2,3.其cap=len=3
	s1 := [] int {1, 2, 3, 4, 5}
	fmt.Println(s1)

	arr1 := [...]int{1, 2, 3, 4}
	s2 := arr1   // 初始化切片s2, 是数组 arr1 的引用
	fmt.Println(s2)

	// 通过内置函数make()初始化切片s3, []int 标识为其元素类型为int的切片
	// s :=make([]int,len,cap)
	s3 := make([]int, 3, 6)
	s3[2] = 4

	// 通过下标访问元素时下标不能超过len大小, 如同数组的下标不能超出len范围一样


	// 函数append是GO专门为切片增加元素而提供的一个内建函数
	// 切片可以通过内置函数 append(slice []Type,elems ...Type)追加元素,
	// elems可以是一排type类型的数据,也可以是slice,
	// 因为追加的一个一个的元素, 因此如果将一个slice追加到另一个slice中需要带上"...",
	// 这样才能表示是将slice中的元素依次追加到另一个slice中
	s3 = append(s3, 1,2,3)

	fmt.Println(s3)

	s4 := [] int {1, 3, 3, 4, 5, 6, 7}

	// range表达式有两个返回值, 第一个是索引, 第二个是元素的值：
	for index, value := range s4 {
		fmt.Println(index, value)
	}
}
func main() {
	//arr_init()
	slice_init()
}