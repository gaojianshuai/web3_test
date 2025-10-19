package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	//map是一种无需的基于key-value得数据结构，map是引用类型，必须初始化才能用
	//1.make创建map类型数据
	var userinfo = make(map[string]string)

	userinfo["username"] = "张三"
	userinfo["age"] = "33"
	userinfo["sex"] = "男"
	fmt.Println(userinfo)

	//2.第二种声明map：map也支持在声明的时候补充元素
	var userinfo1 = map[string]string{
		"username": "james",
		"age":      "30",
		"sex":      "男",
	}
	fmt.Println(userinfo1)

	//3.缺省方式

	userinfo2 := map[string]string{
		"username": "james",
		"age":      "30",
		"sex":      "男",
	}
	fmt.Println(userinfo2)

	//4.遍历
	userinfo3 := map[string]string{
		"username": "james",
		"age":      "30",
		"sex":      "男",
	}
	for k, v := range userinfo3 {
		fmt.Println(k, v)

	}

	//map类型数据的crud
	var userinfo4 = make(map[string]string)
	userinfo4["username"] = "张大大"
	userinfo4["age"] = "44"
	fmt.Println(userinfo4)

	userinfo5 := map[string]string{
		"username": "james",
		"age":      "30",
		"sex":      "男",
	}
	fmt.Println(userinfo5)

	//在切片里面放一系列用户信息，这时候我们就可以定义一个元素为map类型的切片
	//var userinfo6 = []string{"gaojs", "test"}
	var userinfo6 = make([]map[string]string, 2, 2)
	fmt.Println(userinfo6) //map得默认值就是nil
	fmt.Println(userinfo6[0] == nil)

	if userinfo6[0] == nil {
		userinfo6[0] = make(map[string]string, 0)
		userinfo6[0]["username"] = "zhangsan"
		userinfo6[0]["age"] = "30"
		userinfo6[0]["height"] = "180"
	}
	if userinfo6[1] == nil {
		userinfo6[1] = make(map[string]string, 0)
		userinfo6[1]["username"] = "李四"
		userinfo6[1]["age"] = "44"
		userinfo6[1]["height"] = "190"
	}
	fmt.Println(userinfo6)

	for _, v := range userinfo6 {
		fmt.Println(v)
		for key, value := range v {
			fmt.Println(key, value)
		}
	}

	//map类型里面加入切片类型
	var userinfo7 = make(map[string][]string)
	userinfo7["hobby"] = []string{
		"吃饭",
		"学习",
		"篮球",
		"游戏",
	}
	userinfo7["work"] = []string{
		"python",
		"java",
		"python",
		"golang",
	}
	for _, v := range userinfo7 {
		// fmt.Println(v)
		for _, val := range v {
			fmt.Println(val)
		}
	}

	// fmt.Println(userinfo7)

	// 值类型：改变变量副本值的时候，不会改变变量本身的值  比如（基本数据类型，。数组）
	// 引用类型：改变变量副本值得时候，会改变变量本身的值(切片、，map)
	var userinfo8 = make(map[string]string)
	userinfo8["username"] = "李思思"
	userinfo8["age"] = "33"
	userinfo9 := userinfo8

	userinfo9["username"] = "董卿"
	fmt.Println(userinfo9)
	fmt.Println(userinfo8)

	//map中的排序
	map1 := make(map[int]int, 10)
	map1[2] = 22
	map1[4] = 11
	map1[1] = 666
	map1[88] = 120
	map1[24] = 9
	for k, v := range map1 {
		fmt.Println(k, v)
	}
	//结果如下
	// 1 666
	// 88 120
	// 2 22
	// 4 11
	// 24 9
	//按照key升序输出k   v， 首先把key放在切片里
	var keySlice []int
	for key, _ := range map1 {
		keySlice = append(keySlice, key)

	}
	fmt.Println(keySlice)
	// sort让key升序
	sort.Ints(keySlice)
	fmt.Println(keySlice)
	//循环遍历key  输出值
	for _, v := range keySlice {
		fmt.Println(v, map1[v])
	}

	//写一个程序，统计一个字符串每个单词出现的次数，比如“how do you do”
	var str = "how do you do"
	var strSlice = strings.Split(str, " ")
	fmt.Println(strSlice)

	var strMap = make(map[string]int)
	for _, v := range strSlice {
		strMap[v]++
	}
	fmt.Println(strMap)
}
