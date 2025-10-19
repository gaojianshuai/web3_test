package main

import (
	"fmt"
	"strings"
)

func main() {

	//1. 定义字符串
	var num string = "web3"
	fmt.Printf("num=%v 数据类型是%T", num, num)

	//2. 定义转义字符
	str1 := "my name is \ngaojs"
	fmt.Printf("str1=%v 数据类型是%T\n", str1, str1)

	//转义府
	str2 := "C:\\Go\\bin" //或者c:/Go/bin
	fmt.Printf("str1=%v 数据类型是%T", str2, str2)

	//3.多行字符串
	str3 := `
		apple
		banana
		orange
		test1
	`
	fmt.Println(str3)

	// 4.len求字符串长度

	var str4 = "钱包"
	fmt.Println(len(str4))

	//5.拼接字符串:加号拼接     sprintf拼接

	str5 := "你好"
	str6 := "小高"
	str7 := str5 + str6
	fmt.Println(str7)

	str8 := fmt.Sprintf("%v    %v", str5, str6)
	fmt.Println(str8)

	str9 := "教程更新于2025年，所有内容支持最新版本。" +
		"更多Golang+微服务+K8s教程访问"
	fmt.Println(str9)

	//6.分割字符串
	var str10 = "my name is gaojiaoshuai-test web3.0"
	arr := strings.Split(str10, " ")
	fmt.Println(arr)

	//7.strings.Join   表示吧切片连接成字符串
	str11 := strings.Join(arr, "**")
	fmt.Println(str11)

	arr2 := []string{"java", "python", "go"}
	fmt.Println(arr2)
	arr3 := strings.Join(arr2, "---")
	fmt.Println(arr3)

	//8.strings.Contains 判断是否包含
	str12 := "this is my job"
	str13 := "job"
	flag := strings.Contains(str12, str13)
	fmt.Println(flag)

	// 9.strings.HasPrefix    strings.HasSuffix   前缀或后缀
	str14 := "this is my job"
	str15 := "job"
	flag2 := strings.HasPrefix(str14, str15)
	flag3 := strings.HasSuffix(str14, str15)
	fmt.Println(flag2)
	fmt.Println(flag3)

	//10.strings.Index()   strings.LastIndex()   子串出现位置  查找不到返回-1
	str16 := "this is my job test"
	str17 := "is"
	// num := strings.Index(str16,str17)
	fmt.Println(strings.Index(str16, str17))
	fmt.Println(strings.LastIndex(str16, str17))
	// 	num1 := strings.LastIndex(str16,str17)
	// 	print(num)
	// 	print(num1)

}
