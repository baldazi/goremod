# goremod

`goremod` is a Go tool designed to **automatically replace the module name and import paths** in your Go project files (`go.mod` and `.go` files).
It simplifies the process of migrating or renaming a Go module by updating all relevant references with a single command.

---

## 🚀 Features

* Replace module name in `go.mod`
* Update import paths in `.go` files (single-line and multi-line import blocks)
* Process multiple files or entire projects
* Simple CLI interface for quick renaming

---

## 🔧 Installation

```bash
go install github.com/baldazi/goremod@latest
```

---

## 📖 Usage

### CLI usage

Use the following command to replace the module name and import paths in your project:

```bash
go remod "new/module/name"
```

This command updates:

* The `module` directive in your `go.mod` file
* All import statements referencing the old module in `.go` files

### Example

If your current module is `old/module`, and you want to rename it to `new/module`, run:

```bash
go remod "new/module"
```

---

### Programmatic usage

You can also use the `goremod` package directly in Go code:

<!-- ```go
package main

import "github.com/baldazi/goremod/utils"

func main() {
    files := []string{"main.go", "utils/log.go", "core/replace.go"}
    oldImport := "old/module"
    newImport := "new/module"

    utils.ProcessFiles(files, oldImport, newImport)
}
``` -->

## 🤝 Contributing

Contributions are welcome! Feel free to open issues or pull requests.

---

## 📜 License

This project is licensed under the MIT License. See [LICENSE](LICENSE) for details.
