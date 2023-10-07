package main

import (
	"gopkg.in/yaml.v3"
)

func main() {
	var output interface{}
	data := []byte("field: 123")
	yaml.Unmarshal(data, &output)
}
