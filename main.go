package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/segmentio/kafka-go"
)

func main() {
	writer := createWriter()

	queue := []string{"dummy", "fakestore", "platzi"}

	var wg sync.WaitGroup
	for _, source := range queue {
		wg.Add(1)
		go func(s string) {
			defer wg.Done()
			importCatalog(s, writer)
		}(source)
	}
	wg.Wait()
	writer.Close()
}

func importCatalog(source string, w *kafka.Writer) {
	ctx := context.Background()
	switch source {
	case "dummy":
		log.Println("Importing dummy products")
		importDummyProducts(ctx, w)
		log.Println("Finished dummy products")
	case "platzi":
		log.Println("Importing platzi products")
		importPlatziProducts(ctx, w)
		log.Println("Finished platzi products")
	case "fakestore":
		log.Println("Importing fakestore products")
		importFakeStoreProducts(ctx, w)
		log.Println("Finished fakestore products")
	default:
		fmt.Printf("Invalid source %s\n", source)
	}
}

func importDummyProducts(ctx context.Context, w *kafka.Writer) {
	resp, err := http.Get("https://dummyjson.com/products")

	if err != nil {
		log.Fatal(err)
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

	for _, product := range products.Products {
		sendProduct(ctx, w, product.Parse())
	}
}

func importPlatziProducts(ctx context.Context, w *kafka.Writer) {
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

	for _, product := range products {
		sendProduct(ctx, w, product.Parse())
	}
}

func importFakeStoreProducts(ctx context.Context, w *kafka.Writer) {
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

	for _, product := range products {
		sendProduct(ctx, w, product.Parse())
	}
}
