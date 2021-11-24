package db

import (
	mydb "filestore-server/db/mysql"
	"testing"
)


func TestMain(m *testing.M) {
	cleanTable()
	m.Run()
	cleanTable()
}

func cleanTable() {
	mydb.DBConn().Exec("truncate tbl_file")
}

func TestFileWorkFlow(t *testing.T) {
	t.Run("testOnFileUploadFinished", testOnFileUploadFinished)
}

func testOnFileUploadFinished(t *testing.T) {
	ok := OnFileUploadFinished(
		"ac92ff6ba783af448319acd8f568b621c0b348da",
		"R-C.jpg",
		279169,
		"E://golang-fileserver-tmp/R-C.jpg" )
	if ok != true {
		t.Errorf("Error of OnFileUploadFinished: %v", ok)
	}
}