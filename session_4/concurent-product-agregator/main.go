package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

// Struct untuk menampung response JSON dari API
type Product struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

func fetchProduct(id int, wg *sync.WaitGroup, titlesChan chan<- string, errChan chan<- string) {
	defer wg.Done()

	url := fmt.Sprintf("https://dummyjson.com/products/%d", id)
	resp, err := http.Get(url)
	if err != nil {
		errChan <- fmt.Sprintf("ID %d: request error - %v", id, err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		errChan <- fmt.Sprintf("ID %d: HTTP status %d", id, resp.StatusCode)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		errChan <- fmt.Sprintf("ID %d: read error - %v", id, err)
		return
	}

	var product Product
	if err := json.Unmarshal(body, &product); err != nil {
		errChan <- fmt.Sprintf("ID %d: JSON parse error - %v", id, err)
		return
	}

	titlesChan <- product.Title
}

func main() {
	productIDs := []int{1, 2, 3, 4, 5, 6, 7, 8}
	var wg sync.WaitGroup

	titlesChan := make(chan string, len(productIDs))
	errChan := make(chan string, len(productIDs))

	for _, id := range productIDs {
		wg.Add(1)
		go fetchProduct(id, &wg, titlesChan, errChan)
	}

	wg.Wait()
	close(titlesChan)
	close(errChan)

	var titles []string
	var errors []string

	for title := range titlesChan {
		titles = append(titles, title)
	}

	for e := range errChan {
		errors = append(errors, e)
	}

	fmt.Println("=== Product Titles ===")
	for _, t := range titles {
		fmt.Println("-", t)
	}

	fmt.Println("\n=== Errors ===")
	if len(errors) == 0 {
		fmt.Println("No errors occurred.")
	} else {
		for _, e := range errors {
			fmt.Println("-", e)
		}
	}
}
