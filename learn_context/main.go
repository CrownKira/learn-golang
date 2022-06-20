// cancellation propogation
// value propogation: 1 method
// context.background() not reacting to anything, root context
// the rest create from that context
package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	// context allocates some resources for timer
	// release timer resources
	// will send to Done channel
	defer cancel()
	// ctx, cancel := context.WithCancel(ctx)

	// go func() {
	// 	s := bufio.NewScanner(os.Stdin)
	// 	s.Scan()
	// 	// cancel as soon as it reads sth from std input
	// 	// cancel ie. Done() will timeout
	// 	cancel()
	// }()
	// time.AfterFunc(time.Second, cancel)

	mySleepAndTalk(ctx, 5*time.Second, "hello")
}

func mySleepAndTalk(ctx context.Context, d time.Duration, msg string) {
	// sleep for d
	// print msg
	// after() returns a channel value type
	select {
	case <-time.After(d):
		fmt.Println(msg)
	case <-ctx.Done():
		log.Print(ctx.Err())
	}
	fmt.Println(msg)
}
