package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
	
)

// sqlx使用
// 第三方库sqlx能够简化操作，提高开发效率。

// 安装
// go get github.com/jmoiron/sqlx

// 基本使用
// 连接数据库
var db *sqlx.DB

type user struct {
	ID   int
	Name string
	Age  int
}

func initDB() (err error) {
	dsn := "user:password@tcp(127.0.0.1:3306)/test"
	// 也可以使用MustConnect连接不成功就panic
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	return
}

// 查询
// 查询单行数据示例代码如下：
// 查询单条数据示例
func queryRowDemo() {
	sqlStr := "select id, name, age from user where id=?"
	var u user
	err := db.Get(&u, sqlStr, 1) // Get方法通过反射了解user结构体中的字段有哪些（user字段要首字母大写）
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	fmt.Printf("id:%d name:%s age:%d\n", u.ID, u.Name, u.Age)
}

// 查询多行数据示例代码如下：
// 查询多条数据示例
func queryMultiRowDemo() {
	sqlStr := "select id, name, age from user where id > ?"
	var users []user
	err := db.Select(&users, sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	fmt.Printf("users:%#v\n", users)
}

// 插入、更新和删除
// sqlx中的Exec方法与原生sql中的Exec使用基本一致：

// 插入数据
func insertRowDemo() {
	sqlStr := "insert into user(name, age) values (?,?)"
	ret, err := db.Exec(sqlStr, "沙河小王子", 19)
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
func updateRowDemo() {
	sqlStr := "update user set age=? where id = ?"
	ret, err := db.Exec(sqlStr, 39, 6)
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
func deleteRowDemo() {
	sqlStr := "delete from user where id = ?"
	ret, err := db.Exec(sqlStr, 6)
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

// 事务操作
// 对于事务操作，我们可以使用sqlx中提供的db.Beginx()和tx.MustExec()方法来简化错误处理过程。示例代码如下：

func transactionDemo() {
	tx, err := db.Beginx() // 开启事务
	if err != nil {
		if tx != nil {
			tx.Rollback()
		}
		fmt.Printf("begin trans failed, err:%v\n", err)
		return
	}
	sqlStr1 := "Update user set age=40 where id=?"
	tx.MustExec(sqlStr1, 2)
	sqlStr2 := "Update user set age=50 where id=?"
	tx.MustExec(sqlStr2, 4)
	err = tx.Commit() // 提交事务
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("commit failed, err:%v\n", err)
		return
	}
	fmt.Println("exec trans success!")
}

// 不同的数据库中，SQL语句使用的占位符语法不尽相同。
// MySQL		?
// PostgreSQL	$1, $2等
// SQLite		? 和$1
// Oracle		:name

// SQL注入
// 我们任何时候都不应该自己拼接SQL语句！
// 这里我们演示一个自行拼接SQL语句的示例，编写一个根据name字段查询user表的函数如下：

// sql注入示例
func sqlInjectDemo(name string) {
	sqlStr := fmt.Sprintf("select id, name, age from user where name='%s'", name)
	fmt.Printf("SQL:%s\n", sqlStr)

	var users []user
	err := db.Select(&users, sqlStr)
	if err != nil {
		fmt.Printf("exec failed, err:%v\n", err)
		return
	}
	for _, u := range users {
		fmt.Printf("user:%#v\n", u)
	}
}
// 此时以下输入字符串都可以引发SQL注入问题：
// sqlInjectDemo("xxx' or 1=1#")
// sqlInjectDemo("xxx' union select * from user #")
// sqlInjectDemo("xxx' and (select count(*) from user) <10 #")

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init db failed, err:%v\n", err)
	}
	queryRowDemo()
	transactionDemo()
	queryMultiRowDemo()

}
