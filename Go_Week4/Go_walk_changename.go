package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

//var fstr string

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func visit(path string, f os.FileInfo, err error) error {
	//fstr = path
	fmt.Printf("Visited: %s\n", path)
	return nil
}

func main() {
	FileCount := 0
	TotalCount := 0
	var root string
	flag.StringVar(&root, "path", "path", "path string")
	flag.Parse()
	fmt.Println("path:", root)
	OrgPath := root
	Recurse := true
	var newpath string
	var TypeName []string
	//TypeCount := 0
	walkFn := func(path string, info os.FileInfo, err error) error {
		stat, err := os.Stat(path)

		replace_strings := strings.Replace(OrgPath, "/", "_", -1)
		if err != nil {
			return err
		}

		if stat.IsDir() && path != OrgPath && !Recurse {
			fmt.Println("skipping dir:", path)
			return filepath.SkipDir
		}

		if err != nil {
			return err
		}
		if stat.IsDir() {
			FileCount = 0
			newpath = strings.Replace(path, "/", "_", -1)
			newpath = strings.Replace(newpath, replace_strings, "", -1)
			TypeName = strings.Split(newpath, "_")

			fmt.Println(newpath)
		}
		if !stat.IsDir() {
			if strings.Contains(path, "txt") {
				TotalCount += 1
				FileCount += 1

				newfilename := fmt.Sprintf("%05d_%s_%03d", TotalCount, newpath, FileCount)
				fmt.Println(newfilename)
				fmt.Println("org file name : ", path)
				//fmt.Println(filepath.Base(path))
				//fmt.Println(filepath.Dir(path))
				newfullname := filepath.Dir(path) + "/" + newfilename + ".jpg"
				fmt.Println("new file name : ", newfullname)
				//os.Rename(path, newfullname)

			}

		}

		return nil
	}
	err := filepath.Walk(OrgPath, walkFn)
	if err != nil {
		log.Fatal(err)
	}

}
