# Go Race Condition with Mutex and Goroutines

This repository demonstrates a subtle race condition in a Go program that uses goroutines and a mutex.  The program appears to be thread-safe, but under certain conditions, it produces incorrect output. This example highlights the importance of carefully considering how goroutines interact, even when using synchronization primitives like mutexes.

## The Bug
The `bug.go` file contains a Go program that launches multiple goroutines, each incrementing a counter protected by a mutex.  While seemingly correct, if the `time.Sleep` is removed, race conditions may occur leading to unpredictable output.

## The Solution
The `bugSolution.go` file demonstrates a correct solution that addresses the race condition by ensuring proper synchronization between goroutines and avoids the potential data race conditions.

## How to Reproduce
1. Clone this repository.
2. Run the `bug.go` file.
3. Observe the output.  Note any unexpected orderings of the numbers, particularly when `time.Sleep` is removed.
4. Compare the output with that of the `bugSolution.go` file.

This example showcases the importance of careful concurrency programming in Go and the need to avoid subtle race conditions.