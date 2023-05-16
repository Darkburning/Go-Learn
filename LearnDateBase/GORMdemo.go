package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type USER struct { // 结构体对应数据库中的表，其实例对应数据库中的记录
	ID       int `gorm:"primaryKey" gorm:"id"` // 注意结构体中变量名大写才可
	Username string
	Password string
}

func main() {

	dsn := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	//user := []USER{
	//	{ID: 11, Username: "Kate", Password: "abc123"},
	//	{ID: 12, Username: "Leo", Password: "def456"},
	//	{ID: 13, Username: "Maggie", Password: "xyz789"},
	//	{ID: 14, Username: "Nick", Password: "pqr432"},
	//	{ID: 15, Username: "Olivia", Password: "jkl098"},
	//	{ID: 16, Username: "Peter", Password: "mno765"},
	//	{ID: 17, Username: "Queenie", Password: "rst321"},
	//	{ID: 18, Username: "Rachel", Password: "uvw567"},
	//	{ID: 19, Username: "Sam", Password: "ghi234"},
	//	{ID: 20, Username: "Tom", Password: "lmn890"},
	//}

	//// 插入新纪录
	//result := db.Create(&user) // 传入切片即可实现批量插入
	//result := db.CreateInBatches(&user, 5) // 或者通过制定批量插入的数字来控制批量插入
	//fmt.Printf("RowsAffected: %v\n", result.RowsAffected)
	// 删除记录
	//result := db.Delete(&user, []int{1, 2}) // 默认按主键值删除
	// 查询：
	var recUser []USER
	//result = db.First(&recUser) // 获取第一条记录（主键升序）
	//result = db.Last(&recUser) // 获取最后一条记录（主键降序）
	//result := db.Where("username <> ?", "Alice").Find(&recUser) // where方法构造string条件后续参数填充占位符
	//result := db.Where(&struct{ ID int }{10}).Find(&recUser) // 可用匿名结构体填充where方法返回满足条件的所有
	//result := db.Not([]int{1, 2, 3}).Find(&recUser) // Not方法
	//result := db.Or([]int{1, 2, 3}).Find(&recUser) // Or 方法
	//result := db.Select([]string{"ID", "Username"}).Or([]int{1, 2, 3}).Find(&recUser) // Select方法+Or 方法
	result := db.Select([]string{"ID", "Username"}).Or([]int{1, 2, 3}).Order("Username desc").Find(&recUser) // Order方法

	fmt.Printf("%v\n", recUser)
	fmt.Printf("RowAffected: %v\n", result.RowsAffected)
}
