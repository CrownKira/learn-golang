package log

import (
	"context"
	"log"
	"math/rand"
	"net/http"
)

type key int

// qn: how does key(42) works ?
// of type key
// type that only the log package can use
// why use 42 ?
const requestIDKey = key(42)

func Println(ctx context.Context, msg string) {
	// why request id is 42?
	// look for key of type key
	id, ok := ctx.Value(requestIDKey).(int64) // type assertion
	if !ok {
		log.Println("could not find request ID in context")
		return
	}

	log.Printf("[%d] %s", id, msg)
}

// type assertion
// https://stackoverflow.com/questions/24492868/what-is-the-meaning-of-dot-parenthesis-syntax
// var i interface{}
// i = int(42)

// a, ok := i.(int)
// // a == 42 and ok == true

// b, ok := i.(string)
// // b == "" (default value) and ok == false

func Decorate(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		id := rand.Int63()
		// context with value of prev context, the key then the id
		ctx = context.WithValue(ctx, requestIDKey, id)
		// decorate the request before invoking the handler
		f(w, r.WithContext(ctx))
	}
}
