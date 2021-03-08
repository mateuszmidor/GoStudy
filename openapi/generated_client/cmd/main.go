package main

import (
	"context"
	"fmt"

	openapi "github.com/mateuszmidor/GoStudy/openapi/generated_client"
)

func main() {
	cfg := openapi.NewConfiguration()
	client := openapi.NewAPIClient(cfg)

	getProducts(client)
}

func getProducts(client *openapi.APIClient) {
	req := client.DefaultApi.ProductsGet(context.Background())
	result, res, err := client.DefaultApi.ProductsGetExecute(req)

	fmt.Println("Result:",result)
	fmt.Println("Response:",res)
	fmt.Println("Error:",err)
}