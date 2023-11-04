package ratelimit

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTokenBucketRateLimit(t *testing.T) {
	var got []bool
	want := []bool{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, false}
	tb := NewTokenBucket(10, 1)
	for i := 0; i < 20; i++ {
		got = append(got, tb.Request(1))
		time.Sleep(500 * time.Millisecond)
	}
	assert.Equal(t, got, want)
}
