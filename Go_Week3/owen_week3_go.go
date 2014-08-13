package main

import (
	"strings"
	"fmt"
	"os"
)

var bigstr string

func recurs(n string, UpperFlag bool) string {
	if len(n) != 0 {
		if UpperFlag==true {
			fmt.Println(strings.ToUpper(n))
			strings.ToUpper(n)
			bigstr=bigstr+" "+strings.ToUpper(n)
			UpperFlag=false
		}else{
			UpperFlag=true
			fmt.Println(strings.ToLower(n))
			strings.ToLower(n)
			bigstr=bigstr+" "+strings.ToLower(n)
		}
		size := len(n)
		n = n[0 : size-1]
		recurs(n ,UpperFlag)
	}
	return n
}

func main() {
	args := os.Args[1]
	if args == "" {
		fmt.Println("Please insert Argument!")
	} else {
		fmt.Println("Insert args is :" + args)
	}
	UpperFlag:=true
	recurs(args, UpperFlag )
	answer:=strings.Fields(bigstr)

	fmt.Println("answer",answer)
	fmt.Println("answer len",len(answer))
}
