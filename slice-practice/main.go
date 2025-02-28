package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/google/uuid"
)

type Product struct {
	UUID  uuid.UUID `json:"uuid"`
	Title string    `json:"title"`
	Price float64   `json:"price"`
}

type Saver interface {
	Save() error
}

func New(title string, price float64) *Product {
	return &Product{
		UUID:  uuid.New(),
		Title: title,
		Price: price,
	}
}

func (p *Product) String() string {
	return fmt.Sprintf("UUID: %v\nTitle: %v\nPrice %.2f", p.UUID, p.Title, p.Price)
}

func main() {
	n := getNumberOfProducts()
	products := []*Product{}
	for range n {
		title, price := getProductInfo()
		product := New(title, price)
		products = append(products, product)
	}
	SaveProducts(products)
	products = ReadProducts("products.json")
	fmt.Println(products)
}

func getProductInfo() (string, float64) {
	reader := bufio.NewReader(os.Stdin)
	title := ""
	price := 0.0
	fmt.Print("Enter product title: ")
	title, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("error reading", err)
		return "", 0
	}
	title = strings.TrimSuffix(title, "\n")
	title = strings.TrimSuffix(title, "\r")
	fmt.Print("Enter product price: ")
	fmt.Scanln(&price)
	return title, price
}

func getNumberOfProducts() int {
	n := 0
	fmt.Print("Enter number of products: ")
	fmt.Scanln(&n)
	return n
}

func SaveProducts(products []*Product) error {
	json, err := json.MarshalIndent(products, "", "\t")
	if err != nil {
		return fmt.Errorf("error marshaling products slice: %w", err)
	}
	err = os.WriteFile("products.json", json, 0644)
	if err != nil {
		return fmt.Errorf("erro writing json: %w", err)
	}
	return nil
}

func ReadProducts(filename string) []*Product {
	data, err := os.ReadFile(filename)
	if err != nil {
		return []*Product{}
	}
	var products []*Product
	json.Unmarshal(data, &products)
	return products
}
