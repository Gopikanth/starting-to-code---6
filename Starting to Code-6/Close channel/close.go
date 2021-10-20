package main

import "fmt"

func main() {
	a := make(chan string)
	go next(a)
	for {
		c, ok := <-a
		if ok == false {
			fmt.Println("Channel Close ", ok)
			break
		}
		fmt.Println("Channel Open ", c, ok)

	}

}
func next(a chan string) {
	for x := 0; x < 4; x++ {
		a <- "Gopikanth"
	} //when getting outside the loop it closes
	close(a)

}
