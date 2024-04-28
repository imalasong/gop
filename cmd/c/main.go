package main

import "fmt"

func main() {
	fmt.Println(status(nil))
	ch := make(chan int)
	fmt.Printf("%T,%v\n", ch, ch)
	fmt.Println(status(ch))
	close(ch)
	fmt.Println(status(ch))

}

func status(ch chan int) bool {
	select {
	case <-ch:
		return true
	default:
		return false
	}
}
