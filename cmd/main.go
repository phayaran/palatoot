package main

import (
	"fmt"
	"palatoot/pkg/ratelimit"
	"time"
)

func main() {
	tb := ratelimit.NewTokenBucket(10, 1)
	for i := 0; i < 20; i++ {
		fmt.Printf("Request %d: %v\n", i+1, tb.Request(1))
		time.Sleep(500 * time.Millisecond)
	}
}
}
