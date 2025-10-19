package main

import "fmt"

func main() {

	//指针
	var a int = 10

	b := &a
	fmt.Println(a, b)
	c := &b
	fmt.Println(c, &c)
	fmt.Println(*c, **c, *b)
	**c = 30
	fmt.Println(a)

}
