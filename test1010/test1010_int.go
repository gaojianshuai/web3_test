package main

import "fmt"

func main() {

	//1.整形
	var num int = 100
	fmt.Printf("num=%v 数据类型是%T", num, num)

	var a1 int32 = 33
	var a2 int64 = 66
	fmt.Printf("a1=%v 类型是%T a2=%v 类型是%T", a1, a1, a2, a2)
	fmt.Println(int64(a1) + a2)

	//高位向低位转换的时候要注意

	var b1 int16 = 133
	fmt.Println(int8(b1)) //转换结果是-123，就有问题了

	var b2 int16 = 120
	fmt.Println(int8(b2)) //这个时候就没问题，一定要注意高位转低位

	//%v：原样输出  %d十进制输出  %b二进制输出  %o 八进制输出  %x16进制输出
	fmt.Printf("b2=%v\n", b2)
	fmt.Printf("b2=%d\n", b2)
	fmt.Printf("b2=%b\n", b2)
	fmt.Printf("b2=%o\n", b2)
	fmt.Printf("b2=%x\n", b2)
}
