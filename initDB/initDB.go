package initDB

import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import "log"

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/ginhello")

	if err != nil {
		log.Panicln("err:", err.Error())
	}

	// 设置最大的并发打开连接数为10。
	Db.SetMaxOpenConns(10)
	// 设置最大的空闲连接数为10。
	Db.SetMaxIdleConns(10)
}
