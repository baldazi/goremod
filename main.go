package main

import (
	"flag"
	"fmt"
	"golang.org/x/mod/modfile"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: goremod <new/module/name>")
	}
	newModule := os.Args[1]
	// define all flags
	from := flag.String("from", "", "current module name")
	to := flag.String("to", "", "new module name")
	flag.Parse()
	// parse mod file
	data, err := os.ReadFile("go.mod")
	if err != nil {
		log.Fatal("Cannot find go.mod — make sure you're inside a Go module directory.")
	}

	mf, err := modfile.Parse("go.mod", data, nil)

	if err != nil {
		log.Fatal("Failed to parse go.mod: invalid or corrupted format.")
	}
	curModule := mf.Module.Mod.Path

	fmt.Println("nom du module :", curModule)
	fmt.Println(newModule)
	fmt.Println("hello world!", *from, *to)
}
