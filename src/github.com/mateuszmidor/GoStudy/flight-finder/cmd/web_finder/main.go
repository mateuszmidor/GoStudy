package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

func main() {
	finder := newWebPathFinder("../../segments.csv.gz")
	runWEB(finder)
}

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("data", t.filename)))
	})
	data := map[string]interface{}{
		"Host": r.Host,
	}

	t.templ.Execute(w, data)
}

func runWEB(f *webPathFinder) {
	http.Handle("/", &templateHandler{filename: "index.html"})
	http.HandleFunc("/find", withFinder(f))
	http.ListenAndServe(":8080", nil)
}

func withFinder(f *webPathFinder) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		from := strings.ToUpper(r.URL.Query().Get("from"))
		to := strings.ToUpper(r.URL.Query().Get("to"))

		w.Header().Set("Content-Type", "application/text")

		start := time.Now()
		f.findConnections(from, to, w)
		d := time.Now().Sub(start)
		fmt.Fprintf(w, "Took %dms\n", d.Milliseconds())

		fmt.Printf("%s -> %s\n", from, to)
	}
}
