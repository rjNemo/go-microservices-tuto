package main

import (
	"fmt"
	"testing"

	"github.com/rjNemo/go-micro/client/client"
	"github.com/rjNemo/go-micro/client/client/products"
)

func TestOurClient(t *testing.T) {
	conf := client.DefaultTransportConfig().WithHost("localhost:5000")
	c := client.NewHTTPClientWithConfig(nil, conf)
	prod, err := c.Products.ListProducts(products.NewListProductsParams())
	if err != nil {
		t.Fatalf("%v", err)
	}
	fmt.Printf("%#v", prod.GetPayload()[0])
}
