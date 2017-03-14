package main

import "fmt"
import "encoding/json"



/** 假设我们有如下一个类(结构体)student 及其一个实例对象st */
type Student struct {
	Name     string
	Age      int
	Guake    bool
	Classes  []string
	Prince   float32
}



/** 现在需要把这个类的一个对象转换为JSON格式 */

func (s * Student)ShowStu() {
	fmt.Println("show Student:")
	fmt.Println(s.Name)
	fmt.Println(s.Age)
	fmt.Println(s.Guake)
	fmt.Println(s.Prince)
	fmt.Println(s.Classes)
}

func main() {
	st := &Student{
	"Xiao Ming",
	16,
	true,
	[]string{"Math", "English", "Chinese"},
	9.99,
	}

	fmt.Println("before JSON encoding .... ")
	st.ShowStu()

	b, err := json.Marshal(st)
	if err != nil {
		fmt.Println("encoding faild")
	} else {
		fmt.Println("encoded data")
		fmt.Println(b)
		fmt.Println(string(b))
	}

	sts := &Student{}
	sts.ShowStu()
	err = json.Unmarshal([]byte(b), &sts)
	if err != nil{
		fmt.Println("Unmarshal faild")
	} else {
		fmt.Println("Unmarshal ok")
		sts.ShowStu()
	}
}