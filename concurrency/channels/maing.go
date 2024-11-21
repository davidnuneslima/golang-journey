package main

import "fmt"

func producer(ch chan string) {
	ch <- "Message from Producer"
	close(ch)
}

func main() {
	ch := make(chan string)
	go producer(ch)

	for msg := range ch {
		fmt.Println(msg)
	}
}
