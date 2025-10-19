package main

import (
	"fmt"
	"sort"
)

func sortIntAsc(slices []int) []int {
	//升序排序函数
	for i := 0; i < len(slices); i++ {
		for j := 0; j < len(slices); j++ {
			if slices[i] < slices[j] {
				temp := slices[i]
				slices[i] = slices[j]
				slices[j] = temp
			}
		}
	}
	return slices
}

func sortIntDesc(slices []int) []int {
	//降序排序函数
	for i := 0; i < len(slices); i++ {
		for j := 0; j < len(slices); j++ {
			if slices[i] > slices[j] {
				temp := slices[i]
				slices[i] = slices[j]
				slices[j] = temp
			}
		}
	}
	return slices
}

func mapSort(map1 map[string]string) string {
	var sliceKey []string
	//把map对象中的key放在切片里面
	for k, _ := range map1 {
		sliceKey = append(sliceKey, k)
	}
	//2.对key进行升序排序
	sort.Strings(sliceKey)

	var str string
	for _, v := range sliceKey {
		str += fmt.Sprintf("%v=>%v", v, map1[v])
	}
	return str
}

var aa = "quanjubianliang"

func runtest() {
	fmt.Println("runtest")
}
func main() {

	var slice1 = []int{1, 33, 4, 55, 66, 0, 9}
	slice2 := sortIntAsc(slice1)
	fmt.Println(slice2)

	var slice3 = []int{1, 33, 4, 55, 66, 0, 9}
	slice4 := sortIntDesc(slice3)
	fmt.Println(slice4)

	//局部变量  全局变量

	runtest()
	fmt.Println(aa)
}
