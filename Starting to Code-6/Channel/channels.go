package main

import (
	"fmt"
)

func main() {
	var a chan int //creating a channel using var
	fmt.Println(a)
	fmt.Printf("%T", a)
	new_channel := make(chan int) //creating a channel using make
	fmt.Println(new_channel)
	fmt.Printf("%v,%T", new_channel, new_channel)
}
