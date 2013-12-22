package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func main() {
	OrgPath := "c:/HTC/test2/"
	Recurse := true
	testcase := make([]string, 1500, 1500)
	TestCaseNumber := 0
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
		}
		if !stat.IsDir() {
			TypeName := strings.Split(path, "\\")
			if len(TypeName) > 3 && TypeName[3] != "Scenario" {
				newpath := strings.Replace(filepath.Dir(path), "\\", "_", -1)
				newpath = strings.Replace(newpath, replace_strings, "", -1)
				if !stringInSlice(newpath, testcase) || TestCaseNumber == 0 {
					TestCaseNumber++
					testcase[TestCaseNumber] = newpath
					fmt.Println(testcase[TestCaseNumber])
				}

			}
		}

		return nil
	}
	fmt.Println(testcase)

	err := filepath.Walk(OrgPath, walkFn)
	if err != nil {
		log.Fatal(err)
	}

}
