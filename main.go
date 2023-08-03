package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	// dummyProducts := getDummyProducts()
	// platziProducts := getPlatziProducts()
	// fakeStoreProducts := getFakeStoreProducts()
	ctx := context.Background()
	// produce(ctx)
	// runEmitter()
	// runProcessor()
	writer := createWriter()
	sendProduct(ctx, writer, getDummyProducts())
}

func getDummyProducts() []BaseProduct {
	resp, err := http.Get("https://dummyjson.com/products")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var products DummyAPIResponse
	json.Unmarshal(data, &products)
	if err != nil {
		log.Fatal(err)
	}

	parsedProducts := make([]BaseProduct, len(products.Products))

	for i, product := range products.Products {
		parsedProducts[i] = product.Parse()
	}

	return parsedProducts
}

func getPlatziProducts() []BaseProduct {
	resp, err := http.Get("https://api.escuelajs.co/api/v1/products")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var products []PlatziProduct
	json.Unmarshal(data, &products)
	if err != nil {
		log.Fatal(err)
	}

	parsedProducts := make([]BaseProduct, len(products))
	for i, product := range products {
		parsedProducts[i] = product.Parse()
	}

	return parsedProducts
}

func getFakeStoreProducts() []BaseProduct {
	resp, err := http.Get("https://fakestoreapi.com/products")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var products []FakeStoreProduct
	json.Unmarshal(data, &products)
	if err != nil {
		log.Fatal(err)
	}

	parsedProducts := make([]BaseProduct, len(products))
	for i, product := range products {
		parsedProducts[i] = product.Parse()
	}

	return parsedProducts
}
