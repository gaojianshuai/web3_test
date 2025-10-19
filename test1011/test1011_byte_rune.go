package main

import "fmt"

func main() {

	//1.浮点型转换：优先推荐使用64双精度
	var float1 float64 = 12.0
	float2 := 10
	float3 := 20
	float4 := float64(float2)
	fmt.Println(float1)
	fmt.Println(float3)
	fmt.Println(float1 == float4)

	//2.rune类型

	var r1 rune = 'a'
	var r2 rune = 'b'
	fmt.Println(r1, r2)

	//3.rune也可以直接转换成切片
	var s string = "hello， 你好，世界"
	var r3 = []rune(s)
	fmt.Println(r3)

	//一个汉字占用三个字节    utf-8
}
