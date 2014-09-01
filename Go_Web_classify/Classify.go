package main

import (
	"fmt"
	//"html/template"
	//"io/ioutil"
	"log"
	"net/http"
	//"os"
	//"path/filepath"
	//"strings"
)

func classify(w http.ResponseWriter, r *http.Request) {
	fmt.Println(123)
}

func main() {
	http.HandleFunc("/", classify) //设置访问的路由
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("./images"))))
	err := http.ListenAndServe(":22629", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
