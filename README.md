# TokenBucket

[![Build Status](https://github.com/hoodnoah/token_bucket/workflows/CI/badge.svg)](https://github.com/hoodnoah/token_bucket/actions)
[![codecov](https://codecov.io/gh/hoodnoah/token_bucket/branch/main/graph/badge.svg)](https://codecov.io/gh/hoodnoah/token_bucket)
[![Go Report Card](https://goreportcard.com/badge/github.com/hoodnoah/token_bucket)](https://goreportcard.com/report/github.com/hoodnoah/token_bucket)

**TokenBucket** is a simple Go library that implements a token bucket rate limiter. It is designed to help you control the rate of operations (such as external API calls or concurrent workers) by limiting the number of tokens available for execution. The bucket refills tokens at a fixed interval, allowing you to maximize throughput while adhering to rate limits.

## Features

- **Configurable Concurrency:**  
  Set the maximum number of tokens to control concurrent operations.
- **Adjustable Refill Rate:**  
  Specify the rate at which tokens are replenished.
- **Blocking and Non-blocking APIs:**  
  Use `Wait()` to block until a token is available, or `Allow()` to check token availability without blocking.

## Installation

To install the package, run:

```bash
go get github.com/hoodnoah/token_bucket
```

## Usage

Here's a quick example demonstrating how to use **TokenBucket**

```go
package main

import (
  "fmt"
  "time"

  tb "github.com/yourname/token_bucket" // rename for brevity
)

func main() {
  // Create a token bucket with a maximum of 2 concurrent tokens,
  // and a refill rate of 1 token per second
  bucket := tb.NewTokenBucket(2, time.Second)

  // Example 1: Non-blocking token check using Allow().
  if bucket.Allow() {
    fmt.Println("Token acquired (non-blocking); proceed with work...")
  } else {
    fmt.Println("No token available (non-blocking); try again later.")
  }

  // Example 2: Blocking token check using Wait().
  bucket.Wait() // blocks until a token is available, should ideally be no longer than the provided refill rate.
  fmt.Println("Token acquired (blocking); proceed with work...")
}
```
