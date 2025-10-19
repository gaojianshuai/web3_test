package main

import "fmt"

func main() {

	//1.数组的长度时类型的一部分

	var arr [3]int
	var arr1 [4]int
	var strArr [3]string
	fmt.Printf("arr：%T  arr1：%T  strArr：%T", arr, arr1, strArr)

	//2.数组的初始化，第一种方法,长度确定之后不能改变
	var arr2 [3]int
	arr2[0] = 10
	arr2[1] = 20
	arr2[2] = 30
	fmt.Println(arr2) //[10 20 30]

	//3第二种初始化数组
	var arr3 = [3]int{3, 4, 5}
	var string1 = [4]string{"python", "java", "php", "golang"}
	fmt.Println(arr3)
	fmt.Println(string1)

	string2 := [4]string{"python", "java", "php", "c++"}
	fmt.Println(string2)

	//3.数组的另一种初始化  ...   任意动态  推导数组的长度

	var arr4 = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 888}
	fmt.Println(arr4)

	string3 := [4]string{"python", "java", "php", "c++"}
	// string3[4] = "js"
	fmt.Println(string3)
	//改变数组的值
	string3[1] = "js"
	fmt.Println(string3)

	//4.第四种出书画数组
	var arr5 = [...]int{0: 33, 1: 33, 2: 55, 3: 777, 12: 88}

	fmt.Println(len(arr5))
	fmt.Println(arr5)

	//5.数组的循环遍历  for  for  range
	var arr6 = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 888}
	for i := 0; i < len(arr6); i++ {
		fmt.Println(i)
	}

	string4 := [4]string{"python", "java", "php", "c++"}
	for i := 0; i < len(string4); i++ {
		fmt.Println(string4[i])
	}

	//for range 循环
	string5 := [4]string{"python", "java", "php", "c++"}
	for k, v := range string5 {
		fmt.Println(k, v)
	}

	//1.练习1：求出一个数组里面元素的和以及这些元素的平均值，分别用for和for range实现

	var arr7 = [...]int{22, 33, 44, 55, 66, 77}
	var sum = 0
	for i := 0; i < len(arr7); i++ {
		sum = sum + arr7[i]
	}
	fmt.Println(sum)
	fmt.Printf("arr7数组元素的总和时：%v   平均值：%v\n", sum, sum/len(arr7))

	//2.请求出一个数组里面的最大值，并且打印出下标

	var arr8 = [...]int{22, 33, 44, 55, 66, 77}
	var max = arr8[0]
	var index1 = 0
	for i := 0; i < len(arr8); i++ {
		if max < arr8[i] {
			max = arr8[i]
			index1 = i
		}
	}
	fmt.Println(max, index1)

	//3.从[1, 3 , 5, 7 ,8]中找出和为8的俩个元素的小标分别是（0， 3）, (1, 2)
	var arr9 = [...]int{1, 3, 5, 7, 8}
	for i := 0; i < len(arr9); i++ {
		for j := i + 1; j < len(arr9); j++ {
			if arr9[i]+arr9[j] == 8 {
				fmt.Println(i, j)
			}
		}
	}

	//引用类型：改变副本得值，会改变本身的值
	//  值类型：改变变量副本值得时候幕布灰改变变量本身的值

	// 多维数组：以二维数组:  注意多维数组只有第一层可以使用...来让编译器推导数组长度 [...][3]   不支持[3][...]
	// var arr10 = [3]int{1, 3, 4, 5} //一维数组

	//二维数组
	var arr11 = [3][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	//for range获取
	for _, v1 := range arr11 {
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}
	//for获取
	for i := 0; i < len(arr11); i++ {
		for j := 0; j < len(arr11[i]); j++ {
			fmt.Println(arr11[i][j])
		}
	}

}
