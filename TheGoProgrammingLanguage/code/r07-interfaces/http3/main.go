// Project: http server using ServerMux, serving item-price:
//			/list -				item-price list
//			/price?item=socks -	specific item price
// Usage: ./run_all.sh
package main

import (
	"fmt"
	"net/http"
)

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("%.2f USD", d)
}

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s: %s\n", item, price)
}

func main() {
	db := database{"shoes": 50, "socks": 5}
	runDefaultMux(db)
}

// normally, we provide own multiplexer for handling different resources
func runCustomMux(db database) {
	customMux := http.NewServeMux()
	customMux.HandleFunc("/list", db.list)
	customMux.HandleFunc("/price", db.price)
	http.ListenAndServe("localhost:8000", customMux)
}

// for making things easy, http package provides global default multiplexer
func runDefaultMux(db database) {
	http.HandleFunc("/list", db.list)          // register handler to default mux
	http.HandleFunc("/price", db.price)        // register handler to default mux
	http.ListenAndServe("localhost:8000", nil) // nil means: please use default mux
}
