package main

import (
	"database/sql"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserInfo struct {
	ID           uint           // Standard field for the primary key
	Name         string         // A regular string field
	Email        *string        // A pointer to a string, allowing for null values
	Age          uint8          // An unsigned 8-bit integer
	Birthday     *time.Time     // A pointer to time.Time, can be null
	MemberNumber sql.NullString // Uses sql.NullString to handle nullable strings
	ActivatedAt  sql.NullTime   // Uses sql.NullTime for nullable time fields
	CreatedAt    time.Time      // Automatically managed by GORM for creation time
	UpdatedAt    time.Time      // Automatically managed by GORM for update time
	ignored      string         // fields that aren't exported are ignored
}

type Author struct {
	Name  string
	Email string
}

type Blog struct {
	Author
	ID      int
	Upvotes int32
}

// equals
type Blog1 struct {
	ID      int64
	Name    string
	Email   string
	Upvotes int32
}

func Run(db *gorm.DB) {
	db.AutoMigrate(&UserInfo{})
	// UserInfo := UserInfo{Name: "java123", Age: 77}	//创建一条
	//创建多条

	UserInfo := []*UserInfo{
		{Name: "科比", Age: 40},
		{Name: "乔丹", Age: 40},
	}
	// UserInfo.MemberNumber.Valid = true
	// db.Table("test1014")
	result := db.Create(UserInfo)

	fmt.Println(result.RowsAffected)
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
	//延时关闭数据库连接
	// defer db.Close()
	Run(db)

	var u UserInfo
	db.First(&u)
	fmt.Println(u)

	// db.Model(&u).Update(username, "python666")

}
