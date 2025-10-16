package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

// Struct untuk menampung hasil response JSON dari API dummyjson.com
type Product struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

// Fungsi worker untuk mengambil data produk
func fetchProduct(id int, wg *sync.WaitGroup, resultsChan chan<- string, errorsChan chan<- error) {
	defer wg.Done()

	url := fmt.Sprintf("https://dummyjson.com/products/%d", id)
	resp, err := http.Get(url)
	if err != nil {
		errorsChan <- fmt.Errorf("ID %d: request error - %v", id, err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errorsChan <- fmt.Errorf("ID %d: HTTP error - status %d", id, resp.StatusCode)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		errorsChan <- fmt.Errorf("ID %d: read body error - %v", id, err)
		return
	}

	var product Product
	if err := json.Unmarshal(body, &product); err != nil {
		errorsChan <- fmt.Errorf("ID %d: JSON parse error - %v", id, err)
		return
	}

	resultsChan <- product.Title
}

func main() {
	productIDs := []int{1, 2, 3, 4, 5, 6, 7, 8}

	resultsChan := make(chan string, len(productIDs))
	errorsChan := make(chan error, len(productIDs))
	var wg sync.WaitGroup

	for _, id := range productIDs {
		wg.Add(1)
		go fetchProduct(id, &wg, resultsChan, errorsChan)
	}

	wg.Wait()

	// Tutup channel setelah semua goroutine selesai
	close(resultsChan)
	close(errorsChan)

	// Kumpulkan hasil
	var titles []string
	var errors []error

	for title := range resultsChan {
		titles = append(titles, title)
	}

	for err := range errorsChan {
		errors = append(errors, err)
	}

	// Cetak hasil akhir
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
