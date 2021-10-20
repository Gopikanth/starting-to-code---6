package main

import "fmt"

func main() {

	a := make(chan string)

	go func() {
		a <- "1"
		a <- "2"
		a <- "3"
		a <- "4"
		close(a)
	}()

	for v := range a { // Using for loop
		fmt.Println(v)
	}
}
