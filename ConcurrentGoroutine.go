package main

import (
	"fmt"
	"time"
)

// //  WITHOUT GOROUTINE AND CHANNEL AND WITHOUT CUNCURRENCY
// func foo(someValue int) int {
// 	time.Sleep(2 * time.Second)

// 	return someValue * 5

// }

// func main() {
// 	started := time.Now()

// 	v1 := foo(5)
// 	v2 := foo(3)
// 	fmt.Println(v1, v2)
// 	fmt.Println("done in % d", time.Since(started))
// }

// USING GOROUTINE AND CHANNEL CONCURRENT WORKFLOW
func foo(c chan int, someValue int) {
	time.Sleep(2 * time.Second)

	c <- someValue * 5

}

func main() {
	started := time.Now()
	fooVal := make(chan int)
	go foo(fooVal, 5)
	go foo(fooVal, 3)
	v1 := <-fooVal
	v2 := <-fooVal
	fmt.Println(v1, v2)
	fmt.Println("done in % d", time.Since(started))
}
