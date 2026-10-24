package main

import "fmt"

type OrderStatus int

const (
	Pending   OrderStatus = iota //0
	Shipped                      //1
	Delivered                    //2
	Canceled                     //3
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("changing order status to", status)
}

func main() {
	// Example usage
	changeOrderStatus(Canceled)
}
