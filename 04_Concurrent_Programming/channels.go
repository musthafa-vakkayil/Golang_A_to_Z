package main

import (
	"fmt"
	"time"
)

// Calculate Total Sales
func channels() {
	startTime := time.Now() // Record the current time
	var totalSales int
	salesCh := make(chan int)
	errorCh := make(chan error)

	for _, salesRegion := range salesData {
		// Invoke to calculate Total for Each Region
		go calculateRegionSalesUsingChannel(salesRegion, salesCh, errorCh)
	}

	// Allow goroutines to complete
	// Receive total sales from each goroutine asynchronously
	for i := 0; i < len(salesData); i++ {
		select {
		case total := <-salesCh:
			totalSales += total
		case err := <-errorCh:
			fmt.Printf("Error: %v\n", err)
		case <-time.After(3000 * time.Millisecond):
			fmt.Println("Timeout occurred while waiting for sales data")
			return
		}
	}

	fmt.Printf("Total %d , Time taken to calculate %s \n", totalSales, time.Since(startTime))
}

// Function to Calculate Total Sales per Region
func calculateRegionSalesUsingChannel(salesRegion []int, salesCh chan<- int, errorCh chan<- error) {
	regionTotal := 0
	for _, storeSales := range salesRegion {
		// Calculate region total
		if storeSales < 0 {
			errorCh <- fmt.Errorf("SKIPPED REGION! Found store with sale value [%d] less than 0. Please recheck data for region", storeSales)
			return
		}
		regionTotal += storeSales
		time.Sleep(100 * time.Millisecond)
	}

	salesCh <- regionTotal
}
