package main

import (
	"fmt"
	"net/http"
	"os"
)

var (
	ip   = "0.0.0.0"
	port = "8080"
)

func main() {
	args := os.Args
	if len(args) == 2 {
		port = args[1]
	}
	http.Handle("/", http.FileServer(http.Dir("./")))
	fmt.Println("端口号：", port, "， 用浏览器访问localhost:"+port)
	fmt.Println("如果有错误信息，可能是端口号冲突。记得每次修改文件后强制刷新几次，以免加载缓存导致页面没有改变，比如http://localhost:8080/?3，问号后面数字随便填写，强制浏览器重新加载")
	err := http.ListenAndServe(ip+":"+port, nil)
	if err != nil {
		fmt.Println(err.Error())
	}

}
