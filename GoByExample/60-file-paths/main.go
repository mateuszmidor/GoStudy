package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println(p) // dir1/dir2/filename

	fmt.Println(filepath.Join("dir1//", "filename"))       // dir1/filename
	fmt.Println(filepath.Join("dir1/../dir1", "filename")) // dir1/filename

	fmt.Println("Dir(p):", filepath.Dir(p))   // dir1/dir2
	fmt.Println("Base(p):", filepath.Base(p)) // filename

	fmt.Println(filepath.IsAbs("dir/file"))  // false
	fmt.Println(filepath.IsAbs("/dir/file")) // true

	filename := "config.json"

	ext := filepath.Ext(filename)
	fmt.Println(ext) // .json

	fmt.Println(strings.TrimSuffix(filename, ext)) // config

	rel, _ := filepath.Rel("a/b", "a/b/t/file") // t/file
	fmt.Println(rel)

	rel, _ = filepath.Rel("a/b", "a/c/t/file") // ../c/t/file
	fmt.Println(rel)
}
