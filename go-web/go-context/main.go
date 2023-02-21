package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/add", fooAdd)
	http.HandleFunc("/cancel", fooCancel)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func foo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println(ctx)
	_, err := fmt.Fprintln(w, ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func fooAdd(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, "userID", 1234)
	ctx = context.WithValue(ctx, "firstName", "James")

	results := dbAccess(ctx)
	s := fmt.Sprintf("%v and %d", ctx, results)

	_, err := fmt.Fprintln(w, s)
	if err != nil {
		log.Fatal(err)
	}
}

func dbAccess(c context.Context) int {
	return c.Value("userID").(int)
}

func fooCancel(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx = context.WithValue(ctx, "userID", 5678)

	res, err := dbAccessSlow(ctx)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	_, err = fmt.Fprintln(w, res)
	if err != nil {
		log.Fatal(err)
	}
}

func computingTooLong(c context.Context, ch chan<- int) {
	// ridiculous log computation
	uid := c.Value("userID").(int)
	time.Sleep(10 * time.Second)
	if c.Err() != nil {
		return
	}
	ch <- uid
}

func dbAccessSlow(ctx context.Context) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	ch := make(chan int)

	// concurrent computing
	go computingTooLong(ctx, ch)

	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	case i := <-ch:
		return i, nil
	}
}
