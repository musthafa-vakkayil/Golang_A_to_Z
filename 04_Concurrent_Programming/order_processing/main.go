package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var recievedOrdersCh = make(chan order)
	var validOrderCh = make(chan order)
	var invalidOrderCh = make(chan invalidOrder)

	go recieveOrders(recievedOrdersCh)
	go validateOrders(recievedOrdersCh, validOrderCh, invalidOrderCh)
	wg.Add(1)
	go func() {
		vaorder := <-validOrderCh
		fmt.Printf("Valid Order recieved: %v\n", vaorder.ProductCOde)
		wg.Done()
	}()

	go func() {
		inorder := <-invalidOrderCh
		fmt.Printf("invalid Order recieved: %v\n and error is: %v", inorder.order.ProductCOde, inorder.err)
		wg.Done()
	}()

	wg.Wait()
}

var rawOrders = []string{
	`{"productCode": 1111, "quantity": -5, "status": 1}`,
	`{"productCode": 2222, "quantity": 42.3, "status": 1}`,
	`{"productCode": 3333, "quantity": 19, "status": 1}`,
	`{"productCode": 4444, "quantity": 8, "status": 1}`,
}

func recieveOrders(out chan order) {
	for _, rawOrder := range rawOrders {
		var newOrder order
		err := json.Unmarshal([]byte(rawOrder), &newOrder)

		if err != nil {
			log.Println(err)
			continue
		}

		out <- newOrder
	}
}

func validateOrders(in chan order, out chan order, errCh chan invalidOrder) {
	order := <-in
	if order.Quantity <= 0 {
		// error condition
		errCh <- invalidOrder{order: order, err: errors.New("Invalid Quantity")}
	} else {
		out <- order
	}
}
