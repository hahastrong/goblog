package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Hello, 这里是Goblog</h1>")
	} else if r.URL.Path == "/about" {
		fmt.Fprint(w, "<h1>此博客用以记录编程笔记，如有问题请联系</h1> <a href=\"www.baidu.com\"> www.baidu.com </a>")
	} else {
		fmt.Fprint(w, "未找到相关页面")
	}

	fmt.Fprint(w, "请求路径为: "+r.URL.Path)
}

// func handlerFuncAbout(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "<h1>about</h1>")
// 	//fmt.Fprint(w, "请求路径为: "+r.URL.Path)
// }

func main() {
	http.HandleFunc("/", handlerFunc)
	// http.HandleFunc("/about", handlerFuncAbout)
	http.ListenAndServe(":3000", nil)
}
