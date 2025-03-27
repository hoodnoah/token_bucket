package tokenbucket_test

import (
	"sync"
	"testing"
	"time"

	tb "github.com/hoodnoah/token_bucket"
)

// TestAllow verifies that Allow() returns false before a refill, and true afterwards
func TestAllow(t *testing.T) {
	// Create a bucket with a capacity of 2 and a refill rate of 10ms.
	bucket := tb.NewTokenBucket(2, 10*time.Millisecond)

	// The bucket should start empty
	if bucket.Allow() {
		t.Errorf("Allow() should return false before any token is refilled.")
	}

	// Wait for a replenishment
	time.Sleep(20 * time.Millisecond)

	// Now a token should be available
	if !bucket.Allow() {
		t.Errorf("Allow() should return true when a token is available.")
	}
}

// TestWait verifies that Wait() blocks until a token is available
func TestWait(t *testing.T) {
	// Create a bucket with a capacity of 1 and a refill rate of 10ms.
	bucket := tb.NewTokenBucket(1, 10*time.Millisecond)

	start := time.Now()
	done := make(chan struct{})

	// Launch a goroutine which calls Wait()
	go func() {
		bucket.Wait()
		close(done)
	}()

	// The Wait() should block for at least 10ms (refill duration)
	select {
	case <-done:
		elapsed := time.Since(start)
		if elapsed < 10*time.Millisecond {
			t.Errorf("Wait() returned too quickly; elapsed: %v", elapsed)
		}
	case <-time.After(50 * time.Millisecond):
		t.Errorf("Wait() did not return after 50ms")
	}
}

// TestConcurrentAccess simulates concurrent calls to Allow() to verify the bucket does not allow more than capacity
func TestConcurrentAccess(t *testing.T) {
	// create a bucket
	bucket := tb.NewTokenBucket(2, 10*time.Millisecond)

	// Wait for token to fill
	time.Sleep(30 * time.Millisecond)

	var successes int

	for range 10 {
		if bucket.Allow() {
			successes++
		}
	}

	if successes > 2 {
		t.Errorf("Allow() should not return true more than the bucket capacity")
	}
}

// TestIntegration simulates multiple goroutines using Wait() and ensures the expected timing behavior.
func TestIntegration(t *testing.T) {
	// Create a bucket with a capacity of 2 and a refill interval of 50ms.
	bucket := tb.NewTokenBucket(2, 50*time.Millisecond)

	// We will start 3 goroutines that each wait for a token.
	var wg sync.WaitGroup
	wg.Add(3)

	// Channel to record the start times.
	startTimes := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		go func() {
			bucket.Wait()
			startTimes <- time.Now()
			wg.Done()
		}()
	}

	wg.Wait()
	close(startTimes)

	// Collect the times when each goroutine received a token.
	times := make([]time.Time, 0, 3)
	for ts := range startTimes {
		times = append(times, ts)
	}

	// Ensure that the third goroutine started significantly after the first,
	// because only two tokens are available initially.
	if len(times) != 3 {
		t.Fatalf("Expected 3 start times, got %d", len(times))
	}

	// Calculate time differences.
	diff := times[2].Sub(times[0])
	if diff < 45*time.Millisecond {
		t.Errorf("Expected the third Wait() to block for at least one refill interval; got diff: %v", diff)
	}
}
