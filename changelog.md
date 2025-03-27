# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/), and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [1.0.0] - 2025-03-27

### Added

- **Initial Release:**
  - Implemented the core token bucket rate limiter.
  - Public API includes:
    - `NewTokenBucket(maxConcurrent int, refillRate time.Duration) *TokenBucket` for creating a new token bucket.
    - `Allow() bool` for non-blocking token consumption.
    - `Wait()` for blocking until a token is available.
  - Refill mechanism implemented via a goroutine using a `time.Ticker`.
- **Documentation:**
  - Added a README with usage examples and API reference.
- **Testing:**
  - Included unit and integration tests to verify functionality.
