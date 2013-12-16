package main

import (

	"fmt"
	"log"
	"os"

	"path/filepath"
	"strings"
)




func main() {
    FileCount := 0
	OrgPath := "c:\\HTC\\test2\\"
	Recurse := true
	walkFn := func(path string, info os.FileInfo, err error) error {
		stat, err := os.Stat(path)
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
		if !stat.IsDir() {
			filefullname := strings.Split(path, OrgPath)

			fmt.Println(FileCount)
			fmt.Println(filefullname)
			fmt.Println(path)
			FileCount+=1
		}

		return nil
	}
	err := filepath.Walk(OrgPath, walkFn)
	if err != nil {
		log.Fatal(err)
	}

}
