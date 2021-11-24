package db

import (
	"database/sql"
	mydb "filestore-server/db/mysql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)


// OnFileUploadFinished 文件上传完成，保存meta到mysql
func OnFileUploadFinished(filehash, filename string, filesize int64, fileaddr string) bool {
	stmtIns, err := mydb.DBConn().Prepare(`INSERT ignore INTO tbl_file 
    (file_sha1, file_name, file_size, file_addr, status) values (?,?,?,?,1)`)
	if err != nil {
		fmt.Println("Failed to prepare statement, err:"+ err.Error())
		return false
	}
	defer stmtIns.Close()
	res, err := stmtIns.Exec(filehash,filename,filesize,fileaddr)
	if err != nil{
		fmt.Println(err.Error())
		return false
	}
	if rf,err := res.RowsAffected();nil == err {
		if rf <= 0 {
			fmt.Printf("File with hash:%s has been uploaded\n", filehash)
		}
		return true
	}
	return false
}

type TableFile struct {
	FileHash	string
	FileName	sql.NullString
	FileSize	sql.NullInt64
	FileAddr	sql.NullString
}

// GetFileMetaDB 从mysql获取文件元信息
func GetFileMetaDB(filehash string) (*TableFile, error) {
	dbStr := `SELECT file_sha1,file_name,file_size,file_addr from tbl_file where file_sha1=? and status=1`
	stmtGet, err := mydb.DBConn().Prepare(dbStr)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer stmtGet.Close()
	tfile := TableFile{}
	err = stmtGet.QueryRow(filehash).Scan(&tfile.FileHash,&tfile.FileName,&tfile.FileSize,&tfile.FileAddr)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return &tfile, nil
}
