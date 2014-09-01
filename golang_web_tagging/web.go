package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type File struct {
	FileName string
}

func tagging(w http.ResponseWriter, r *http.Request) {
	fmt.Println("tag=======start======")
	//请求的是登陆数据，那么执行登陆的逻辑判断
	fmt.Println("tag1:", r.Form["tag1"])
	fmt.Println("tag2:", r.Form["tag2"])
	fmt.Println("tag3:", r.Form["tag3"])

	filses, _ := ioutil.ReadDir("./images")
	var Filename string = filses[0].Name()

	r.ParseForm() //解析url传递的参数，对于POST则解析响应包的主体（request body）
	//注意:如果没有调用ParseForm方法，下面无法获取表单的数据
	fmt.Println("path", r.URL.Path)
	var tag string = ""
	var tag_tmp string = ""
	for k, v := range r.Form {
		//fmt.Fprintf(w, "key:"+ k)
		fmt.Println("key:", k)
		//fmt.Fprintf(w, "val:"+ strings.Join(v, ""))
		fmt.Println("val:", strings.Join(v, ""))
		tag_tmp = strings.Trim(strings.Join(v, ""), " ")
		if tag_tmp != "" {
			tag = tag + "_" + tag_tmp
		}
		fmt.Println("tag", tag)
	}
	// fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的

	// move file
	if tag != "" {
		var oldpath string = "images/" + Filename
		var ext string = filepath.Ext(Filename)
		var newbasename = Filename[:len(Filename)-len(ext)]
		fmt.Println("newbasename : ", newbasename)
		fmt.Println("ext : ", ext)
		var newpath string = "processed/" + newbasename + tag + ext

		var err error = os.Rename(oldpath, newpath)

		if err != nil {
			fmt.Println("failed")
		}

		Filename = filses[1].Name()
		fmt.Println("Filename", Filename)
	}

	fmt.Println("method:", r.Method) //获取请求的方法
	fmt.Println("Filename", Filename)
	t := template.New("fieldname example")
	//t, _ := template.ParseFiles("login.gtpl"  )
	var htmlcss string = `
    <style type="text/css">
    img {
        max-width: 700px; 
        max-height: 700px; 
    }
    </style>`
	var htmlhead string = "<html><head><title></title>" + htmlcss + "</head><body><form action='/tagging' method='post'><table>"
	var htmlinput string = "<td><tr>TAG1:<input type='text' name='tag1'></tr><tr>TAG2:<input type='text' name='tag2'></tr><tr>TAG3:<input type='text' name='tag3'> "
	var htmlsubmit string = "<input type='submit' value='rename'></tr></td>"
	var htmlfoot string = "</table></form></body></html>"
	t.Parse(htmlhead + "<tr><td><img src=\"/static/{{.FileName}}\" >{{.FileName}}<br></td>" + htmlinput + htmlsubmit + htmlfoot)

	p := File{FileName: Filename}
	t.Execute(w, p)

}

func main() {
	http.HandleFunc("/", tagging) //设置访问的路由
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("./images"))))
	err := http.ListenAndServe(":22629", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
