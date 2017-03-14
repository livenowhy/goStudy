package main

import "fmt"


// 定义函数
type myType struct {
	a int
	b int
}

// 该 method 属于 myType 类型对象中的方法
func (m myType) geCount() int  {
	// m.a 即为 myType 类型对象中的属性
	return 2 + m.a + m.b
}

func main() {
	var m1 myType
	m1.a = 10
	m1.b = 20

	fmt.Println(m1.geCount())
}


