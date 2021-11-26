package handler

import (
	dblayer "filestore-server/db"
	"filestore-server/util"
	"io/ioutil"
	"net/http"
)

var (
	pwdSalt = "25@*4xfe"
)

//UserSignupHandler 处理用户注册请求
func UserSignupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet{
		data, err := ioutil.ReadFile("./static/view/signup.html")
		if err != nil{
			w.WriteHeader(500)
			return
		}
		w.Write(data)
	}
	r.ParseForm()
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	if len(username) < 3 || len(password) < 5{
		w.Write([]byte("用户名密码格式不正确"))
		return
	}
	encPasswd := util.Sha1([]byte(password + pwdSalt))
	suc := dblayer.UserSignup(username, encPasswd)
	if suc {
		w.Write([]byte("SUCCESS"))
	} else {
		w.Write([]byte("FAILED"))
	}
}
