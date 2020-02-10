package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// 初始化连接

// Open函数可能只是验证其参数，而不创建与数据库的连接。如果要检查数据源的名称是否合法，应调用返回值的Ping方法。

// 返回的DB可以安全的被多个goroutine同时使用，并会维护自身的闲置连接池。这样一来，Open函数只需调用一次。很少需要关闭DB。

// 定义一个全局对象db
var db *sql.DB

// 定义一个初始化数据库的函数
func initDB() (err error) {
	// DSN:Data Source Name
	// dsn := "user:password@tcp(127.0.0.1:3306)/test"
	dsn := "root:Kimalan91@tcp(127.0.0.1:3306)/learninggolang"
	// 不会校验账号密码是否正确
	// 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}


// 其中sql.DB是一个数据库（操作）句柄，代表一个具有零到多个底层连接的连接池
// 它可以安全的被多个go程同时使用
// database/sql包会自动创建和释放连接；它也会维护一个闲置连接的连接池。

// SetMaxOpenConns
// func (db *DB) SetMaxOpenConns(n int)
// SetMaxOpenConns设置与数据库建立连接的最大数目
// 如果n大于0且小于最大闲置连接数，会将最大闲置连接数减小到匹配最大开启连接数的限制
// 如果n<=0，不会限制最大开启连接数，默认为0（无限制）

// SetMaxIdleConns
// func (db *DB) SetMaxIdleConns(n int)
// SetMaxIdleConns设置连接池中的最大闲置连接数
// 如果n大于最大开启连接数，则新的最大闲置连接数会减小到匹配最大开启连接数的限制
// 如果n<=0，不会保留闲置连接


// CRUD
// 建库建表
// 我们先在MySQL中创建一个名为sql_test的数据库
// CREATE DATABASE sql_test;
// 进入该数据库:
// use sql_test;
// 执行以下命令创建一张用于测试的数据表：
// CREATE TABLE `user` (
//     `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
//     `name` VARCHAR(20) DEFAULT '',
//     `age` INT(11) DEFAULT '0',
//     PRIMARY KEY(`id`)
// )ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

// select id,name,age from user where id=1;

// 查询
// 单行查询
// 单行查询db.QueryRow()执行一次查询，并期望返回最多一行结果（即Row）。
// QueryRow总是返回非nil的值，直到返回值的Scan方法被调用时，才会返回被延迟的错误。（如：未找到结果）
// func (db *DB) QueryRow(query string, args ...interface{}) *Row
// 查询单条数据示例
type user struct {
	id int
	name string
	age int
}
func queryRowDemo(n int) {
	sqlStr := "select id, name, age from user where id=?"
	var u user
	// 非常重要：确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
	err := db.QueryRow(sqlStr, n).Scan(&u.id, &u.name, &u.age)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return
	}
	fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
}

// 多行查询
// 多行查询db.Query()执行一次查询，返回多行结果（即Rows），一般用于执行select命令。参数args表示query中的占位参数。
// func (db *DB) Query(query string, args ...interface{}) (*Rows, error)
// 查询多条数据示例
func queryMultiRowDemo() {
	sqlStr := "select id, name, age from user where id > ?"
	rows, err := db.Query(sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	// 非常重要：关闭rows释放持有的数据库链接
	defer rows.Close()

	// 循环读取结果集中的数据
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
	}
}

// 插入数据
// 插入、更新和删除操作都使用方法
// func (db *DB) Exec(query string, args ...interface{}) (Result, error)
// Exec执行一次命令（包括查询、删除、更新、插入等），返回的Result是对已执行的SQL命令的总结。参数args表示query中的占位参数。
// 插入数据具体代码：
func insertRowDemo() {
	sqlStr := "insert into user(name, age) values (?,?)"
	ret, err := db.Exec(sqlStr, "王五", 38)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	theID, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", theID)
}

// 更新数据
// 具体更新数据示例代码如下：
func updateRowDemo(i int) {
	sqlStr := "update user set age=? where id = ?"
	ret, err := db.Exec(sqlStr, 39, i)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("update success, affected rows:%d\n", n)
}

// 删除数据
// 具体删除数据的示例代码如下：
func deleteRowDemo(i int) {
	sqlStr := "delete from user where id = ?"
	ret, err := db.Exec(sqlStr, i)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("delete success, affected rows:%d\n", n)
}

// MySQL预处理
// 什么是预处理？

// 普通SQL语句执行过程：
// 客户端对SQL语句进行占位符替换得到完整的SQL语句。
// 客户端发送完整SQL语句到MySQL服务端
// MySQL服务端执行完整的SQL语句并将结果返回给客户端。

// 预处理执行过程：
// 把SQL语句分成两部分，命令部分与数据部分。
// 先把命令部分发送给MySQL服务端，MySQL服务端进行SQL预处理。
// 然后把数据部分发送给MySQL服务端，MySQL服务端对SQL语句进行占位符替换。
// MySQL服务端执行完整的SQL语句并将结果返回给客户端。

// 为什么要预处理？
// 优化MySQL服务器重复执行SQL的方法，可以提升服务器性能，提前让服务器编译，一次编译多次执行，节省后续编译的成本。
// 避免SQL注入问题。

// Go实现MySQL预处理
// func (db *DB) Prepare(query string) (*Stmt, error)
// Prepare方法会先将sql语句发送给MySQL服务端，返回一个准备好的状态用于之后的查询和命令。返回值可以同时执行多个查询和命令。
// 查询操作的预处理示例代码如下：
// // 预处理查询示例
// func prepareQueryDemo() {
// 	sqlStr := "select id, name, age from user where id > ?"
// 	stmt, err := db.Prepare(sqlStr)
// 	if err != nil {
// 		fmt.Printf("prepare failed, err:%v\n", err)
// 		return
// 	}
// 	defer stmt.Close()
// 	rows, err := stmt.Query(0)
// 	if err != nil {
// 		fmt.Printf("query failed, err:%v\n", err)
// 		return
// 	}
// 	defer rows.Close()
// 	// 循环读取结果集中的数据
// 	for rows.Next() {
// 		var u user
// 		err := rows.Scan(&u.id, &u.name, &u.age)
// 		if err != nil {
// 			fmt.Printf("scan failed, err:%v\n", err)
// 			return
// 		}
// 		fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
// 	}
// }

// 插入、更新和删除操作的预处理十分类似，这里以插入操作的预处理为例：
// // 预处理插入示例
// func prepareInsertDemo() {
// 	sqlStr := "insert into user(name, age) values (?,?)"
// 	stmt, err := db.Prepare(sqlStr)
// 	if err != nil {
// 		fmt.Printf("prepare failed, err:%v\n", err)
// 		return
// 	}
// 	defer stmt.Close()
// 	_, err = stmt.Exec("小王子", 18)
// 	if err != nil {
// 		fmt.Printf("insert failed, err:%v\n", err)
// 		return
// 	}
// 	_, err = stmt.Exec("沙河娜扎", 18)
// 	if err != nil {
// 		fmt.Printf("insert failed, err:%v\n", err)
// 		return
// 	}
// 	fmt.Println("insert success.")
// }

func main() {
	err := initDB() // 调用输出化数据库的函数
	if err != nil {
		fmt.Printf("init db failed,err:%v\n", err)
		return
	}
	queryRowDemo(2)
	// insertRowDemo()
	updateRowDemo(6)
	// deleteRowDemo(4)
	queryMultiRowDemo()
}