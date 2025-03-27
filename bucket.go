package tokenbucket

import "time"

type TokenBucket struct {
	tokens chan struct{}
	ticker time.Ticker
}

// Constructor for a new token bucket.
//
// maxConcurrent is the maximum number of tokens that can be in the bucket at any time. This is essentially the number of "workers" which can be active at a single time. e.g., with a refresh rate of 1s, and a maxConcurrent of 2, you can have 2 workers running at the same time, and a new worker can start every 1s.
//
// refillRate is the rate at which the bucket is refilled with tokens. This is the rate at which new workers can start.
func NewTokenBucket(maxConcurrent int, refillRate time.Duration) *TokenBucket {
	bucket := &TokenBucket{
		tokens: make(chan struct{}, maxConcurrent),
		ticker: *time.NewTicker(refillRate),
	}

	// Start the goroutine that will refill the bucket
	go bucket.replenish()

	return bucket
}

// replenish is a goroutine that will refill the bucket with tokens at a fixed rate
func (b *TokenBucket) replenish() {
	for range b.ticker.C {
		select {
		case b.tokens <- struct{}{}: // Add a token to the bucket if it's not full
		default: // If the channel is full, do nothing
		}
	}
}

// Wait will block until a token is available in the bucket
func (b *TokenBucket) Wait() {
	<-b.tokens
}

// Allow will return true if a token is available, otherwise false.
// Non-blocking; the caller must decide when to poll again.
func (b *TokenBucket) Allow() bool {
	select {
	case <-b.tokens:
		return true
	default:
		return false
	}
}
