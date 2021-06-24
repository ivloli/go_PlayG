package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/time/rate"
)

func main() {
	fmt.Println("vim-go")
	limiter := rate.NewLimiter(rate.Every(200*time.Millisecond), 10)
	fmt.Println(rate.Every(200 * time.Millisecond))
	for {
		limterHandler(limiter)
	}
}

func limterHandler(l *rate.Limiter) {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	e := l.Wait(ctx)
	fmt.Println(e)
}
