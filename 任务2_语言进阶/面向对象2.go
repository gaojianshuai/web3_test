package main

import "fmt"

// 面向对象编程
type Person struct {
	Name string
	Age  int
}

// Employee 结构体, 组合Person
type Employee struct {
	Person     // Ensure Employee implements the Person interface by defining the required methods
	EmployeeID string
}

// printInfo 打印员工信息
func (e Employee) printInfo() {
	fmt.Println("员工信息如下：\n")
	fmt.Printf("姓名: %s\n", e.Name)
	fmt.Printf("年龄: %d\n", e.Age)
	fmt.Printf("员工ID: %s\n", e.EmployeeID)
}

//或者直接创建员工函数

func NewEmployee(Name string, Age int, EmployeeID string) Employee {
	return Employee{
		Person:     Person{Name: Name, Age: Age},
		EmployeeID: EmployeeID,
	}
}

func main() {
	// 直接创建实例的方式
	emp1 := Employee{
		Person: Person{
			Name: "王五",
			Age:  25,
		},
		EmployeeID: "E1003",
	}
	emp1.printInfo()

	// 使用构造函数创建实例
	emp2 := NewEmployee("赵六", 28, "E1004")
	emp2.printInfo()

}
