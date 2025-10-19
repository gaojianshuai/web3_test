package main //把test.go归属到main
import "fmt"   //引入一个包fmt


func main()  {
	// println会换行
	fmt.Println("hello")
	fmt.Println("我叫web3.0")
	// printf不会换行
	fmt.Printf("你好")
	fmt.Printf("2222")

	fmt.Print("rrrr")
	fmt.Print("tttt")
	// print输出没有空格
	fmt.Print("a", "b", "c")
	// println有空格
	fmt.Println("c", "d", "e")

	// printf和println的区别：变量的时候有区别
	var a = "ddd"	// 变量定义了必须要使用
	fmt.Println(a)
	// 格式化输出
	fmt.Printf("%+v\n", a)


	// 定义变量:第一种：var 变量名称 变量类型（可以省略） 表达式 var name string = "gaojs"    
	// 变量定义第二种：推导模式   变量名 := 表达式

	var b int = 10
	var c int = 11
	var d int = 12
	fmt.Println("b=", b, "c=", c, "d=", d)

	fmt.Printf("b= %v c=%v d=%v\n", b, c, d)

	// 类型推到方式定义变量
	e := 20
	f := 30
	g := 40
	h := 50
	fmt.Printf("e=%v f=%v g=%v h=%v\n", e, f, g, h)

	// 使用printf打印一个变量类型
	fmt.Printf("e=%v e得类型是%T", e, e)
	// 用的最多的就是println和printf


}

