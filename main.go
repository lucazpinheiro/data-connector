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

	// queue := []string{"dummy", "fakestore", "platzi"}
	queue := []string{"dummy"}

	for _, source := range queue {
		go func(s string) {
			importCatalog(s, ch)
		}(source)
	}

	go func() {
		for {
			p, ok := <-ch
			if !ok {
				break
			}
			fmt.Println(p.ID)
			sendProduct(ctx, writer, p)
		}
		close(ch)
	}()

	// wg.Wait()
	writer.Close()
}

func importCatalog(source string, ch chan<- BaseProduct) {
	switch source {
	case "dummy":
		log.Println("Importing dummy products")
		importDummyProducts(ch)
		log.Println("Finished dummy products")
	case "platzi":
		log.Println("Importing platzi products")
		importPlatziProducts(ch)
		log.Println("Finished platzi products")
	case "fakestore":
		log.Println("Importing fakestore products")
		importFakeStoreProducts(ch)
		log.Println("Finished fakestore products")
	default:
		fmt.Printf("Invalid source %s\n", source)
	}
}

func importDummyProducts(ch chan<- BaseProduct) {
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
		ch <- product.Parse()
	}
}

func importPlatziProducts(ch chan<- BaseProduct) {
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
		ch <- product.Parse()
	}
}

func importFakeStoreProducts(ch chan<- BaseProduct) {
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
		ch <- product.Parse()
	}
}
