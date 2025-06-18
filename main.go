package main

import (
	"flag"
	"fmt"
	"golang.org/x/mod/modfile"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
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

	// changing the module name
	cmd := exec.Command("go", "mod", "edit", "-module", newModule)
	output, err := cmd.CombinedOutput()

	if err != nil {
		log.Fatalf("Failed to run command: %v\nOutput: %s", err, string(output))
	}

	// find all go file

	root := "./" // root directory

	err = filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && strings.HasSuffix(d.Name(), ".go") {
			fmt.Println(path)
		}
		return nil
	})

	if err != nil {
		fmt.Println("An error occur when exploring :", err)
	}
}
