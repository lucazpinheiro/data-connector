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
	ctx := context.Background()

	writer := createWriter()

	ch := make(chan BaseProduct)

	queue := []string{"dummy", "fakestore"}

	for _, source := range queue {
		go importCatalog(source, ch)
	}

	for product := range ch {
		sendProduct(ctx, writer, product)
	}

	close(ch)
}

func importCatalog(source string, ch chan<- BaseProduct) {
	switch source {
	case "dummy":
		log.Println("Importing dummy products")
		importDummyProducts()
		for _, product := range importDummyProducts() {
			ch <- product
		}
	case "platzi":
		log.Println("Importing platzi products")
		importPlatziProducts()
		for _, product := range importPlatziProducts() {
			ch <- product
		}
	case "fakestore":
		log.Println("Importing fakestore products")
		importFakeStoreProducts()
		for _, product := range importFakeStoreProducts() {
			ch <- product
		}
	default:
		fmt.Println("Invalid source")
	}
}

func importDummyProducts() []BaseProduct {
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

func importPlatziProducts() []BaseProduct {
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

func importFakeStoreProducts() []BaseProduct {
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
