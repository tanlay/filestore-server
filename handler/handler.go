package handler

import (
	"encoding/json"
	"filestore-server/meta"
	"filestore-server/util"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

//FileUploadHandler 处理文件上传
func FileUploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		//返回上传html页面
		data, err := ioutil.ReadFile("./static/view/index.html")
		if err != nil {
			io.WriteString(w, "internel server error")
		}
		io.WriteString(w, string(data))
	} else if r.Method == "POST" {
		//接收用户上传的文件流并储存到本地目录
		file, head, err := r.FormFile("file")
		if err != nil {
			fmt.Printf("Failed to get data err: %s\n", err.Error())
			return
		}
		defer file.Close()

		fileMeta := meta.FileMeta{
			FileName: head.Filename,
			Location: "E://golang-fileserver-tmp/"+head.Filename,
			UploadAt: time.Now().Format("2016-01-02 15:04:05"),
		}

		newFile, err := os.Create(fileMeta.Location)
		if err != nil {
			fmt.Printf("Failed to create file, err: %s\n", err.Error())
			return
		}
		defer newFile.Close()

		fileMeta.FileSize, err = io.Copy(newFile, file)
		if err != nil {
			fmt.Printf("Failed to save data into file, err: %s\n", err.Error())
			return
		}
		newFile.Seek(0,0)
		fileMeta.FileSha1 = util.FileSha1(newFile)
		//meta.UpdateFileMeta(fileMeta)
		_ = meta.UpdateFileMetaDB(fileMeta)

		http.Redirect(w, r, "/file/upload/suc", http.StatusFound)
	}
}

// UploadSucHandler 上传完成
func UploadSucHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Upload finished!")
}

// GetFileMetaHandler 通过filehash获取文件元信息
func GetFileMetaHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	filehash := r.Form["filehash"][0]
	//fMeta := meta.GetFileMeta(filehash)
	fMeta, err := meta.GetFileMetaDB(filehash)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	data, err := json.Marshal(fMeta)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fsha1 := r.Form.Get("filehash")
	fm := meta.GetFileMeta(fsha1)
	f, err := os.Open(fm.Location)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//设置浏览器下载头
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment")

	w.Write(data)
}

// FileMetaUpdateHandler 更新元信息接口（重命名）
func FileMetaUpdateHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	opType := r.Form.Get("op")
	fileSha1 := r.Form.Get("filehash")
	newFileName:= r.Form.Get("filename")

	if opType != "0" {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	if r.Method != "POST"{
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	curFileMeta:= meta.GetFileMeta(fileSha1)
	curFileMeta.FileName = newFileName
	meta.UpdateFileMeta(curFileMeta)

	data, err := json.Marshal(curFileMeta)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

// FileDeleteHandler 删除文件及元信息
func FileDeleteHandler(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	fileSha1:= r.Form.Get("filehash")
	fMeta:= meta.GetFileMeta(fileSha1)
	os.Remove(fMeta.Location)

	meta.RemoveFileMeta(fileSha1)



	w.WriteHeader(http.StatusOK)
}