package main

import (
	"fmt"
	"microservices/client/client"
	"microservices/client/client/products"
	"testing"
)

func TestOurClient(t *testing.T) {
	cfg:=client.DefaultTransportConfig().WithHost("localhost:9090")//config local host manually for testing
	c := client.NewHTTPClientWithConfig(nil,cfg)
	params := products.NewListProductsParams()
	prod,err:=c.Products.ListProducts(params)
	if err!=nil{
		t.Fatal(err)
	}
	fmt.Println(prod)
}
