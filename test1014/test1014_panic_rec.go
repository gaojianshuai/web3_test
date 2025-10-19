package main

import (
	"errors"
	"fmt"
)

// go语言中没有异常机制，但是可以使用panic和recover模式来处理异常机制
// panic 可以放在任何地方引发，但是recover 只有在defer调用的函数中有效
func mm() {
	fmt.Println("开始测试")
}

func mm1() {

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("err:", err)
		}
	}()
	panic("抛出异常") //遇到panic直接结束执行   recover可以监听  且继续执行
}

func mm2(a int, b int) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("error:", err)
		}
	}()
	fmt.Println(a / b)
}

func readFile(fileName string) error {
	if fileName == "main.go" {
		return nil
	} else {
		return errors.New("读取文件失败")
	}
}

func myFn() {
	defer func() {
		e := recover()
		if e != nil {
			fmt.Println("发送邮件给管理员")
		}
	}()
	err := readFile("xxx.go")
	if err != nil {
		panic(err)
	}
}

func main() {

	mm()
	mm1()

	mm2(23, 4)

	myFn()
	fmt.Println("继续")
}
