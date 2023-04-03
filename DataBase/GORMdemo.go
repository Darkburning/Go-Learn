package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type USER struct { // 结构体对应数据库中的表，其实例对应数据库中的记录
	Id       int // 注意结构体中变量名大写才可
	Username string
	Password int
}

func main() {

	dsn := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	user := USER{
		Id:       003,
		Username: "Jason",
		Password: 123456,
	}
	//// 插入新纪录
	//result := db.Create(&user)
	//fmt.Printf("RowsAffected: %v\n", result.RowsAffected)
	// 删除记录
	db.Delete(&user) // 默认按主键值删除
	// 查询：
	var recUser []USER
	//result = db.First(&recUser) // 获取第一条记录（主键升序）
	//result = db.Last(&recUser) // 获取最后一条记录（主键降序）
	result := db.Find(&recUser) // 返回所有记录,可以用切片接收
	fmt.Printf("%v\n", recUser)
	fmt.Printf("RowAffected: %v\n", result.RowsAffected)
}
