package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
)

// isPrime checks if a number is prime
func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// findMillionthPrime finds the nth prime number using concurrency
func findMillionthPrime(n int) int {
	primeCount := 0
	currentNum := 1
	numWorkers := runtime.NumCPU() // Use all available CPU cores

	// Channel for distributing numbers to check
	nums := make(chan int)
	// Channel for reporting primes
	primes := make(chan int)

	var wg sync.WaitGroup
	var stopLock sync.Mutex
	stopped := false

	stop := func() {
		stopLock.Lock()
		stopped = true
		stopLock.Unlock()
	}

	isStopped := func() bool {
		stopLock.Lock()
		defer stopLock.Unlock()
		return stopped
	}

	// Worker goroutines to check primes
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for num := range nums {
				if isStopped() {
					return
				}
				if isPrime(num) {
					primes <- num
				}
			}
		}()
	}

	// Goroutine to close the primes channel after all workers are done
	go func() {
		wg.Wait()
		close(primes)
	}()

	// Goroutine to generate numbers to check
	go func() {
		for {
			if isStopped() {
				close(nums)
				return
			}
			currentNum++
			nums <- currentNum
		}
	}()

	// Read primes and count until we reach the nth prime
	var nthPrime int
	for prime := range primes {
		primeCount++
		if primeCount == n {
			nthPrime = prime
			stop() // Signal to stop the number generator
			break
		}
	}

	return nthPrime
}

func main() {
	millionthPrime := findMillionthPrime(1_000_000)
	fmt.Printf("The millionth prime number is: %d\n", millionthPrime)
}
