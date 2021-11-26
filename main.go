package main

import (
	"filestore-server/handler"
	"fmt"
	"net/http"
)

func main() {
	//文件
	http.HandleFunc("/file/upload", handler.FileUploadHandler)
	http.HandleFunc("/file/upload/suc", handler.UploadSucHandler)
	http.HandleFunc("/file/meta", handler.GetFileMetaHandler)
	http.HandleFunc("/file/download", handler.DownloadHandler)
	http.HandleFunc("/file/update", handler.FileMetaUpdateHandler)
	http.HandleFunc("/file/delete", handler.FileDeleteHandler)
	//用户
	http.HandleFunc("/user/signup", handler.UserSignupHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Failed to start server: err %s\n", err.Error())
	}
}
