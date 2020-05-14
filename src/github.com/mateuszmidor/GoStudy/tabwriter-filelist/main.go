package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"text/tabwriter"
)

func lsdir(dir string) []os.FileInfo {
	files, _ := ioutil.ReadDir(dir)
	return files
}

// file/dir to String
func fileOrDir(dir bool) string {
	if dir == true {
		return "DIR"
	}
	return "FILE"
}

// prepare tabwriter for writing file list in form: Name  Size  Type
func makeTabWriter(format string) *tabwriter.Writer {
	const minWidth = 0  // minimal cell width including any padding
	const tabWidth = 2  // width of tab characters (equivalent number of spaces)
	const padding = 4   // distance between cells
	const padchar = ' ' // SCII char used for padding
	const flags = 0     // formatting control
	w := tabwriter.NewWriter(os.Stdout, minWidth, tabWidth, padding, padchar, flags)
	return w
}

func main() {
	// prepare tabwriter
	const format = "%v\t%v\t%v\t\n"
	w := makeTabWriter(format)

	// get sorted list of files
	files := lsdir("/boot/")
	byType := func(i, j int) bool { return files[i].IsDir() && !files[j].IsDir() }
	sort.Slice(files, byType)

	// print files using tabwriter
	fmt.Fprintf(w, format, "Name", "Size", "Type")
	fmt.Fprintf(w, format, "----", "----", "----")
	for _, fi := range files {
		fmt.Fprintf(w, format, fi.Name(), fi.Size(), fileOrDir(fi.IsDir()))
	}

	w.Flush()
}
