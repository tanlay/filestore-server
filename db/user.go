package db

import (
	mydb "filestore-server/db/mysql"
	"fmt"
)

//UserSignup 用过用户名密码完成user表的注册
func UserSignup(username, passwd string) bool {
	stmt, err := mydb.DBConn().Prepare(`INSERT INTO tbl_user (user_name,user_pwd) values (?,?)`)
	if err != nil {
		fmt.Printf("Failed to insert,err: %s\n", err.Error())
		return false
	}
	defer stmt.Close()

	res, err := stmt.Exec(username,passwd)
	if err != nil {
		fmt.Printf("Failed to insert, err: %s\n", err.Error())
		return false
	}
	if rf, err := res.RowsAffected();nil==err && rf > 0 {
		return true
	}
	return false
}

