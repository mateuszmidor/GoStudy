package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"text/template"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	data := map[string]interface{}{
		"API_ENDPOINT": os.Getenv("API_ENDPOINT"), // set in AWS Software configuration
	}
	t.templ.Execute(w, data)
}

func main() {
	for _, pair := range os.Environ() {
		fmt.Println(pair)
	}
	const awsServiceAddress = ":5000" // AWS runs the app at port 5000
	http.Handle("/", &templateHandler{filename: "main.html"})

	log.Println("Running the server at", awsServiceAddress)
	if err := http.ListenAndServe(awsServiceAddress, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
