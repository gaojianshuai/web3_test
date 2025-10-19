package main

import (
	"fmt"
	"sort"
)

func main() {

	//选择排序:第一个和后面的作对比，然后交换位置

	var numSlice = []int{9, 8, 5, 6, 7, 11}
	for i := 0; i < len(numSlice); i++ {
		for j := i + 1; j < len(numSlice); j++ {
			if numSlice[i] < numSlice[j] {
				temp := numSlice[i]
				numSlice[i] = numSlice[j]
				numSlice[j] = temp
			}
		}

	}
	fmt.Println(numSlice)

	//冒泡排序
	var numSlice1 = []int{9, 8, 5, 6, 7, 11}
	for i := 0; i < len(numSlice1); i++ {
		for j := 0; j < len(numSlice1)-1-i; j++ {
			if numSlice1[j] > numSlice1[j+1] {
				temp := numSlice1[j]
				numSlice1[j] = numSlice1[j+1]
				numSlice1[j+1] = temp
			}
		}
	}
	fmt.Println(numSlice1)

	//sort包：升序  降序
	//升序

	intList := []int{1, 3, 2, 4, 5, 6, 66, 5, 10}
	floatList := []float64{1.1, 3.3, 2.2, 4.3, 55.4, 0.4}
	stringList := []string{"a", "c", "b", "t", "z", "h"}

	sort.Ints(intList)
	fmt.Println(intList)
	sort.Float64s(floatList)
	fmt.Println(floatList)
	sort.Strings(stringList)
	fmt.Println(stringList)

	//降序
	sort.Sort(sort.Reverse(sort.IntSlice(intList)))

	sort.Sort(sort.Reverse(sort.Float64Slice(floatList)))

	sort.Sort(sort.Reverse(sort.StringSlice(stringList)))
	fmt.Println(intList, floatList, stringList)

}
