package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 测试连接是否成功
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Successfully connected to database!")

	// 往表中插入数据
	//stmt, err := db.Prepare("INSERT INTO test1 (name, number) VALUES (?, ?), (?, ?)")
	//if err != nil {
	//	panic(err.Error())
	//}
	//defer stmt.Close()
	//
	//// 执行 SQL 语句，将值绑定到占位符中
	//result, err := stmt.Exec("Kite", "123", "John", "456")
	//if err != nil {
	//	panic(err.Error())
	//}

	// 获取插入操作的结果
	//lastInsertId, err := result.LastInsertId()
	//if err != nil {
	//	panic(err.Error())
	//} else {
	//	fmt.Println("Last inserted ID:", lastInsertId)
	//}
	//
	//rowsAffected, err := result.RowsAffected()
	//if err != nil {
	//	panic(err.Error())
	//} else {
	//	fmt.Println("Rows affected:", rowsAffected)
	//}
	//定义要删除记录的ID
	name := "kite"

	//准备要执行的SQL语句
	stmt, err := db.Prepare("DELETE FROM test1 WHERE name=?")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	//执行SQL语句
	result, err := stmt.Exec(name)
	if err != nil {
		panic(err.Error())
	}

	//获取受影响的行数
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("成功删除%d条记录\n", rowsAffected)

	// 查看当前表中的所有记录
	// 执行 SQL 语句，展示当前表的所有记录
	rows, err := db.Query("SELECT * FROM test1")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	// 获取所有列名
	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}

	// 创建一个切片用于存储所有行
	values := make([]interface{}, len(columns))
	valuePtrs := make([]interface{}, len(columns))
	for rows.Next() {
		// 将一行的值放入 values 切片中，并将指向 values 中每个值的指针放入 valuePtrs 切片中
		for i := range columns {
			valuePtrs[i] = &values[i]
		}

		// 扫描一行的值并将它们放入 valuePtrs 切片中
		if err := rows.Scan(valuePtrs...); err != nil {
			panic(err.Error())
		}

		// 打印当前行的值
		for i, col := range columns {
			fmt.Printf("%s: %s\t", col, values[i])
		}
		fmt.Println()
	}
}
