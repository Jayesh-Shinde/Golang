package main

import (
	"fmt"
	"time"
)

func greet(phrase string, done chan bool) {
	fmt.Println("Hello!", phrase)
	done <- true
}

func slowGreet(phrase string, done chan bool) {
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)
	done <- true
	//close(done)
}

func main() {
	done := make(chan bool, 4) // channel is basically a way of awaiting the async function
	// like awaiting a goroutine
	go greet("Nice to meet you!", done)
	go greet("How are you?", done)
	go slowGreet("How ... are ... you ...?", done)
	go greet("I hope you're liking the course!", done)

	//for value := range done {
	for range done {
		//fmt.Println("value: ", value)
	}
}

/*
below is the way it works closely with promise.all in node js

func main() {
    var (
        user      User
        orders    []Order
        userErr   error
        ordersErr error
        wg        sync.WaitGroup
    )

    wg.Add(2)
    go func() {
        defer wg.Done()
        user, userErr = fetchUser()
    }()

    go func() {
        defer wg.Done()
        orders, ordersErr = fetchOrders()
    }()

    wg.Wait()

    // Now you can handle them specifically
    if ordersErr != nil {
        log.Fatalf("CRITICAL: Orders failed: %v", ordersErr) // Hard stop
    }

    if userErr != nil {
        fmt.Println("Warning: User profile not found, continuing...") // Log and continue
    }
}

*/
