package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

func main() {
    // make a request to the Coingecko API for the latest Bitcoin price in PLN
    response, err := http.Get("https://api.coingecko.com/api/v3/simple/price?ids=bitcoin&vs_currencies=pln")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer response.Body.Close()

    // read the response body
    var data map[string]map[string]float64
    decoder := json.NewDecoder(response.Body)
    if err := decoder.Decode(&data); err != nil {
        fmt.Println("Error:", err)
        return
    }

    // get the latest Bitcoin price in PLN from the response
    price := data["bitcoin"]["pln"]

    // print the latest Bitcoin price in PLN
    fmt.Printf("Latest Bitcoin price in PLN: %.2f\n", price)
}

