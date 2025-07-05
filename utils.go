package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	lg "github.com/charmbracelet/lipgloss"
)

func listFile(dir string) []string {
	var files []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".go") {
			files = append(files, path)
		}

		return nil
	})

	if err != nil {
		log.Fatalf("Error Listing files : %v", err)
	}
	return files
}

// styling CLI

func fileHeader(filename string, msg string) string {
	return lg.NewStyle().
		Foreground(lg.Color(69)).
		Bold(true).
		Render(fmt.Sprintf("%s : %s", msg, filename))
}

func statusDone(lines int, statusMsg string, msg string) string {
	check := lg.NewStyle().
		Foreground(lg.Color("42")).
		Bold(true).
		Render(statusMsg)

	count := lg.NewStyle().
		Foreground(lg.Color("241")).
		Bold(true).
		Render(fmt.Sprintf("%d %msg", lines, msg))

	return fmt.Sprintf("%s %s", check, count)
}

func statusNoChange(msg string) string {
	return lg.NewStyle().
		Foreground(lg.Color("245")).
		Render(msg)
}
