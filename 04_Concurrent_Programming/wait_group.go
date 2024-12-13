package main

import (
	"fmt"
	// Step 3,4 - Import
	"sync"
	"time"
)

// Calculate Total Sales
func main() {
	startTime := time.Now() // Record the current time
	var totalSales int
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(len(salesData))

	for _, salesRegion := range salesData {
		// Invoke to calculate Total for Each Region
		go calculateRegionSales(salesRegion, &totalSales, &wg, &mu)
	}

	// Allow goroutines to complete

	wg.Wait()

	fmt.Printf("Total %d , Time taken to calculate %s \n", totalSales, time.Since(startTime))
}

// Function to Calculate Total Sales per Region
func calculateRegionSales(salesRegion []int, totalSales *int, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	regionTotal := 0
	for _, storeSales := range salesRegion {
		// Calculate region total
		regionTotal += storeSales
		time.Sleep(100 * time.Millisecond)
	}
	mu.Lock()
	// Post Calculation of regionTotal
	*totalSales += regionTotal

	mu.Unlock()
}
