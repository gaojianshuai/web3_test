package main

/*

题目1：使用SQL扩展库进行查询
假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
要求 ：
编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
*/
import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Employee struct {
	ID         int
	Name       string
	Department string
	Salary     int
}

func main() {
	//配置MySQL连接参数
	username := "root"      //账号
	password := "gjs199074" //密码
	host := "127.0.0.1"     //数据库地址，可以是Ip或者域名
	port := 3306            //数据库端口
	Dbname := "web3_test"   //数据库名
	timeout := "10s"        //连接超时，10秒

	//拼接dsn参数，这里使用Sprintf动态拼接
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	fmt.Println(db)
	fmt.Println("连接成功")

	db.AutoMigrate(&Employee{})
	Employees := []*Employee{
		{ID: 1, Name: "小张", Department: "技术部", Salary: 7000},
		{ID: 2, Name: "小王", Department: "市场部", Salary: 6000},
		{ID: 3, Name: "小李", Department: "技术部", Salary: 8000},
		{ID: 4, Name: "小可", Department: "技术部", Salary: 9000},
	}
	result1 := db.Create(&Employees) //通过数据指针来创建
	println(result1.RowsAffected)
	// 查询技术部员工
	// 查询技术部员工
	var techEmployees []Employee
	result := db.Where("Department = ?", "技术部").Find(&techEmployees)
	if result.Error != nil {
		fmt.Println("查询失败:", result.Error)
		return
	}

	// 打印技术部员工信息
	fmt.Println("技术部员工信息:")
	for _, emp := range techEmployees {
		fmt.Printf("ID: %d, Name: %s, Department: %s, Salary: %d\n", emp.ID, emp.Name, emp.Department, emp.Salary)
	}

	// 并将结果映射到一个自定义的 Employee 结构体切片中
	// 查询工资最高的员工
	var highestPaidEmployee Employee
	result2 := db.Order("Salary DESC").First(&highestPaidEmployee)
	if result2.Error != nil {
		fmt.Println("查询失败:", result2.Error)
		return
	}

	// 打印工资最高的员工信息
	fmt.Println("工资最高的员工信息:")
	fmt.Printf("ID: %d, Name: %s, Department: %s, Salary: %d\n", highestPaidEmployee.ID, highestPaidEmployee.Name, highestPaidEmployee.Department, highestPaidEmployee.Salary)
}
