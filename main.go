package main

import (
	"path/filepath"

	"github.com/adimyth/go-autodoc/autodoc"
)

func main() {
	files, err := filepath.Glob("src/*.go")
	if err != nil {
		panic("error listing all go files")
	}
	for _, file := range files {
		autodoc.DocGenerator(file)
	}
}
