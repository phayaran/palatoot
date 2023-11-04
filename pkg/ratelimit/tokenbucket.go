package ratelimit

import (
	"fmt"
	"math"
	"time"
)

type TokenBucket struct {
	Tokens         float64
	MaxTokens      float64
	refillRate     float64
	lastRefillTime time.Time
}

func NewTokenBucket(maxTokens, refillRate float64) *TokenBucket {
	return &TokenBucket{
		Tokens:         maxTokens,
		MaxTokens:      maxTokens,
		refillRate:     refillRate,
		lastRefillTime: time.Now(),
	}
}

func (tb *TokenBucket) refill() {
	now := time.Now()
	duration := now.Sub(tb.lastRefillTime)
	tokensToAdd := tb.refillRate * duration.Seconds()
	fmt.Println("Adding tokens", tokensToAdd)
	tb.Tokens = math.Min(tb.Tokens+tokensToAdd, tb.MaxTokens)
	tb.lastRefillTime = now
}

func (tb *TokenBucket) Request(tokens float64) bool {
	tb.refill()
	if tokens <= tb.Tokens {
		tb.Tokens -= tokens
		return true
	}
	return false
}
