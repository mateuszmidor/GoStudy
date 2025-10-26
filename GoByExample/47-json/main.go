package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type catalog struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	// simple type
	// 1. encode
	inputBool := true
	jsonBool, _ := json.Marshal(inputBool)
	fmt.Println(string(jsonBool))

	// 2. decode
	var outputBool bool
	json.Unmarshal(jsonBool, &outputBool)
	fmt.Println(outputBool)

	// complex type
	// 1. encode
	inputCatalog := catalog{Page: 7, Fruits: []string{"apple", "pear", "grape"}}
	var jsonCatalog bytes.Buffer
	json.NewEncoder(&jsonCatalog).Encode(&inputCatalog)
	fmt.Print(jsonCatalog.String())

	// 2. decode
	var outputCatalog catalog
	json.NewDecoder(&jsonCatalog).Decode(&outputCatalog)
	fmt.Printf("%+v\n", outputCatalog)
}
