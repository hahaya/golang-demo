package main

import (
	"fmt"
	"net/http"
)

func main() {
	//在8080端口开启一个简单的http服务
	http.HandleFunc("/", httpHandle)
	http.ListenAndServe(":8080", nil)
}

func httpHandle(w http.ResponseWriter, r *http.Request) {
	//定义cookie结构体
	ck := http.Cookie{Name: "username", Value: "hahaya", Path: "/", Domain: "localhost", MaxAge: 120}
	//设置cookie
	http.SetCookie(w, &ck)

	//读取cookie
	_ck, err := r.Cookie("username")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	fmt.Fprintf(w, _ck.Value)
}
