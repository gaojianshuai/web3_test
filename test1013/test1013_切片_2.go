package main

import "fmt"

func main() {

	// make()函数创建切片  make([]T, size, cap)
	var slice1 = make([]int, 4, 99)
	fmt.Println(slice1)
	fmt.Println(len(slice1), cap(slice1))
	//修改切片值
	slice1[0] = 11
	slice1[1] = 22
	slice1[2] = 33
	slice1[3] = 44
	fmt.Println(slice1)

	var slice2 []int
	fmt.Println(len(slice2), cap(slice2)) //长度为0   容量为0

	// append  给切片库容：动态添加元素,可以同时添加多个数据
	slice2 = append(slice2, 22, 44, 55, 66)
	fmt.Println(slice2, len(slice2), cap(slice2))

	//apend可以合并切片

	sliceA := []string{"python", "java"}
	sliceB := []string{"golang"}
	sliceA = append(sliceA, sliceB...)
	fmt.Println(sliceA)

	//切片扩容
	var sliceC []int
	for i := 0; i <= 10; i++ {
		sliceC = append(sliceC, i)
		fmt.Println(sliceC, len(sliceC), cap(sliceC))
	}

	//golang中的copy()函数
	sliceD := []int{1, 2, 3, 4, 5, 6}
	sliceE := make([]int, 6, 7)
	copy(sliceE, sliceD)
	fmt.Println(sliceE)
	fmt.Println(sliceD)
	sliceE[0] = 33
	fmt.Println(sliceE)

	//从切片中删除元素
	sliceF := []int{1, 2, 3, 4, 5, 6}

	//删除索引为2的元素:合并切片的时候最后一个要加三个点...
	sliceF = append(sliceF[:2], sliceF[3:]...)
	fmt.Println(sliceF)

	//
	s1 := "你好,golang"
	runeStr := []rune(s1)
	fmt.Println(runeStr) //[20320 22909 44 103 111 108 97 110 103]

	runeStr[0] = '大'
	fmt.Println(string(runeStr)) //大好,golang
}
