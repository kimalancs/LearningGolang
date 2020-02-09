package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // 导入但没有直接使用，使用其init()函数
)

// 面试题
// 如何判断一个链表有没有闭环
// 每一个节点会记录下一个节点的地址或指针
type a struct {
	val  int
	next *a
}

// 从出发开始，x一次走一步，y一次走两步，当x与y在某节点相遇时说明有闭环
// 经常在leetcode刷算法题

// MySQL，关系型数据库
// 支持插件式的存储引擎
// MyISAM
// 查询速度快、只支持表锁、不支持事务
// InnoDB
// 整体速度快、支持表锁和行锁、支持事务

// 事务
// 索引的原理：B树和B+树
// 索引的类型
// 索引的命中
// 分库分表
// SQL注入
// SQL慢查询优化
// MySQL主从，所有写入操作都写入主库，操作记入日志binlog，从库根据传输过来的binlog的日志做一遍所有操作
// MySQL读写分离

// Go语言中的database/sql包提供了保证SQL或类SQL数据库的泛用接口，并不提供具体的数据库驱动
// 使用database/sql包时必须注入（至少）一个数据库驱动

// 原生支持连接池，是并发安全的

// 只有一个标准，列出了需要第三方库实现的具体内容，并没有具体实现

// 常用的数据库基本上都有完整的第三方实现
// go get -u github.com/go-sql-driver/mysql

// func Open(driverName, dataSourceName string) (*DB, error)
// Open打开一个dirverName指定的数据库，dataSourceName指定数据源，一般包至少括数据库文件名和（可能的）连接信息。

func main() {
	// DSN:Data Source Name
	dsn := "user:password@tcp(127.0.0.1:3306)/dbname"
	db, err := sql.Open("mysql", dsn) // 不会校验用户名和密码是否正确，只会校验提供的参数格式是否正确，Ping()方法用来验证
	if err != nil {
		panic(err) // dsn格式不正确,无法解析出需要的数据时才会报错
	}
	err = db.Ping() // 验证用户名和密码是否正确，是否可以连接数据库
	if err != nil {
		fmt.Printf("open %s failed, err:%v\n", dsn, err)

	}
	defer db.Close() // 注意,defer db.Close()语句不应该写在if err != nil的前面呢？
}

// vscode 跳转到定义后再返回
// windows系统 alt+左右方向箭头
// mac control+减号和contrl+shift+减号
