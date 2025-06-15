package main

import (
	"flag"
	"fmt"
	"golang.org/x/mod/modfile"
	"log"
	"os"
)

func main() {
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

	fmt.Println("nom du module :", mf.Module.Mod.Path)
	fmt.Println("hello world!", *from, *to)
}
