package main

/*

题目2：实现类型安全映射
假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
要求 ：
定义一个 Book 结构体，包含与 books 表对应的字段。
编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。
*/
import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	ID     uint    `gorm:"primaryKey"`
	Title  string  `gorm:"size:255"`
	Author string  `gorm:"size:255"`
	Price  float64 `gorm:"type:decimal(10,2)"`
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

	db.AutoMigrate(&Book{})
	Books := []*Book{
		{ID: 1, Title: "水浒传", Author: "施耐庵", Price: 12},
		{ID: 2, Title: "三国演义", Author: "罗贯中", Price: 55},
		{ID: 3, Title: "西游记", Author: "吴承恩", Price: 33},
		{ID: 4, Title: "红楼梦", Author: "曹雪芹", Price: 66},
	}
	result1 := db.Create(&Books) //通过数据指针来创建
	println(result1.RowsAffected)

	var books []Book
	result := db.Where("price > ?", 50).Find(&books)
	if result.Error != nil {
		log.Fatal("Query failed:", result.Error)
	}

	fmt.Println("价格大于50元的书籍:")
	for _, book := range books {
		fmt.Printf("ID: %d, 书名: %s, 作者: %s, 价格: %.2f\n",
			book.ID, book.Title, book.Author, book.Price)
	}
	fmt.Printf("books得类型是%T\n", books)
}
