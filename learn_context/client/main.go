package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

// less resource wasting
// Value(): must be something that is related to the request, should be extra info that is useful
// eg. request-id, etc
func main() {
	ctx := context.Background()
	// timeout after 1 second, if not getting response in 1 second
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	ctx = context.WithValue(ctx, "foo", "bar")
	defer cancel()

	// set the request
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
	// will allow u to customise the context
	// context field in request struct
	req = req.WithContext(ctx)

	// wait for response here
	// blocking call
	// do the request and get the response
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Fatal(res.Status)
	}

	// write the body to stdout
	io.Copy(os.Stdout, res.Body)
}
