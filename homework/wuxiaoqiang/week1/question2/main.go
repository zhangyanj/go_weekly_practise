package main

import (
	"fmt"
	"log"
	"net/http"
)

/*
使用协程实现一个服务端，要求：
	通过浏览器访问响应Hello world
	不可以使用snow框架
*/
func main() {
	http.HandleFunc("/", helloWorld)
	log.Fatal(http.ListenAndServe(":9999", nil))
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world!")
}
