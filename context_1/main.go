package main

import (
	"context"
	"fmt"
	"time"
)

// propogate value in code
func enrichContext(ctx context.Context) context.Context {
	// must have context as first arg
	return context.WithValue(ctx, "request-id", "12345") // take in key value pair
}

// context is like hashmap that stores value and get passed around
// pass info between layers of application
// used only for things
// suplementray object: request id
// deadlines, timeout, etc
func doSomethingCool(ctx context.Context) {
	rID := ctx.Value("request-id") // get the request id
	// if context is time out context, it will time out
	// context is like information carrier but will timeout
	// it is just a hashmap that contains information and has a time limit
	fmt.Println(rID)
	for {
		// simulate a long running process
		select {
		// Done() returns a channel
		// <- channel: channel receiving something
		// when timeout, something will be sent to the channel
		case <-ctx.Done():
			fmt.Println("timed out")
			return
		default:
			fmt.Println("doing something cool")
		}
		// while loop every .5 seconds
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	fmt.Println("Go Context Tutorial")
	// context.Background() is empty; it is the basic parent
	// ctx := context.Background()
	// wrpa with deadline, wehn time frame exceed, done channel will be called
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	ctx = enrichContext(ctx)
	go doSomethingCool(ctx)
	select {
	case <-ctx.Done():
		fmt.Println("oh no, I have exceeded the deadline")
		// will show err if done() channel is called (under the hood)
		fmt.Println(ctx.Err())
	}

	// why no print timed out
	// so that will not end b4 other thread done

	time.Sleep(2 * time.Second)
}
