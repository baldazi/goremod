package main

import (
	"fmt"
	"flag"
)

func main(){

	from := flag.String("from", "", "current module name")
	to := flag.String("to", "", "new module name")

	flag.Parse()

	fmt.Println("hello world!", *from, *to)
}