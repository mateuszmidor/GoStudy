// Project: http server serving item-price:
//			/list -				item-price list
//			/price?item=socks -	specific item price
// Usage: go run . & ncfirefox localhost:8000/price?item=shoes
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

func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/list":
		for item, price := range db {
			fmt.Fprintf(w, "%s: %s\n", item, price)
		}
	case "/price":
		item := req.URL.Query().Get("item")
		price, ok := db[item]
		if !ok {
			w.WriteHeader(http.StatusNotFound) // 404
			fmt.Fprintf(w, "no such item: %q\n", item)
			return
		}
		fmt.Fprintf(w, "%s: %s\n", item, price)
	default:
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such page: %q\n", req.URL)
	}

}

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.ListenAndServe("localhost:8000", db)
}