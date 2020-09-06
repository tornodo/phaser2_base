package main

import (
	"fmt"
	"net/http"
	"os"
	"text/template"
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

	fmt.Println("端口号：", port, "， 用浏览器访问localhost:"+port)
	fmt.Println("如果有错误信息，可能是端口号冲突。记得每次修改文件后强制刷新几次，以免加载缓存导致页面没有改变，比如http://localhost:8080/?3，问号后面数字随便填写，强制浏览器重新加载")

	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./"))))
	http.Handle("/assets", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))
	http.Handle("/js", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))
	// http.HandleFunc("/", indexHandler)

	err := http.ListenAndServe(ip+":"+port, nil)
	if err != nil {
		fmt.Println(err)
	}

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./index.html")
	if err != nil {
		fmt.Println(err)
	} else {
		t.Execute(w, nil)
	}

}
