package main

import "fmt"

func main() {

	//1.数组
	var arr3 = [3]int{3, 4, 5}
	var string1 = [4]string{"python", "java", "php", "golang"}
	fmt.Println(arr3)
	fmt.Println(string1)

	//2.切片
	//第一种：不推荐
	var arr1 []int
	fmt.Printf("%v    %T   长度%v\n", arr1, arr1, len(arr1))
	//切面默认值nil
	fmt.Println(arr1 == nil)

	// 第二种
	var arr2 = []int{3, 4, 5}
	fmt.Printf("%v    %T   长度%v\n", arr2, arr2, len(arr2))

	//第三种
	var arr4 = []int{3: 4, 4: 5, 5: 6}
	fmt.Printf("%v    %T   长度%v\n", arr4, arr4, len(arr4))

	//切片的循环遍历
	var strSlice = []string{"python", "java", "php", "golang"}
	for i := 0; i < len(strSlice); i++ {
		fmt.Println(strSlice[i])
	}

	for k, v := range strSlice {
		fmt.Println(k, v)
	}

	//基于数组定义切片

	a := [5]int{1, 2, 3, 4, 5}
	b := a[:] //获取数组里面的所有值
	fmt.Println(b, b)
	//左包右不包
	c := a[1:3]
	fmt.Println(c)

	d := a[2:]
	fmt.Println(d)

	e := a[:3] //表示获取第三个下标前的数据
	fmt.Println(e)

	//基于切片的切片
	var slice1 = []string{"python", "java", "php", "golang"}
	slice2 := slice1[1:]
	fmt.Println(slice2)

	//切片的长度和容量  len()获取长度   cap()获取容量：切片的容量是从他的第一个元素开始数，到其底层数组元素末尾的个数
	tt := []int{1, 2, 3, 4, 5}
	fmt.Println(len(tt), cap(tt)) //长度5  容量5

	aa := tt[1:]
	fmt.Println(len(aa), cap(aa)) //长度4  容量4

	bb := tt[1:3]
	fmt.Println(len(bb), cap(bb)) //长度2，容量4

	cc := tt[:3]
	fmt.Println(len(cc), cap(cc)) //长度2，容量5
}
