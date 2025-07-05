package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"golang.org/x/mod/modfile"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: goremod <new/module/name>")
	}

	// Ret the new module name from cli argument
	newModule := os.Args[1]

	// Read the module file to ensure there is a Go module in the working directory
	data, err := os.ReadFile("go.mod")
	if err != nil {
		log.Fatal("Cannot find go.mod — make sure you're inside a Go module directory.")
	}
	// get the module file informations
	mf, err := modfile.Parse("go.mod", data, nil)
	if err != nil {
		log.Fatal("Failed to parse go.mod: invalid or corrupted format.")
	}
	// get the current module name
	curModule := mf.Module.Mod.Path

	// process the replacement
	files := listFile(".")
	processFiles(files, curModule, newModule)

	// changing the module name
	cmd := exec.Command("go", "mod", "edit", "-module", newModule)
	output, err := cmd.CombinedOutput()

	if err != nil {
		log.Fatalf("Failed to run command: %v\nOutput: %s", err, string(output))
	}
}

func processFiles(files []string, oldImportBase string, newImportBase string) {

	for _, f := range files {
		fmt.Println(fileHeader(f, "➜ File"))

		content, err := os.ReadFile(f)

		if err != nil {
			log.Fatalf("Error occur when reading file %s : %v", f, err)
		}

		lines := strings.Split(string(content), "\n")
		inImportBlock := false
		changedLines := 0

		for i, line := range lines {
			trimmed := strings.TrimSpace(line)

			if strings.HasPrefix(trimmed, "import") {
				if strings.HasPrefix(trimmed, "import (") {
					inImportBlock = true
					continue
				}

				if strings.Contains(trimmed, oldImportBase) {
					lines[i] = strings.ReplaceAll(line, oldImportBase, newImportBase)
					changedLines++
				}

			} else if inImportBlock {
				if trimmed == ")" {
					inImportBlock = false
					continue
				}
				if strings.Contains(trimmed, oldImportBase) {
					lines[i] = strings.ReplaceAll(line, oldImportBase, newImportBase)
					changedLines++
				}
			}

			if changedLines > 0 {
				if err := os.WriteFile(f, []byte(strings.Join(lines, "\n")), 0644); err != nil {
					log.Fatalf("Error writing : %v", err)
				}
				fmt.Println(statusDone(changedLines, "✔️ Modified", "Lines Replaced"))
			}else{
				fmt.Println(statusNoChange("No replacement is needed"))
			}

		}
	}
}
