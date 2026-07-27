package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := "abc123!?$*&()'-=@~"

	stringEncoded := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(stringEncoded)

	stringDecoded, _ := base64.StdEncoding.DecodeString(stringEncoded)
	fmt.Println(string(stringDecoded))

	urlEncoded := base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(urlEncoded)

	urlDecoded, _ := base64.URLEncoding.DecodeString(urlEncoded)
	fmt.Println(string(urlDecoded))
}
