package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
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
		importDummyProducts(source, ctx, w)
	case "platzi":
		importPlatziProducts(source, ctx, w)
	case "fakestore":
		importFakeStoreProducts(source, ctx, w)
	default:
		fmt.Printf("Invalid source %s\n", source)
	}
}

func makeRequest(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return data, nil
}

func importDummyProducts(source string, ctx context.Context, w *kafka.Writer) {
	logImportStart(source)
	dummyURL := "https://dummyjson.com/products"
	data, err := makeRequest(dummyURL)
	if err != nil {
		log.Fatal(err)
	}

	var products DummyAPIResponse
	json.Unmarshal(data, &products)
	if err != nil {
		log.Fatal(err)
	}

	for _, product := range products.Products {
		parsedProduct := product.Parse()
		msgKey := fmt.Sprintf("%s:%d", parsedProduct.Source, parsedProduct.ID)
		msgData, _ := json.Marshal(parsedProduct)
		sendMessage(ctx, w, string(msgKey), msgData)
		logProduct(parsedProduct.Source, parsedProduct.ID)
	}
	logImportEnd(source)
}

func importPlatziProducts(source string, ctx context.Context, w *kafka.Writer) {
	logImportStart(source)
	platziURL := "https://api.escuelajs.co/api/v1/products"

	data, err := makeRequest(platziURL)
	if err != nil {
		log.Fatal(err)
	}

	var products []PlatziProduct
	json.Unmarshal(data, &products)
	if err != nil {
		log.Fatal(err)
	}

	for _, product := range products {
		parsedProduct := product.Parse()
		msgKey := fmt.Sprintf("%s:%d", parsedProduct.Source, parsedProduct.ID)
		msgData, _ := json.Marshal(parsedProduct)
		sendMessage(ctx, w, string(msgKey), msgData)
		logProduct(parsedProduct.Source, parsedProduct.ID)
	}
	logImportEnd(source)
}

func importFakeStoreProducts(source string, ctx context.Context, w *kafka.Writer) {
	logImportStart(source)
	fakestoreURL := "https://fakestoreapi.com/products"

	data, err := makeRequest(fakestoreURL)
	if err != nil {
		log.Fatal(err)
	}

	var products []FakeStoreProduct
	json.Unmarshal(data, &products)
	if err != nil {
		log.Fatal(err)
	}

	for _, product := range products {
		parsedProduct := product.Parse()
		msgKey := fmt.Sprintf("%s:%d", parsedProduct.Source, parsedProduct.ID)
		msgData, _ := json.Marshal(parsedProduct)
		sendMessage(ctx, w, string(msgKey), msgData)
		logProduct(parsedProduct.Source, parsedProduct.ID)
	}
	logImportEnd(source)
}
