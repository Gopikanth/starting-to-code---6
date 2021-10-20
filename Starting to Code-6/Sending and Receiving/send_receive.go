package main

import "fmt"

func main() {
	fmt.Println("creating the channel")

	b := make(chan int)

	go new_function(b)
	b <- 67
	fmt.Println("Ending the channel")

}

func new_function(b chan int) {
	var a int
	a = 234 + <-b
	fmt.Println("Channel Received sucessfully", a)
}
