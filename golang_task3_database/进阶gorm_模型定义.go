package main

/*
题目1：模型定义
假设你要开发一个博客系统，有以下几个实体： User （用户）、 Article （文章）、 Comment （评论）。
要求 ：
使用Gorm定义 User 、 Article 和 Comment 模型，其中 User 与 Article 是一对多关系（一个用户可以发布多篇文章），
Article 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
编写Go代码，使用Gorm创建这些模型对应的数据库表。

*/

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string    `gorm:"unique;not null"`
	Email    string    `gorm:"type:varchar(255);not null"`
	Password string    `gorm:"not null"`
	Articles []Article `gorm:"foreignKey:UserID"`
}

type Article struct {
	gorm.Model
	Title    string    `gorm:"not null"`
	Content  string    `gorm:"type:text"`
	UserID   uint      `gorm:"not null"`
	Comments []Comment `gorm:"foreignKey:ArticleID"`
}

type Comment struct {
	gorm.Model
	Content   string `gorm:"type:text"`
	ArticleID uint   `gorm:"not null"`
	UserID    uint   `gorm:"not null"`
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

	// 自动迁移模型
	err = db.AutoMigrate(&User{}, &Article{}, &Comment{})
	if err != nil {
		log.Fatal("Failed to migrate tables:", err)
	}

	fmt.Println("Database tables created successfully!")

	//创建用户、文章、评论
	var users = []User{
		{Username: "科比", Email: "112345678@QQ.COM", Password: "123456", Articles: []Article{
			{Title: "GORM教程", Content: "这是关于GORM的教程文章", UserID: 1, Comments: []Comment{
				{Content: "很有帮助的文章！", UserID: 1, ArticleID: 1}}}}},
		{Username: "乔丹", Email: "112345678@QQ.COM", Password: "123456", Articles: []Article{
			{Title: "GORM高级用法", Content: "这是关于GORM的高级用法文章", UserID: 2, Comments: []Comment{
				{Content: "学到了很多新知识！", UserID: 2, ArticleID: 2}, {Content: "太强了学到好多知识，给你点赞！", UserID: 2, ArticleID: 2}}}}}}

	result1 := db.Create(users) //通过数据切片来创建
	println(result1.RowsAffected)
	// 查询某个用户发布的所有文章及其对应的评论信息
	var user User
	userID := 1 // 假设要查询的用户 ID 为 1
	err = db.Preload("Articles.Comments").First(&user, userID).Error
	if err != nil {
		fmt.Println("查询失败:", err)
		return
	}

	// 打印结果
	fmt.Printf("用户: %s\n", user.Username)
	for _, article := range user.Articles {
		fmt.Printf("文章: %s\n", article.Title)
		for _, comment := range article.Comments {
			fmt.Printf("评论: %s\n", comment.Content)
		}
	}

	// 查询评论数量最多的文章信息
	var article Article
	err = db.Model(&Article{}).
		Select("articles.*, COUNT(comments.id) as comment_count").
		Joins("left join comments on comments.article_id = articles.id").
		Group("articles.id").
		Order("comment_count DESC").
		Limit(1).
		Find(&article).Error
	if err != nil {
		fmt.Println("查询失败:", err)
		return
	}

	// 打印结果
	fmt.Printf("评论数量最多的文章: %s\n", article.Title)

}
