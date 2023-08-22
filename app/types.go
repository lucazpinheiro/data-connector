package main

import "time"

type BaseProduct struct {
	Source      string
	ID          int
	IsAvailable bool
	Title       string
	Price       float64
	Description string
	Images      []string
}

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

func (p *DummyProduct) Parse() BaseProduct {
	return BaseProduct{
		Source:      "dummy",
		ID:          p.ID,
		IsAvailable: p.Stock > 0,
		Title:       p.Title,
		Price:       float64(p.Price),
		Description: p.Description,
		Images:      p.Images,
	}
}

type DummyAPIResponse struct {
	Products []DummyProduct `json:"products"`
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

func (p *PlatziProduct) Parse() BaseProduct {
	return BaseProduct{
		Source:      "platzi",
		ID:          p.ID,
		IsAvailable: true,
		Title:       p.Title,
		Price:       float64(p.Price),
		Description: p.Description,
		Images:      p.Images,
	}
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

func (p *FakeStoreProduct) Parse() BaseProduct {
	return BaseProduct{
		Source:      "fakestore",
		ID:          p.ID,
		IsAvailable: true,
		Title:       p.Title,
		Price:       p.Price,
		Description: p.Description,
		Images:      []string{p.Image},
	}
}
