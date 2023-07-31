package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

// https://dummyjson.com/products
type DummyProduct struct {
	ID                 int      `json:"id"`
	Title              string   `json:"title"`
	Description        string   `json:"description"`
	Price              int      `json:"price"`
	DiscountPercentage float64  `json:"discountPercentage"`
	Rating             float64  `json:"rating"`
	Stock              int      `json:"stock"`
	Brand              string   `json:"brand"`
	Category           string   `json:"category"`
	Thumbnail          string   `json:"thumbnail"`
	Images             []string `json:"images"`
}

// https://api.escuelajs.co/api/v1/products
type PlatziProduct struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Price       int       `json:"price"`
	Description string    `json:"description"`
	Images      []string  `json:"images"`
	CreationAt  time.Time `json:"creationAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	Category    struct {
		ID         int       `json:"id"`
		Name       string    `json:"name"`
		Image      string    `json:"image"`
		CreationAt time.Time `json:"creationAt"`
		UpdatedAt  time.Time `json:"updatedAt"`
	} `json:"category"`
}

// https://fakestoreapi.com/products
type FakeStoreProduct struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Image       string  `json:"image"`
	Rating      struct {
		Rate  float64 `json:"rate"`
		Count int     `json:"count"`
	} `json:"rating"`
}

func main() {
	getDummyProduct()
}

func getDummyProduct() {
	resp, err := http.Get("https://dummyjson.com/products")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var products []DummyProduct
	json.Unmarshal(data, &products)
	fmt.Println(products)
}
