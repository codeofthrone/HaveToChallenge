package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1]
	if args == "" {
		fmt.Println("Please insert Argument!")
	} else {
		fmt.Println("Insert args is :" + args)
	}
	iargs, err := strconv.Atoi(args)
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
	for i := 0; i <= iargs; i++ {
		if (i % 3) == 0 {
			fmt.Println(i)
		}
	}

}
