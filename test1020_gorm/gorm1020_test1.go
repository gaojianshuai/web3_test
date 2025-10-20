package test1020_gorm

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type User struct {
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
	CreateCard   CreateCard     // Added field for CreateCard
}

// gorm.Model 的定义
type Model struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
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
// type Blog struct {
// 	ID      int64
// 	Name    string
// 	Email   string
// 	Upvotes int32
// }

// type Author struct {
// 	Name  string
// 	Email string
// }

type Blog2 struct {
	ID int
	//嵌入字段
	// Author  Author `gorm:"embedded"`
	//指定字段前缀
	Author Author `gorm:"embedded;embeddedPrefix:author_"`
	//修改列明
	Upvotes int32 `gorm:"column:votes;"`
	// Upvotes int32
}

// // 等效于
//
//	type Blog struct {
//		ID      int64
//		Name    string
//		Email   string
//		Upvotes int32
//	}
// type Blog struct {
// 	ID      int
// 	Author  Author `gorm:"embedded;embeddedPrefix:author_"`
// 	Upvotes int32
// }

// 等效于
// type Blog struct {
// 	ID          int64
// 	AuthorName  string
// 	AuthorEmail string
// 	Upvotes     int32
// }

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Name = u.Name + "first_create_name"
	return
}

// func (u *User) AfterCreate(tx *gorm.DB) (err error) {
// 	u.Name = u.Name + "last_create_name"
// 	return
// }

type CreateCard struct {
	Model
	Number string
	UserId int64
}

type User6 struct {
	Model
	Name       string
	CreateCard CreateCard
}

func Run(db *gorm.DB) {
	// //自动创建表
	// db.AutoMigrate(&User{})
	// db.AutoMigrate(&Blog2{})
	// //差一条空数据
	// user := &User{}
	// // 不生成指针
	// user.MemberNumber.Valid = true
	// //创建一个表数据
	// db.Create(user)
	// db.Create(&User{
	// 	Name:       "jinzhu",
	// 	CreateCard: CreateCard{Number: "5555555555555555"},
	// })
	// db.Omit("CreateCard").Create(&user)
	// //创建一条数据
	// user1 := User{Name: "python", Age: 33}
	// result := db.Create(&user1) //通过数据指针来创建
	// //返回找到的记录数
	// println(result.RowsAffected)

	// users := []*User{
	// 	{Name: "golang", Age: 44},
	// 	{Name: "java", Age: 55},
	// 	{Name: "js", Age: 66},
	// }
	// result1 := db.Create(&users) //通过数据指针来创建
	// println(result1.RowsAffected)
	// //获取第一条记录（主键升序）
	// db.Debug().First(&user1)
	// //获取最后一条数据（主键降序）
	// db.Debug().Last(&user1)
	// //获取一条记录没有指定排序字段
	// result2 := db.Take(&user1)
	// println(result2.RowsAffected)
	// users3 := []*User{
	// 	{Name: "科比", Age: 44},
	// 	{Name: "ja詹姆斯va", Age: 55},
	// 	{Name: "乔丹", Age: 66},
	// }
	// //指定字段创建记录
	// db.Select("Name", "Age").Create(&users3)

	// //【批量插入

	// var users4 = []User{
	// 	{Name: "张三", Age: 30},
	// 	{Name: "李四", Age: 25},
	// 	{Name: "王五", Age: 28},
	// }
	// for _, user := range users4 {
	// 	db.Create(&user)
	// }
	// //通过db.CreateInBatches 批量插入
	// var users5 = make([]User, 1000)
	// for i := 0; i < 100; i++ {
	// 	users5[i] = User{Name: "gaojs" + fmt.Sprintf("%d", i+1)}
	// }
	// db.CreateInBatches(users5, 10) //每批次10条

	// //跳过钩子函数
	// db.Session(&gorm.Session{SkipHooks: true}).Create(&User{Name: "跳过钩子函数"})

	// //排序
	// result3 := db.Order("age").Find(&users)
	// fmt.Println(result3)

	// //更新
	// db.Model(&User{}).Where("active = pythonfirst_create_name", true).Update("name", "metaweb3")
	//删除
	db.Delete("pythonfirst_create_name")
	db.Where("name = ?", "pythonfirst_create_name").Delete(&db.Name)
}
