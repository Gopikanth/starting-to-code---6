package main

import "fmt"

func main() {
	fmt.Println("creating a channel")
	a := make(chan string)
	go next(a)
	a <- "Gopikanth"
	fmt.Println("ending the channel")

}
func next(a chan string) {
	fmt.Println(<-a+ "hello")

}
