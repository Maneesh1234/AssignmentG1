//Random int generator using channel
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// USING GOROUTINE AND CHANNEL CONCURRENT WORKFLOW
func foo(c chan int, someValue int) {
	time.Sleep(2 * time.Second)

	c <- rand.Intn(100 * someValue)

}

func main() {
	started := time.Now()
	fooVal := make(chan int)
	go foo(fooVal, 5)
	go foo(fooVal, 4)
	v1 := <-fooVal
	v2 := <-fooVal
	fmt.Println(v1, v2)
	fmt.Println("done in % d", time.Since(started))
}
