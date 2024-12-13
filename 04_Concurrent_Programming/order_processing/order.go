package main

import "fmt"

type order struct {
	ProductCOde int
	Quantity    float64
	Status      orderStatus
}

type invalidOrder struct {
	order order
	err   error
}

func (o order) String() string {
	return fmt.Sprintf("Product code: %v, Quantity: %v, Status: %v \n",
		o.ProductCOde, o.Quantity, orderStatusToText(o.Status))
}

func orderStatusToText(o orderStatus) string {
	switch o {
	case none:
		return "none"
	case new:
		return "new"
	case recieved:
		return "recieved"
	case reserved:
		return "reserved"
	case filled:
		return "filled"
	default:
		return "Unknown Status"
	}
}

type orderStatus int

const (
	none orderStatus = iota
	new
	recieved
	reserved
	filled
)

var orders = []order{}
