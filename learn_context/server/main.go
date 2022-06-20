package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"learn_context/log"
)

func main() {
	http.HandleFunc("/", log.Decorate(handler))
	panic(http.ListenAndServe("127.0.0.1:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	// qn: where is this context created?
	// at the place, where we can decide whether to cancel this context
	// qn: when cancel what will happen? ans: you decide what happen, need to explicitly handle it
	ctx := r.Context()
	// to show that int and key does not clash
	ctx = context.WithValue(ctx, int(42), int64(100))

	log.Println(ctx, "handler started")
	defer log.Println(ctx, "handler ended")

	// does not work
	fmt.Printf("value for foo is:", ctx.Value("foo"))

	select {
	case <-time.After(5 * time.Second):
		fmt.Fprintln(w, "hello")
	case <-ctx.Done():
		// Done; when client cancel the request
		// handle this case, so that will not do unnecessary processing
		// so as to stop processing when client cancel their request
		// when cancel, then internal server error
		// when cancel before writing to the response channel
		err := ctx.Err()
		log.Println(ctx, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// after request, process abt 5s then send back to response
	time.Sleep(5 * time.Second)
	fmt.Fprint(w, "hello")
}
