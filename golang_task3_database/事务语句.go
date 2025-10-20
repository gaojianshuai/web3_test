package main

/*
题目2：事务语句
假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表
（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
要求 ：
编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，
如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。
*/
import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Account struct {
	ID      uint `gorm:"primaryKey"`
	Balance int
}

type Transaction struct {
	ID            uint `gorm:"primaryKey"`
	FromAccountID uint
	ToAccountID   uint
	Amount        int
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
	// 自动迁移表结构
	db.AutoMigrate(&Account{}, &Transaction{})
	// 初始化账户A余额为500
	db.FirstOrCreate(&Account{ID: 1, Balance: 500})
	db.FirstOrCreate(&Account{ID: 2, Balance: 0})

	// 示例转账
	fromAccountID := uint(1)
	toAccountID := uint(2)
	amount := 100

	if err := TransferMoney(db, fromAccountID, toAccountID, amount); err != nil {
		log.Printf("转账失败: %v", err)
	} else {
		log.Println("转账成功")
	}
}

func TransferMoney(db *gorm.DB, fromID, toID uint, amount int) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	// 检查转出账户余额
	var fromAccount Account
	if err := tx.First(&fromAccount, fromID).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("查询转出账户失败: %w", err)
	}

	if fromAccount.Balance < amount {
		tx.Rollback()
		return fmt.Errorf("账户余额不足")
	}

	// 更新账户余额
	if err := tx.Model(&Account{}).Where("id = ?", fromID).
		Update("balance", gorm.Expr("balance - ?", amount)).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("扣除转出账户失败: %w", err)
	}

	if err := tx.Model(&Account{}).Where("id = ?", toID).
		Update("balance", gorm.Expr("balance + ?", amount)).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("增加转入账户失败: %w", err)
	}

	// 记录交易
	transaction := Transaction{
		FromAccountID: fromID,
		ToAccountID:   toID,
		Amount:        amount,
	}
	if err := tx.Create(&transaction).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("记录交易失败: %w", err)
	}

	return tx.Commit().Error
}
