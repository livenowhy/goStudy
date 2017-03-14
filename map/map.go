package  main

import "fmt"


func test1()  {
	// 映射的创建
	// make(map[KeyType] ValueType, initialCapacity)
	// make(map[KeyType] ValueType)
	// map [KeyType] ValueType {}
	// map [KeyType] ValueType {key1: value1, key2:value2, ...}
	// 用4种方式分别创建数组
	// 其中第一种和第二种的区别在于，有没有指定初始容量，不过使用的时候则无需在意这些，因为map的本质决定了，一旦容量不够，它会自动扩容
	map1 := make(map[string]string, 5)
	map2 := make(map[string]string)
	map3 := map[string]string{}
	map4 := map[string]string{"a": "1", "b": "2"}
	fmt.Println(map1, map2, map3, map4)

		// map指定key取对应的value时,可以指定返回两个值,第一个是对应的value,第二个是一个bool,表示是否有值s
	if v, ok := map4["a"]; ok{
		fmt.Println(v)
		fmt.Println("Key is Found")
	} else {
		fmt.Println("Key Not Found")
	}
}

func test2()  {
	// 映射的填充和遍历
	map1 := make(map[string]string)
	map1["a"] = "1"
	map1["b"] = "2"
	for key, value := range map1 {
		fmt.Printf("%s->%s \n", key, value)
	}
	// 数组的填充使用 map[key] = value 的方式,遍历映射的时候,每一项都返回2个值,键和值
}

func test3()  {
	// 映射的查找、修改和删除
	map4 := map[string]string{"a": "1", "b": "2"}

	// 修改映射和添加映射的操作没什么区别,若指定的键不存在则创建;否则,修改之
	map4["a"] = "4"

	// 删除则是使用go的内置函数delete
	delete(map4, "a")
}

func main() {
	test1()
	test2()
	test3()
}

