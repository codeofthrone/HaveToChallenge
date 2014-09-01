package main

import (
	"flag"
	"fmt"
	"os"
	"io/ioutil"
	"path/filepath"
)

var fstr string

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func visit(path string, f os.FileInfo, err error) error {
	fstr = fstr + path +"\n"
	//fmt.Printf("Visited: %s\n", path)
  return nil
} 

func main() {
	var root string
	flag.StringVar(&root ,"path", "path", "path string")
	flag.Parse()
	fmt.Println("path:", root)
	f,err := os.Create("./output.txt")
    check(err)

    err = filepath.Walk(root, visit)
    f.WriteString(fstr)
    defer f.Close()

	dat, err := ioutil.ReadFile("./output.txt")
    check(err)
    fmt.Print(string(dat))

 //   err = filepath.Walk(root, visit)

}
