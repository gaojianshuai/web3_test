/*
目1：基本CRUD操作
假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
要求 ：
编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
*/

// -- 插入新记录
// INSERT INTO students (name, age, grade) VALUES ('张三', 20, '三年级');

// -- 查询年龄大于18岁的学生
// SELECT * FROM students WHERE age > 18;

// -- 更新张三的年级
// UPDATE students SET grade = '四年级' WHERE name = '张三';

// -- 删除年龄小于15岁的学生

// DELETE FROM students WHERE age < 15;

// 使用golang去执行上述SQL语句
package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	ID    int
	Name  string
	Age   int
	Grade string
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
		fmt.Println(db)
		fmt.Println("连接成功")

	}
	db.AutoMigrate(&Student{})
	students := []*Student{
		{ID: 1, Name: "张三", Age: 20, Grade: "三年级"},
		{ID: 2, Name: "小王", Age: 34, Grade: "八年级"},
		{ID: 3, Name: "小李", Age: 12, Grade: "九年级"},
		{ID: 4, Name: "小可", Age: 21, Grade: "六年级"},
	}
	result1 := db.Debug().Create(&students) //通过数据指针来创建
	println(result1.RowsAffected)

	//查询数据
	result2 := db.Debug().Where("age > ?", 18).Find(&students)
	println(result2.RowsAffected)
	//更新数据
	result3 := db.Debug().Model(&Student{}).Where("name = ?", "张三").Update("grade", "四年级")
	println(result3.RowsAffected)
	//删除数据
	result4 := db.Debug().Where("age < ?", 15).Delete(&Student{})
	println(result4.RowsAffected)
}
