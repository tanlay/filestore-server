package meta

import (
	mydb "filestore-server/db"
)

//FileMeta 文件原信息结构体
type FileMeta struct {
	FileSha1	string
	FileName 	string
	FileSize	int64
	Location	string
	UploadAt	string
}

var fileMetas map[string]FileMeta

func init() {
	fileMetas = make(map[string]FileMeta)
}

//UpdateFileMeta 新增/更新文件元信息
func UpdateFileMeta(fmeta FileMeta) {
	fileMetas[fmeta.FileSha1] = fmeta
}

//UpdateFileMetaDB 新增/更新文件元信息到mysql
func UpdateFileMetaDB(fmeta FileMeta) bool {
	return mydb.OnFileUploadFinished(
		fmeta.FileSha1,fmeta.FileName,fmeta.FileSize,fmeta.Location)
}


//GetFileMeta 通过sha1获取文件对象
func GetFileMeta(fileSha1 string) FileMeta {
	return fileMetas[fileSha1]
}

//RemoveFileMeta 删除元信息
func RemoveFileMeta(fileSha1 string) {
	delete(fileMetas, fileSha1)
}