package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
)

type catalog struct {
	Page   int      `xml:"page"`
	Fruits []string `xml:"fruits"`
}

func main() {
	// simple type
	// 1. encode
	inputBool := true
	xmlBool, _ := xml.Marshal(inputBool)
	fmt.Println(string(xmlBool))

	// 2. decode
	var outputBool bool
	xml.Unmarshal(xmlBool, &outputBool)
	fmt.Println(outputBool)

	// complex type
	// 1. encode
	inputCatalog := catalog{Page: 7, Fruits: []string{"apple", "pear", "grape"}}
	var xmlCatalog bytes.Buffer
	xml.NewEncoder(&xmlCatalog).Encode(&inputCatalog)
	fmt.Println(xmlCatalog.String())

	// 2. decode
	var outputCatalog catalog
	xml.NewDecoder(&xmlCatalog).Decode(&outputCatalog)
	fmt.Printf("%+v\n", outputCatalog)
}
