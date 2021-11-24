package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
	err error
)

func init() {
	dbStr := `root:123456@tcp(127.0.0.1:3308)/fileserver?charset=utf8`
	db, _ = sql.Open("mysql", dbStr)
	db.SetMaxOpenConns(1000)
	if err := db.Ping();nil != err {
		panic("conn db failed")
	}
}

func DBConn() *sql.DB {
	return db
}